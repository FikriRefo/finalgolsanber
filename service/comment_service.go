package service

import (
	"socmed/dto"
	"socmed/entity"
	"socmed/errorhandle"
	"socmed/repository"
)

type CommentService interface {
	Create(req *dto.CommentRequest) error
	GetAllByPostID(postID int) ([]dto.CommentResponse, error)
	Update(commentID int, req *dto.CommentRequest) error
	Delete(commentID int) error
}

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) *commentService {
	return &commentService{
		repository: r,
	}
}

func (s *commentService) Create(req *dto.CommentRequest) error {
	comment := entity.Comment{
		PostID:  req.PostID,
		UserID:  req.UserID,
		Comment: req.Comment,
	}

	if err := s.repository.Create(&comment); err != nil {
		return &errorhandle.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *commentService) GetAllByPostID(postID int) ([]dto.CommentResponse, error) {
	comments, err := s.repository.FindByPostID(postID)
	if err != nil {
		return nil, &errorhandle.InternalServerError{Message: err.Error()}
	}

	var commentResponses []dto.CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, dto.CommentResponse{
			ID:        comment.ID,
			UserID:    comment.UserID,
			PostID:    comment.PostID,
			Comment:   comment.Comment,
			CreatedAt: comment.CreatedAt.String(),
			UpdatedAt: comment.UpdatedAt.String(),
		})
	}

	return commentResponses, nil
}

func (s *commentService) Update(commentID int, req *dto.CommentRequest) error {
	existingComment, err := s.repository.FindByID(commentID)
	if err != nil {
		return &errorhandle.NotFoundError{Message: "Comment not found"}
	}

	existingComment.Comment = req.Comment

	if err := s.repository.Update(existingComment); err != nil {
		return &errorhandle.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *commentService) Delete(commentID int) error {
	if err := s.repository.Delete(commentID); err != nil {
		return &errorhandle.InternalServerError{Message: err.Error()}
	}

	return nil
}
