package dto

import "mime/multipart"

type ProfileResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"user"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber int    `json:"phone_number"`
	AvatarUrl   string `json:"avatar_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProfileRequest struct {
	UserID      int                   `form:"user_id"`
	Name        string                `form:"name"`
	Age         int                   `form:"age"`
	PhoneNumber int                   `form:"phone_number"`
	Avatar      *multipart.FileHeader `form:"avatar"`
	AvatarUrl   string                `json:"avatar_url"` // Add AvatarUrl field
}
