package repository

import (
	"socmed/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindByUserID(userID int) ([]entity.Post, error)
	Create(post *entity.Post) error
	GetByUserID(userID int) ([]entity.Post, error)
	FindByID(postID int) (*entity.Post, error) // <-- Add this line
	Update(post *entity.Post) error
	Delete(postID int) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) FindByID(postID int) (*entity.Post, error) {
	var post entity.Post
	err := r.db.Where("id = ?", postID).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Create(post *entity.Post) error {
	return r.db.Create(&post).Error
}

func (r *postRepository) GetByUserID(userID int) ([]entity.Post, error) {
	var posts []entity.Post
	err := r.db.Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

func (r *postRepository) FindByUserID(userID int) ([]entity.Post, error) {
	var posts []entity.Post
	err := r.db.Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) Update(post *entity.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(postID int) error {
	err := r.db.Where("id = ?", postID).Delete(&entity.Post{}).Error
	if err != nil {
		return err
	}
	return nil
}
