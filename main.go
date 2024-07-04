package main

import (
	"chatapp/api/router"
	"chatapp/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	app := fiber.New()

	router.SetupRouter(app)

	app.Listen(":8000")
}
