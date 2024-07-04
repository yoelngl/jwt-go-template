package service

import (
	"chatapp/internal/repository"
	"chatapp/model"
	"chatapp/pkg/constant"
	"chatapp/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) model.Response {
	username := c.Locals("username").(string)

	result := repository.GetUsersDetail(username)
	if result.ID == "" {
		return utils.Response(constant.NOT_FOUND_CODE, constant.ACCOUNT_NOT_FOUND, nil)
	}

	users := model.UserDetail{
		ID:       result.ID,
		Username: result.Username,
		Email:    result.Email,
	}

	return utils.Response(constant.OK_CODE, constant.ACCOUNT_ACTIVE, users)
}

func SearchUser(c *fiber.Ctx) model.Response {
	username := c.FormValue("username")
	if username == "" {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.BAD_REQUEST, nil)
	}

	result := repository.SearchUserByUsername(username)
	if len(result) <= 0 {
		return utils.Response(constant.NOT_FOUND_CODE, constant.SEARCH_ACCOUNT_DOESNT_EXISTS, nil)
	}

	return utils.Response(constant.OK_CODE, constant.ACCOUNT_ACTIVE, result)
}
