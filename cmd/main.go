package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"twitt-service/genproto/tweet"
	"twitt-service/pkg/config"
	"twitt-service/pkg/logger"
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

	tweetSt := postgres.NewTweetRepo(db)

	tweetSt1 := postgres.NewCommentRepo(db)

	tweetSt2 := postgres.NewLikeRepo(db)
	tweetSr := service.NewTweetService(tweetSt, tweetSt2, tweetSt1, logger)

	listen, err := net.Listen("tcp", cfg.TWITT_SERVICE)
	if err != nil {
		logger.Error("Error listening on port "+cfg.TWITT_SERVICE, "error", err)
		log.Fatal(err)
	}

	go func() {
		server := grpc.NewServer()
		tweet.RegisterTweetServiceServer(
			server,
			tweetSr,
		)
		logger.Info("Starting server on port " + cfg.TWITT_SERVICE)
		log.Panicln("Starting server on port " + cfg.TWITT_SERVICE)

		if err := server.Serve(listen); err != nil {
			logger.Error("Error starting server on port "+cfg.TWITT_SERVICE, "error", err)
			log.Fatal(err)
		}
	}()
}
