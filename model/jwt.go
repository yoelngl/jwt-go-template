package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
