package handler

import (
	"chatapp/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	response := service.RegisterUser(c)

	return c.JSON(response)
}

func Login(c *fiber.Ctx) error {
	response := service.LoginUser(c)

	return c.JSON(response)
}
