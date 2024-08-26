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
	return &pb.LikeRes{}, nil
}

func (s *likeService) DeleteLIke(in *pb.LikeReq) (*pb.DLikeRes, error) {
	return &pb.DLikeRes{}, nil
}

func (s *likeService) GetUserLikes(in *pb.UserId) (*pb.TweetTitles, error) {
	return &pb.TweetTitles{}, nil
}

func (s *likeService) GetCountTweetLikes(in *pb.TweetId) (*pb.Count, error) {
	return &pb.Count{}, nil
}

func (s *likeService) MostLikedTweets(in *pb.Void) (*pb.Tweet, error) {
	return &pb.Tweet{}, nil
}
