package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type CommentRepo struct {
	db *sqlx.DB
}

func NewCommentRepo(db *sqlx.DB) storage.CommentsStorage {
	return &CommentRepo{
		db: db,
	}
}

func (c *CommentRepo) PostComment(in *pb.Comment) (*pb.CommentRes, error) {
	return &pb.CommentRes{}, nil
}

func (c *CommentRepo) UpdateComment(in *pb.UpdateAComment) (*pb.CommentRes, error) {
	return &pb.CommentRes{}, nil
}

func (c *CommentRepo) DeleteComment(in *pb.CommentId) (*pb.Message, error) {
	return &pb.Message{}, nil
}

func (c *CommentRepo) GetComment(in *pb.CommentId) (*pb.Comment, error) {
	return &pb.Comment{}, nil
}

func (c *CommentRepo) GetAllComments(in *pb.CommentFilter) (*pb.Comments, error) {
	return &pb.Comments{}, nil
}

func (c *CommentRepo) GetUserComments(in *pb.UserId) (*pb.Comments, error) {
	return &pb.Comments{}, nil
}

func (c *CommentRepo) AddLikeToComment(in *pb.UserId) (*pb.Message, error) {
	return &pb.Message{}, nil
}

func (c *CommentRepo) DeleteLikeComment(in *pb.UserId) (*pb.Message, error) {
	return &pb.Message{}, nil
}
