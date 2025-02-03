package entity

import "time"

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
