package model

import "time"

type User struct {
	ID          int           `gorm:"primary_key" json:"id"`
	Transaction []Transaction `gorm:"ForeignKey:UserID;" json:"-"`
	FullName    string        `json:"full_name" validate:"required"`
	Mobile      string        `json:"mobile" validate:"required"`
	Email       string        `json:"email" validate:"required,email"`
	Password    string        `json:"password" validate:"required"`
	Address     string        `json:"address" validate:"required"`
	IsAdmin     bool          `json:"is_admin,omitempty"`
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserResponse struct {
	UserName string
	Mobile   string
	Email    string
	Token    string
}
