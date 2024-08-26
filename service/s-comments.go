package service

import (
	"log/slog"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type CommentsService interface {
	PostComment(in *pb.Comment) (*pb.CommentRes, error)
	UpdateComment(in *pb.UpdateAComment) (*pb.CommentRes, error)
	DeleteComment(in *pb.CommentId) (*pb.Message, error)
	GetComment(in *pb.CommentId) (*pb.Comment, error)
	GetAllComments(in *pb.CommentFilter) (*pb.Comments, error)
	GetUserComments(in *pb.UserId) (*pb.Comments, error)
	AddLikeToComment(in *pb.UserId) (*pb.Message, error)
	DeleteLikeComment(in *pb.UserId) (*pb.Message, error)
}

type commentsService struct {
	storage storage.CommentsStorage
	logger  *slog.Logger
}

func NewCommentsService(st storage.CommentsStorage, logger *slog.Logger) CommentsService {
	return &commentsService{
		storage: st,
		logger:  logger,
	}
}

func (s *commentsService) PostComment(in *pb.Comment) (*pb.CommentRes, error) {
	return &pb.CommentRes{}, nil
}

func (s *commentsService) UpdateComment(in *pb.UpdateAComment) (*pb.CommentRes, error) {
	return &pb.CommentRes{}, nil
}

func (s *commentsService) DeleteComment(in *pb.CommentId) (*pb.Message, error) {
	return &pb.Message{}, nil
}

func (s *commentsService) GetComment(in *pb.CommentId) (*pb.Comment, error) {
	return &pb.Comment{}, nil
}

func (s *commentsService) GetAllComments(in *pb.CommentFilter) (*pb.Comments, error) {
	return &pb.Comments{}, nil
}

func (s *commentsService) GetUserComments(in *pb.UserId) (*pb.Comments, error) {
	return &pb.Comments{}, nil
}

func (s *commentsService) AddLikeToComment(in *pb.UserId) (*pb.Message, error) {
	return &pb.Message{}, nil
}

func (s *commentsService) DeleteLikeComment(in *pb.UserId) (*pb.Message, error) {
	return &pb.Message{}, nil
}
