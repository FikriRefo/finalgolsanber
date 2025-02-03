package service

import (
	"socmed/dto"
	"socmed/entity"
	"socmed/errorhandle"
	"socmed/repository"
)

type PostService interface {
	Create(req *dto.PostRequest) error
	GetAllByUserID(userID int) ([]dto.PostResponse, error)
	Update(postID int, req *dto.PostRequest) error
	Delete(postID int, userID int) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *dto.PostRequest) error {
	post := entity.Post{
		UserID: req.UserID,
		Tweet:  req.Tweet,
	}

	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &errorhandle.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *postService) GetAllByUserID(userID int) ([]dto.PostResponse, error) {
	posts, err := s.repository.FindByUserID(userID)
	if err != nil {
		return nil, &errorhandle.InternalServerError{Message: err.Error()}
	}

	var postResponses []dto.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, dto.PostResponse{
			ID:         post.ID,
			UserID:     post.UserID,
			Tweet:      post.Tweet,
			PictureUrl: *post.PictureUrl,
			CreatedAt:  post.CreatedAt.String(),
			UpdatedAt:  post.UpdatedAt.String(),
		})
	}

	return postResponses, nil
}

func (s *postService) Update(postID int, req *dto.PostRequest) error {
	existingPost, err := s.repository.FindByID(postID)
	if err != nil {
		return &errorhandle.NotFoundError{Message: "Post not found"}
	}

	existingPost.Tweet = req.Tweet
	if req.Picture != nil {
		existingPost.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Update(existingPost); err != nil {
		return &errorhandle.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *postService) Delete(postID int, userID int) error {
	existingPost, err := s.repository.FindByID(postID)
	if err != nil {
		return &errorhandle.NotFoundError{Message: "Post not found"}
	}

	if existingPost.UserID != userID {
		return &errorhandle.UnauthorizedError{Message: "Unauthorized"}
	}

	if err := s.repository.Delete(postID); err != nil {
		return &errorhandle.InternalServerError{Message: err.Error()}
	}

	return nil
}
