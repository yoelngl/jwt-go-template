package repository

import (
	"chatapp/config"
	"chatapp/model"

	"github.com/google/uuid"
)

func CreateNewUser(user *model.User) error {
	user.ID = uuid.New().String()
	err := config.DB.Create(&user)
	if err != nil {
		return err.Error
	}

	return nil
}

func FindUser(user model.User) model.User {
	var users model.User
	result := config.DB.Where("email = ?", user.Email).First(&users)
	if result.Error != nil {
		return model.User{}
	}

	return users
}