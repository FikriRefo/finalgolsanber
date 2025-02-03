package entity

import "time"

type Profile struct {
	ID          int
	UserID      int
	Name        string
	Age         int
	PhoneNumber int
	AvatarUrl   *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
