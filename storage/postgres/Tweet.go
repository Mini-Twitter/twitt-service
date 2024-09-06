package postgres

import (
	"context"
	"database/sql"
	"time"
	"twitt-service/storage"

	"github.com/jmoiron/sqlx"
	pb "twitt-service/genproto/tweet"
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
	query := `INSERT INTO tweets (user_id, hashtag, title, content, image_url) 
	          VALUES ($1, $2, $3, $4, $5) 
	          RETURNING id, user_id, hashtag, title, content, image_url, created_at, updated_at`

	var res pb.TweetResponse
	err := t.db.QueryRowContext(context.Background(), query, in.UserId, in.Hashtag, in.Title, in.Content, in.ImageUrl).Scan(
		&res.Id, &res.UserId, &res.Hashtag, &res.Title, &res.Content, &res.ImageUrl, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (t *TweetRepo) UpdateTweet(in *pb.UpdateATweet) (*pb.TweetResponse, error) {
	query := `UPDATE tweets SET hashtag = $1, title = $2, content = $3, updated_at = $4 
	          WHERE id = $5 
	          RETURNING id, user_id, hashtag, title, content, image_url, created_at, updated_at`

	now := time.Now().Format(time.RFC3339)

	var res pb.TweetResponse
	err := t.db.QueryRowContext(context.Background(), query, in.Hashtag, in.Title, in.Content, now, in.Id).Scan(
		&res.Id, &res.UserId, &res.Hashtag, &res.Title, &res.Content, &res.ImageUrl, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (t *TweetRepo) AddImageToTweet(in *pb.Url) (*pb.Message, error) {
	query := `UPDATE tweets SET image_url = $1, updated_at = $2 WHERE id = $3`

	now := time.Now().Format(time.RFC3339)

	_, err := t.db.ExecContext(context.Background(), query, in.Url, now, in.TweetId)
	if err != nil {
		return nil, err
	}

	return &pb.Message{Message: "Image added successfully"}, nil
}

func (t *TweetRepo) UserTweets(in *pb.UserId) (*pb.Tweets, error) {
	query := `SELECT id, user_id, hashtag, title, content, image_url, created_at, updated_at 
	          FROM tweets WHERE user_id = $1`

	rows, err := t.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets pb.Tweets
	for rows.Next() {
		var tweet pb.TweetResponse
		if err := rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			return nil, err
		}
		tweets.Tweets = append(tweets.Tweets, &tweet)
	}

	return &tweets, nil
}

func (t *TweetRepo) GetTweet(in *pb.TweetId) (*pb.TweetResponse, error) {
	query := `SELECT id, user_id, hashtag, title, content, image_url, created_at, updated_at 
	          FROM tweets WHERE id = $1`

	var res pb.TweetResponse
	err := t.db.QueryRowContext(context.Background(), query, in.Id).Scan(
		&res.Id, &res.UserId, &res.Hashtag, &res.Title, &res.Content, &res.ImageUrl, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (t *TweetRepo) GetAllTweets(in *pb.TweetFilter) (*pb.Tweets, error) {
	query := `SELECT id, user_id, hashtag, title, content, image_url, created_at, updated_at 
	          FROM tweets WHERE hashtag LIKE '%' || $1 || '%' AND title LIKE '%' || $2 || '%' 
	          LIMIT $3 OFFSET $4`

	rows, err := t.db.QueryContext(context.Background(), query, in.Hashtag, in.Title, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets pb.Tweets
	for rows.Next() {
		var tweet pb.TweetResponse
		if err := rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			return nil, err
		}
		tweets.Tweets = append(tweets.Tweets, &tweet)
	}

	return &tweets, nil
}

func (t *TweetRepo) RecommendTweets(in *pb.UserId) (*pb.Tweets, error) {
	query := `SELECT id, user_id, hashtag, title, content, image_url, created_at, updated_at 
	          FROM tweets WHERE hashtag IN 
	          (SELECT hashtag FROM tweets WHERE user_id = $1) LIMIT 10`

	rows, err := t.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets pb.Tweets
	for rows.Next() {
		var tweet pb.TweetResponse
		if err := rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			return nil, err
		}
		tweets.Tweets = append(tweets.Tweets, &tweet)
	}

	return &tweets, nil
}

func (t *TweetRepo) GetNewTweets(in *pb.UserId) (*pb.Tweets, error) {
	query := `SELECT id, user_id, hashtag, title, content, image_url, created_at, updated_at 
	          FROM tweets WHERE user_id = $1 ORDER BY created_at DESC LIMIT 10`

	rows, err := t.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets pb.Tweets
	for rows.Next() {
		var tweet pb.TweetResponse
		if err := rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			return nil, err
		}
		tweets.Tweets = append(tweets.Tweets, &tweet)
	}

	return &tweets, nil
}
func (t *TweetRepo) AddReTweet(in *pb.ReTweetReq) (*pb.Message, error) {
	query := `UPDATE tweets SET tweet_id = $1, is_retweeted=true WHERE id = $2`

	err := t.db.QueryRowContext(context.Background(), query, in.RetweetId, in.TweetId)
	if err != nil {
		return nil, err.Err()
	}

	return &pb.Message{Message: "ReTweeted successfully"}, nil
}
