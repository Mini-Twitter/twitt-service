package service

import (
	"log/slog"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type TweetService interface {
	PostTweet(in *pb.Tweet) (*pb.TweetResponse, error)
	UpdateTweet(in *pb.UpdateATweet) (*pb.TweetResponse, error)
	AddImageToTweet(in *pb.Url) (*pb.Message, error)
	UserTweets(in *pb.UserId) (*pb.Tweets, error)
	GetTweet(in *pb.TweetId) (*pb.TweetResponse, error)
	GetAllTweets(in *pb.TweetFilter) (*pb.Tweets, error)
	RecommendTweets(in *pb.UserId) (*pb.Tweets, error)
	GetNewTweets(in *pb.UserId) (*pb.Tweets, error)
}

type tweetService struct {
	storage storage.TweetStorage
	logger  *slog.Logger
}

func NewTweetService(st storage.TweetStorage, logger *slog.Logger) TweetService {
	return &tweetService{
		storage: st,
		logger:  logger,
	}
}

func (s *tweetService) PostTweet(in *pb.Tweet) (*pb.TweetResponse, error) {
	return nil, nil
}

func (s *tweetService) UpdateTweet(in *pb.UpdateATweet) (*pb.TweetResponse, error) {
	return nil, nil
}

func (s *tweetService) AddImageToTweet(in *pb.Url) (*pb.Message, error) {
	return nil, nil
}

func (s *tweetService) UserTweets(in *pb.UserId) (*pb.Tweets, error) {
	return nil, nil
}

func (s *tweetService) GetTweet(in *pb.TweetId) (*pb.TweetResponse, error) {
	return nil, nil
}

func (s *tweetService) GetAllTweets(in *pb.TweetFilter) (*pb.Tweets, error) {
	return nil, nil
}

func (s *tweetService) RecommendTweets(in *pb.UserId) (*pb.Tweets, error) {
	return nil, nil
}

func (s *tweetService) GetNewTweets(in *pb.UserId) (*pb.Tweets, error) {
	return nil, nil
}
