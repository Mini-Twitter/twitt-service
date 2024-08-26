package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type LikeRepo struct {
	db *sqlx.DB
}

func NewLikeRepo(db *sqlx.DB) storage.LikesStorage {
	return &LikeRepo{
		db: db,
	}
}

func (l *LikeRepo) AddLike(in *pb.LikeReq) (*pb.LikeRes, error) {
	return &pb.LikeRes{}, nil
}

func (l *LikeRepo) DeleteLIke(in *pb.LikeReq) (*pb.DLikeRes, error) {
	return &pb.DLikeRes{}, nil
}

func (l *LikeRepo) GetUserLikes(in *pb.UserId) (*pb.TweetTitles, error) {
	return &pb.TweetTitles{}, nil
}

func (l *LikeRepo) GetCountTweetLikes(in *pb.TweetId) (*pb.Count, error) {
	return &pb.Count{}, nil
}

func (l *LikeRepo) MostLikedTweets(in *pb.Void) (*pb.Tweet, error) {
	return &pb.Tweet{}, nil
}
