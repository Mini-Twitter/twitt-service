package service

import (
	"log/slog"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type LikeService interface {
	AddLike(in *pb.LikeReq) (*pb.LikeRes, error)
	DeleteLIke(in *pb.LikeReq) (*pb.DLikeRes, error)
	GetUserLikes(in *pb.UserId) (*pb.TweetTitles, error)
	GetCountTweetLikes(in *pb.TweetId) (*pb.Count, error)
	MostLikedTweets(in *pb.Void) (*pb.Tweet, error)
}

type likeService struct {
	storage storage.LikesStorage
	logger  *slog.Logger
}

func NewLikeService(st storage.LikesStorage, logger *slog.Logger) LikeService {
	return &likeService{
		storage: st,
		logger:  logger,
	}
}

func (s *likeService) AddLike(in *pb.LikeReq) (*pb.LikeRes, error) {
	res, err := s.storage.AddLike(in)
	if err != nil {
		s.logger.Error("failed to create liked tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *likeService) DeleteLIke(in *pb.LikeReq) (*pb.DLikeRes, error) {
	res, err := s.storage.DeleteLIke(in)
	if err != nil {
		s.logger.Error("failed to delete liked tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *likeService) GetUserLikes(in *pb.UserId) (*pb.TweetTitles, error) {
	res, err := s.storage.GetUserLikes(in)
	if err != nil {
		s.logger.Error("failed to get liked tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *likeService) GetCountTweetLikes(in *pb.TweetId) (*pb.Count, error) {
	res, err := s.storage.GetCountTweetLikes(in)
	if err != nil {
		s.logger.Error("failed to get liked tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *likeService) MostLikedTweets(in *pb.Void) (*pb.Tweet, error) {
	res, err := s.storage.MostLikedTweets(in)
	if err != nil {
		s.logger.Error("failed to get liked tweets", err)
		return nil, err
	}
	return res, nil
}
