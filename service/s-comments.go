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
	res, err := s.storage.PostComment(in)
	if err != nil {
		s.logger.Error("failed to post comment", err)
		return nil, err
	}
	return res, nil
}

func (s *commentsService) UpdateComment(in *pb.UpdateAComment) (*pb.CommentRes, error) {
	res, err := s.storage.UpdateComment(in)
	if err != nil {
		s.logger.Error("failed to update comment", err)
		return nil, err
	}
	return res, nil
}

func (s *commentsService) DeleteComment(in *pb.CommentId) (*pb.Message, error) {
	res, err := s.storage.DeleteComment(in)
	if err != nil {
		s.logger.Error("failed to delete comment", err)
		return nil, err
	}
	return res, nil
}

func (s *commentsService) GetComment(in *pb.CommentId) (*pb.Comment, error) {
	res, err := s.storage.GetComment(in)
	if err != nil {
		s.logger.Error("failed to get comment", err)
		return nil, err
	}
	return res, nil
}

func (s *commentsService) GetAllComments(in *pb.CommentFilter) (*pb.Comments, error) {
	res, err := s.storage.GetAllComments(in)
	if err != nil {
		s.logger.Error("failed to get comments", err)
		return nil, err
	}
	return res, nil
}

func (s *commentsService) GetUserComments(in *pb.UserId) (*pb.Comments, error) {
	res, err := s.storage.GetUserComments(in)
	if err != nil {
		s.logger.Error("failed to get comments", err)
		return nil, err
	}
	return res, nil
}

func (s *commentsService) AddLikeToComment(in *pb.UserId) (*pb.Message, error) {
	res, err := s.storage.AddLikeToComment(in)
	if err != nil {
		s.logger.Error("failed to add like to comment", err)
		return nil, err
	}
	return res, nil
}

func (s *commentsService) DeleteLikeComment(in *pb.UserId) (*pb.Message, error) {
	res, err := s.storage.DeleteLikeComment(in)
	if err != nil {
		s.logger.Error("failed to delete like comment", err)
		return nil, err
	}
	return res, nil
}
