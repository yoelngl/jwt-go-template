package service

import (
	"chatapp/config"
	"chatapp/internal/repository"
	"chatapp/model"
	"chatapp/pkg/constant"
	"chatapp/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) model.Response {
	var user model.User
	var err error

	if err = c.BodyParser(&user); err != nil {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.INVALID_INPUT, nil)
	}

	if err = config.Validate.Struct(&user); err != nil {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.BAD_PARAMATER, nil)
	}

	uniqueEmail := repository.FindUser(user)
	if uniqueEmail.Email != "" {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.EMAIL_ALREADY_EXISTS, nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.FAILED_HASH_PASSWORD, nil)
	}

	user.Password = string(hashedPassword)

	err = repository.CreateNewUser(&user)
	if err != nil {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.FAILED_TO_CREATE_USER, nil)
	}

	return utils.Response(constant.OK_CODE, constant.SUCESSFULLY_CREATED, nil)
}

func LoginUser(c *fiber.Ctx) model.Response {
	var input model.Login
	var err error

	if err = c.BodyParser(&input); err != nil {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.INVALID_INPUT, nil)
	}

	if err = config.Validate.Struct(&input); err != nil {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.BAD_PARAMATER, nil)
	}

	request := model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	user := repository.FindUser(request)
	if user.ID == "" {
		return utils.Response(constant.NOT_FOUND_CODE, constant.ACCOUNT_NOT_FOUND, nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.PASSWORD_MISMATCH, nil)
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return utils.Response(constant.BAD_REQUEST_CODE, constant.FAILED_TO_GENERATE_TOKEN, nil)
	}

	return utils.Response(constant.OK_CODE, constant.SUCESSFULLY_LOGIN, token)
}
