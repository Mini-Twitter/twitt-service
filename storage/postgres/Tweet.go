package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	pb "twitt-service/genproto/tweet"
	pu "twitt-service/genproto/user"
)

type TweetRepo struct {
	db   *sqlx.DB
	user pu.UserServiceClient
}

func NewTweetRepo(db *sqlx.DB, user pu.UserServiceClient) *TweetRepo {
	return &TweetRepo{
		db:   db,
		user: user,
	}
}

func (t *TweetRepo) PostTweet(in *pb.Tweet) (*pb.TweetResponse, error) {
	query := `INSERT INTO tweets (id, user_id, hashtag, title, content, image_url, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, user_id, hashtag, title, content, image_url, created_at, updated_at`

	id := uuid.New().String()
	now := time.Now()

	var tweet pb.TweetResponse
	err := t.db.QueryRowContext(context.Background(), query,
		id, in.UserId, in.Hashtag, in.Title, in.Content, in.ImageUrl, now, now).Scan(
		&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to insert tweet: %v", err)
	}

	return &tweet, nil
}

func (t *TweetRepo) UpdateTweet(in *pb.UpdateATweet) (*pb.TweetResponse, error) {
	query := `UPDATE tweets SET hashtag = $1, title = $2, content = $3, updated_at = $4 
	          WHERE id = $5 RETURNING id, user_id, hashtag, title, content, image_url, created_at, updated_at`

	now := time.Now()

	var tweet pb.TweetResponse
	err := t.db.QueryRowContext(context.Background(), query,
		in.Hashtag, in.Title, in.Content, now, in.Id).Scan(
		&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to update tweet: %v", err)
	}

	return &tweet, nil
}

func (t *TweetRepo) AddImageToTweet(in *pb.Url) (*pb.Message, error) {
	query := `UPDATE tweets SET image_url = $1, updated_at = $2 WHERE id = $3`

	now := time.Now()

	_, err := t.db.ExecContext(context.Background(), query, in.Url, now, in.TweetId)
	if err != nil {
		return nil, fmt.Errorf("failed to add image to tweet: %v", err)
	}

	return &pb.Message{Message: "Image added successfully"}, nil
}

func (t *TweetRepo) UserTweets(in *pb.UserId) (*pb.Tweets, error) {
	query := `SELECT id, user_id, hashtag, title, content, image_url, created_at, updated_at 
	          FROM tweets WHERE user_id = $1`

	rows, err := t.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user tweets: %v", err)
	}
	defer rows.Close()

	var tweets pb.Tweets
	for rows.Next() {
		var tweet pb.TweetResponse
		if err := rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan tweet: %v", err)
		}
		tweets.Tweets = append(tweets.Tweets, &tweet)
	}

	return &tweets, nil
}

func (t *TweetRepo) GetTweet(in *pb.TweetId) (*pb.TweetResponse, error) {
	query := `SELECT id, user_id, hashtag, title, content, image_url, created_at, updated_at 
	          FROM tweets WHERE id = $1`

	var tweet pb.TweetResponse
	err := t.db.QueryRowContext(context.Background(), query, in.Id).Scan(
		&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to get tweet: %v", err)
	}

	return &tweet, nil
}

func (t *TweetRepo) GetAllTweets(in *pb.TweetFilter) (*pb.Tweets, error) {
	query := `SELECT id, user_id, hashtag, title, content, image_url, created_at, updated_at 
	          FROM tweets WHERE (hashtag = $1 OR $1 = '') AND (title = $2 OR $2 = '')
	          ORDER BY created_at DESC LIMIT $3 OFFSET $4`

	rows, err := t.db.QueryContext(context.Background(), query, in.Hashtag, in.Title, in.Limit, in.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tweets: %v", err)
	}
	defer rows.Close()

	var tweets pb.Tweets
	for rows.Next() {
		var tweet pb.TweetResponse
		if err := rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan tweet: %v", err)
		}
		tweets.Tweets = append(tweets.Tweets, &tweet)
	}

	return &tweets, nil
}

func (t *TweetRepo) RecommendTweets(in *pb.UserId) (*pb.Tweets, error) {
	// Assuming "recommendation" logic is based on hashtags the user has used or interacted with
	query := `SELECT t.id, t.user_id, t.hashtag, t.title, t.content, t.image_url, t.created_at, t.updated_at
	          FROM tweets t
	          JOIN follows f ON t.user_id = f.following_id
	          WHERE f.follower_id = $1
	          ORDER BY t.created_at DESC`

	rows, err := t.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get recommended tweets: %v", err)
	}
	defer rows.Close()

	var tweets pb.Tweets
	for rows.Next() {
		var tweet pb.TweetResponse
		if err := rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan tweet: %v", err)
		}
		tweets.Tweets = append(tweets.Tweets, &tweet)
	}

	return &tweets, nil
}

func (t *TweetRepo) GetNewTweets(in *pb.UserId) (*pb.Tweets, error) {
	query := `SELECT id, user_id, hashtag, title, content, image_url, created_at, updated_at 
	          FROM tweets WHERE created_at > (
	              SELECT last_login FROM users WHERE id = $1
	          )
	          ORDER BY created_at DESC`

	rows, err := t.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get new tweets: %v", err)
	}
	defer rows.Close()

	var tweets pb.Tweets
	for rows.Next() {
		var tweet pb.TweetResponse
		if err := rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Hashtag, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan tweet: %v", err)
		}
		tweets.Tweets = append(tweets.Tweets, &tweet)
	}

	return &tweets, nil
}
