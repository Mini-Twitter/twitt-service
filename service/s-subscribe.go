package service

import (
	"log/slog"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type SubscribeService interface {
	Follow(in *pb.FollowReq) (*pb.FollowRes, error)
	Unfollow(in *pb.FollowReq) (*pb.DFollowRes, error)
	GetUserFollowers(in *pb.UserId) (*pb.Count, error)
	GetUserFollows(in *pb.UserId) (*pb.Count, error)
	MostPopularUser(in *pb.Void) (*pb.User, error)
}

type subscribeService struct {
	storage storage.SubscribeStorage
	logger  *slog.Logger
}

func NewSubscribeService(st storage.SubscribeStorage, logger *slog.Logger) SubscribeService {
	return &subscribeService{
		storage: st,
		logger:  logger,
	}
}

func (s *subscribeService) Follow(in *pb.FollowReq) (*pb.FollowRes, error) {
	return &pb.FollowRes{}, nil
}

func (s *subscribeService) Unfollow(in *pb.FollowReq) (*pb.DFollowRes, error) {
	return &pb.DFollowRes{}, nil
}

func (s *subscribeService) GetUserFollowers(in *pb.UserId) (*pb.Count, error) {
	return &pb.Count{}, nil
}

func (s *subscribeService) GetUserFollows(in *pb.UserId) (*pb.Count, error) {
	return &pb.Count{}, nil
}

func (s *subscribeService) MostPopularUser(in *pb.Void) (*pb.User, error) {
	return &pb.User{}, nil
}
