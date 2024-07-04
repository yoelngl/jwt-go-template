package repository

import (
	"chatapp/config"
	"chatapp/model"
)

func GetUsersDetail(username string) model.User {
	var users model.User
	result := config.DB.Where("username = ?", username).First(&users)
	if result.Error != nil {
		return model.User{}
	}

	return users
}

func SearchUserByUsername(username string) []model.SearchUserResult {
	var users []model.SearchUserResult
	result := config.DB.Table("users").Where("username LIKE ?", "%"+username+"%").Scan(&users)
	if result.Error != nil {
		return []model.SearchUserResult{}
	}

	return users
}
