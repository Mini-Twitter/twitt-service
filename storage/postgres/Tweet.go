package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type TweetRepo struct {
	db *sqlx.DB
}

func NewTweetRepo(db *sqlx.DB) storage.TweetStorage {
	return &TweetRepo{
		db: db,
	}
}

func (t *TweetRepo) PostTweet(in *pb.Tweet) (*pb.TweetResponse, error) {
	return nil, nil
}

func (t *TweetRepo) UpdateTweet(in *pb.UpdateATweet) (*pb.TweetResponse, error) {
	return nil, nil
}

func (t *TweetRepo) AddImageToTweet(in *pb.Url) (*pb.Message, error) {
	return nil, nil
}

func (t *TweetRepo) UserTweets(in *pb.UserId) (*pb.Tweets, error) {
	return nil, nil
}

func (t *TweetRepo) GetTweet(in *pb.TweetId) (*pb.TweetResponse, error) {
	return nil, nil
}

func (t *TweetRepo) GetAllTweets(in *pb.TweetFilter) (*pb.Tweets, error) {
	return nil, nil
}

func (t *TweetRepo) RecommendTweets(in *pb.UserId) (*pb.Tweets, error) {
	return nil, nil
}

func (t *TweetRepo) GetNewTweets(in *pb.UserId) (*pb.Tweets, error) {
	return nil, nil
}
