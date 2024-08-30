package rebbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"twitt-service/genproto/tweet"
	"twitt-service/pkg/logger"
	"twitt-service/service"
)

type (
	MsgBroker struct {
		service          *service.TweetService
		channel          *amqp.Channel
		PostComment      <-chan amqp.Delivery
		UpdateComment    <-chan amqp.Delivery
		AddLike          <-chan amqp.Delivery
		PostTweet        <-chan amqp.Delivery
		UpdateTweet      <-chan amqp.Delivery
		logger           *slog.Logger
		wg               *sync.WaitGroup
		numberOfServices int
		Db               *sqlx.DB
	}
)

func New(
	service *service.TweetService,
	channel *amqp.Channel,
	PostComment <-chan amqp.Delivery,
	UpdateComment <-chan amqp.Delivery,
	AddLike <-chan amqp.Delivery,
	PostTweet <-chan amqp.Delivery,
	UpdateTweet <-chan amqp.Delivery,
	wg *sync.WaitGroup,
	numberOfServices int,
	Db *sqlx.DB) *MsgBroker {
	return &MsgBroker{
		service:          service,
		channel:          channel,
		PostComment:      PostComment,
		UpdateComment:    UpdateComment,
		AddLike:          AddLike,
		PostTweet:        PostTweet,
		UpdateTweet:      UpdateTweet,
		logger:           logger.InitLogger(),
		wg:               wg,
		numberOfServices: numberOfServices,
		Db:               Db,
	}
}

func (m *MsgBroker) StartToConsume(ctx context.Context) {
	m.wg.Add(m.numberOfServices)
	consumerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go m.consumeMessages(consumerCtx, m.PostComment, "PostComment")
	go m.consumeMessages(consumerCtx, m.UpdateComment, "UpdateComment")
	go m.consumeMessages(consumerCtx, m.AddLike, "AddLike")
	go m.consumeMessages(consumerCtx, m.PostTweet, "PostTweet")
	go m.consumeMessages(consumerCtx, m.UpdateTweet, "UpdateTweet")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	m.logger.Info("Shutting down, waiting for consumers to finish")
	cancel()
	m.wg.Wait()
	m.logger.Info("All consumers have stopped")
}

func (m *MsgBroker) consumeMessages(ctx context.Context, messages <-chan amqp.Delivery, logPrefix string) {
	defer m.wg.Done()
	for {
		select {
		case val := <-messages:
			var err error

			switch logPrefix {
			case "PostComment":
				var req tweet.Comment

				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					return
				}
				_, err = m.service.PostComment(ctx, &req)
				if err != nil {
					m.logger.Error("Error while creating booking", "error", err)
					val.Nack(false, false)
					return
				}
				val.Ack(false)

				fmt.Println(req)

			case "UpdateComment":
				var req tweet.UpdateAComment

				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					return
				}
				_, err = m.service.UpdateComment(ctx, &req)

				fmt.Println(req)
			case "PostTweet":
				var req tweet.Tweet

				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					return
				}
				_, err = m.service.PostTweet(ctx, &req)

				fmt.Println(req)

			case "UpdateTweet":
				var req tweet.UpdateATweet

				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					return
				}
				_, err = m.service.UpdateTweet(ctx, &req)

				fmt.Println(req)

			case "AddLike":
				var req tweet.LikeReq
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					return
				}
				_, err = m.service.AddLike(ctx, &req)

				fmt.Println(req)

			}

			if err != nil {
				m.logger.Error("Failed in %s: %s", logPrefix, err.Error())
				val.Nack(false, false)
				return
			}

			val.Ack(false)

		case <-ctx.Done():
			m.logger.Info("Context done, stopping consumer", "consumer", logPrefix)
			return
		}
	}
}
