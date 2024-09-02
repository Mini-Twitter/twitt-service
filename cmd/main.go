package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"twitt-service/genproto/tweet"
	"twitt-service/pkg/config"
	"twitt-service/pkg/logger"
	messagebroker "twitt-service/pkg/rebbitmq"
	"twitt-service/service"
	"twitt-service/storage/postgres"
)

func main() {
	logger := logger.InitLogger()
	cfg := config.Load()

	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		logger.Error("Error connecting to database")
		log.Fatal(err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@api-gateway-rabbit-1:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	cr1, err := getQueue(ch, "PostComment")
	if err != nil {
		log.Fatalf("Failed to declare PostComment queue: %v", err)
	}

	cr2, err := getQueue(ch, "UpdateComment")
	if err != nil {
		log.Fatalf("Failed to declare UpdateComment queue: %v", err)
	}

	cr3, err := getQueue(ch, "PostTweet")
	if err != nil {
		log.Fatalf("Failed to declare PostTweet queue: %v", err)
	}

	cr4, err := getQueue(ch, "UpdateTweet")
	if err != nil {
		log.Fatalf("Failed to declare UpdateTweet queue: %v", err)
	}

	cr5, err := getQueue(ch, "AddLike")
	if err != nil {
		log.Fatalf("Failed to declare AddLike queue: %v", err)
	}

	rc1, err := getMessageQueue(ch, cr1)
	if err != nil {
		log.Fatalf("Failed to consume PostComment messages: %v", err)
	}

	rc2, err := getMessageQueue(ch, cr2)
	if err != nil {
		log.Fatalf("Failed to consume UpdateComment messages: %v", err)
	}

	rc3, err := getMessageQueue(ch, cr3)
	if err != nil {
		log.Fatalf("Failed to consume PostTweet messages: %v", err)
	}

	rc4, err := getMessageQueue(ch, cr4)
	if err != nil {
		log.Fatalf("Failed to consume UpdateTweet messages: %v", err)
	}

	rc5, err := getMessageQueue(ch, cr5)
	if err != nil {
		log.Fatalf("Failed to consume AddLike messages: %v", err)
	}

	tweetSt := postgres.NewTweetRepo(db)

	tweetSt1 := postgres.NewCommentRepo(db)

	tweetSt2 := postgres.NewLikeRepo(db)
	tweetSr := service.NewTweetService(tweetSt, tweetSt2, tweetSt1, logger)

	listen, err := net.Listen("tcp", cfg.TWITT_SERVICE)
	fmt.Println("Listening on " + cfg.TWITT_SERVICE)
	if err != nil {
		logger.Error("Error listening on port "+cfg.TWITT_SERVICE, "error", err)
		log.Fatal(err)
	}

	res := messagebroker.New(tweetSr, ch, rc1, rc2, rc3, rc5, rc4, &sync.WaitGroup{}, 5, db)
	go res.StartToConsume(context.Background())

	server := grpc.NewServer()
	tweet.RegisterTweetServiceServer(
		server,
		tweetSr,
	)

	if err := server.Serve(listen); err != nil {
		logger.Error("Error starting server on port "+cfg.TWITT_SERVICE, "error", err)
		log.Fatal(err)
	}

}

func getQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
}

func getMessageQueue(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}
