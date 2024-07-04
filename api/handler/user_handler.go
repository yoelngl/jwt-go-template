package handler

import (
	"chatapp/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	response := service.GetUser(c)

	return c.JSON(response)
}

func SearchUser(c *fiber.Ctx) error {
	response := service.SearchUser(c)

	return c.JSON(response)
}
