package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type SubscribeRepo struct {
	db *sqlx.DB
}

func NewSubscribeRepo(db *sqlx.DB) storage.SubscribeStorage {
	return &SubscribeRepo{
		db: db,
	}
}

func (s *SubscribeRepo) Follow(in *pb.FollowReq) (*pb.FollowRes, error) {
	return nil, nil
}

func (s *SubscribeRepo) Unfollow(in *pb.FollowReq) (*pb.DFollowRes, error) {
	return nil, nil
}

func (s *SubscribeRepo) GetUserFollowers(in *pb.UserId) (*pb.Count, error) {
	return nil, nil
}

func (s *SubscribeRepo) GetUserFollows(in *pb.UserId) (*pb.Count, error) {
	return nil, nil
}

func (s *SubscribeRepo) MostPopularUser(in *pb.Void) (*pb.User, error) {
	return nil, nil
}
