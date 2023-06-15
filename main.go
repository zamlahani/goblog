package main

import (
	"goblog/configs"
	"goblog/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.PostRoute(app)

	app.Listen(":6000")
}
