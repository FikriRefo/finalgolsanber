package repository

import (
	"socmed/entity"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *entity.Comment) error
	FindByPostID(postID int) ([]entity.Comment, error)
	FindByID(commentID int) (*entity.Comment, error)
	Update(comment *entity.Comment) error
	Delete(commentID int) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(comment *entity.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) FindByPostID(postID int) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}

func (r *commentRepository) FindByID(commentID int) (*entity.Comment, error) {
	var comment entity.Comment
	err := r.db.First(&comment, commentID).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) Update(comment *entity.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) Delete(commentID int) error {
	return r.db.Delete(&entity.Comment{}, commentID).Error
}
