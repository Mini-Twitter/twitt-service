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
	res, err := s.storage.PostTweet(in)
	if err != nil {
		s.logger.Error("failed to post tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *tweetService) UpdateTweet(in *pb.UpdateATweet) (*pb.TweetResponse, error) {
	res, err := s.storage.UpdateTweet(in)
	if err != nil {
		s.logger.Error("failed to update tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *tweetService) AddImageToTweet(in *pb.Url) (*pb.Message, error) {
	res, err := s.storage.AddImageToTweet(in)
	if err != nil {
		s.logger.Error("failed to add image to tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *tweetService) UserTweets(in *pb.UserId) (*pb.Tweets, error) {
	res, err := s.storage.UserTweets(in)
	if err != nil {
		s.logger.Error("failed to user tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *tweetService) GetTweet(in *pb.TweetId) (*pb.TweetResponse, error) {
	res, err := s.storage.GetTweet(in)
	if err != nil {
		s.logger.Error("failed to get tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *tweetService) GetAllTweets(in *pb.TweetFilter) (*pb.Tweets, error) {
	res, err := s.storage.GetAllTweets(in)
	if err != nil {
		s.logger.Error("failed to get tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *tweetService) RecommendTweets(in *pb.UserId) (*pb.Tweets, error) {
	res, err := s.storage.RecommendTweets(in)
	if err != nil {
		s.logger.Error("failed to recommend tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *tweetService) GetNewTweets(in *pb.UserId) (*pb.Tweets, error) {
	res, err := s.storage.GetNewTweets(in)
	if err != nil {
		s.logger.Error("failed to get tweets", err)
		return nil, err
	}
	return res, nil
}
