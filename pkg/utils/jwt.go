package utils

import (
	"chatapp/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(user model.User) (string, error) {
	expirationTime := time.Now().Add(8 * time.Hour)
	claims := model.JWT{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
