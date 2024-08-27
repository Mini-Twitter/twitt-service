package postgres

import (
	"context"
	"fmt"
	"time"
	pu "twitt-service/genproto/user"
	"twitt-service/storage"

	"github.com/jmoiron/sqlx"
	pb "twitt-service/genproto/tweet"
)

type SubscribeRepo struct {
	db   *sqlx.DB
	user pu.UserServiceClient
}

func NewSubscribeRepo(db *sqlx.DB, user pu.UserServiceClient) storage.SubscribeStorage {
	return &SubscribeRepo{
		db:   db,
		user: user,
	}
}

func (s *SubscribeRepo) Follow(in *pb.FollowReq) (*pb.FollowRes, error) {
	query := `INSERT INTO follows (follower_id, following_id, followed_at) 
	          VALUES ($1, $2, $3) RETURNING follower_id, following_id, followed_at`

	now := time.Now()

	var res pb.FollowRes
	err := s.db.QueryRowContext(context.Background(), query, in.FollowerId, in.FollowingId, now).Scan(
		&res.FollowerId, &res.FollowingId, &res.FollowedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to follow user: %v", err)
	}

	return &res, nil
}

func (s *SubscribeRepo) Unfollow(in *pb.FollowReq) (*pb.DFollowRes, error) {
	query := `DELETE FROM follows WHERE follower_id = $1 AND following_id = $2 
	          RETURNING follower_id, following_id, NOW()`

	var res pb.DFollowRes
	err := s.db.QueryRowContext(context.Background(), query, in.FollowerId, in.FollowingId).Scan(
		&res.FollowerId, &res.FollowingId, &res.UnfollowedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to unfollow user: %v", err)
	}

	return &res, nil
}

func (s *SubscribeRepo) GetUserFollowers(in *pb.UserId) (*pb.Count, error) {
	query := `SELECT COUNT(*) FROM follows WHERE following_id = $1`

	var count pb.Count
	err := s.db.QueryRowContext(context.Background(), query, in.Id).Scan(&count.Count)

	if err != nil {
		return nil, fmt.Errorf("failed to get user followers count: %v", err)
	}

	count.Description = "Number of followers"
	return &count, nil
}

func (s *SubscribeRepo) GetUserFollows(in *pb.UserId) (*pb.Count, error) {
	query := `SELECT COUNT(*) FROM follows WHERE follower_id = $1`

	var count pb.Count
	err := s.db.QueryRowContext(context.Background(), query, in.Id).Scan(&count.Count)

	if err != nil {
		return nil, fmt.Errorf("failed to get user follows count: %v", err)
	}

	count.Description = "Number of follows"
	return &count, nil
}

func (s *SubscribeRepo) MostPopularUser(in *pb.Void) (*pb.User, error) {
	query := `SELECT u.id, u.first_name, p.username, p.bio, p.profile_image, 
	                 COUNT(f.follower_id) AS followers_count,
	                 (SELECT COUNT(*) FROM follows WHERE follower_id = u.id) AS following_count,
	                 (SELECT COUNT(*) FROM tweets WHERE user_id = u.id) AS posts_count,
	                 u.created_at, u.updated_at
	          FROM users u
	          JOIN user_profile p ON u.id = p.user_id
	          JOIN follows f ON u.id = f.following_id
	          GROUP BY u.id, p.username, p.bio, p.profile_image
	          ORDER BY followers_count DESC
	          LIMIT 1`

	var user pb.User
	err := s.db.QueryRowContext(context.Background(), query).Scan(
		&user.UserId, &user.FirstName, &user.Username, &user.Bio, &user.ProfileImage,
		&user.FollowersCount, &user.FollowingCount, &user.PostsCount, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to get most popular user: %v", err)
	}

	return &user, nil
}
