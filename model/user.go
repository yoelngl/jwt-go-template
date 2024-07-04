package model

import "time"

type User struct {
	ID        string    `json:"id" `
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,gte=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDetail struct {
	ID       string `json:"id" `
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SearchUserResult struct {
	Username string `json:"username" `
	Email    string `json:"email"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

func (User) TableName() string {
	return "users"
}


