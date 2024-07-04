package router

import (
	"chatapp/api/handler"
	"chatapp/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)

	api.Use(middleware.AuthMiddleware)
	api.Get("/users", handler.SearchUser)
	api.Get("/users", handler.GetUser)
}
