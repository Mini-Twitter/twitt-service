package postgres

//
//import (
//	"context"
//	"database/sql"
//	"fmt"
//
//	"github.com/jmoiron/sqlx"
//	pb "twitt-service/genproto/tweet"
//)
//
//type SubscribeRepo struct {
//	db *sqlx.DB
//}
//
//func NewSubscribeRepo(db *sqlx.DB) *SubscribeRepo {
//	return &SubscribeRepo{
//		db: db,
//	}
//}
//
//func (s *SubscribeRepo) Follow(in *pb.FollowReq) (*pb.FollowRes, error) {
//	query := `INSERT INTO follows (follower_id, followed_id, created_at)
//	          VALUES ($1, $2, NOW())
//	          RETURNING follower_id, followed_id, created_at`
//
//	var res pb.FollowRes
//	err := s.db.QueryRowContext(context.Background(), query, in.FollowerId, in.FollowedId).Scan(
//		&res.FollowerId, &res.FollowedId, &res.CreatedAt)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &res, nil
//}
//
//func (s *SubscribeRepo) Unfollow(in *pb.FollowReq) (*pb.DFollowRes, error) {
//	query := `DELETE FROM follows WHERE follower_id = $1 AND followed_id = $2
//	          RETURNING follower_id, followed_id`
//
//	var res pb.DFollowRes
//	err := s.db.QueryRowContext(context.Background(), query, in.FollowerId, in.FollowedId).Scan(
//		&res.FollowerId, &res.FollowedId)
//
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, fmt.Errorf("no such follow relation exists")
//		}
//		return nil, err
//	}
//
//	return &res, nil
//}
//
//func (s *SubscribeRepo) GetUserFollowers(in *pb.UserId) (*pb.Count, error) {
//	query := `SELECT COUNT(*) FROM follows WHERE followed_id = $1`
//
//	var count pb.Count
//	err := s.db.QueryRowContext(context.Background(), query, in.Id).Scan(&count.Count)
//	if err != nil {
//		return nil, err
//	}
//
//	return &count, nil
//}
//
//func (s *SubscribeRepo) GetUserFollows(in *pb.UserId) (*pb.Count, error) {
//	query := `SELECT COUNT(*) FROM follows WHERE follower_id = $1`
//
//	var count pb.Count
//	err := s.db.QueryRowContext(context.Background(), query, in.Id).Scan(&count.Count)
//	if err != nil {
//		return nil, err
//	}
//
//	return &count, nil
//}
//
//func (s *SubscribeRepo) MostPopularUser(in *pb.Void) (*pb.User, error) {
//	query := `SELECT followed_id, COUNT(followed_id) AS followers
//	          FROM follows
//	          GROUP BY followed_id
//	          ORDER BY followers DESC
//	          LIMIT 1`
//
//	var user pb.User
//	err := s.db.QueryRowContext(context.Background(), query).Scan(&user.Id, &user.FollowersCount)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, nil
//		}
//		return nil, err
//	}
//
//	return &user, nil
//}
