package dto

type CommentResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID" json:"user"`
	PostID    int    `json:"post_id"`
	Post      Post   `gorm:"foreignKey:PostID" json:"post"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommentRequest struct {
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
	Comment string `json:"comment"`
}

type Post struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	User       User   `gorm:"foreignKey:UserID" json:"user"`
	Tweet      string `json:"tweet"`
	PictureUrl string `json:"picture_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
