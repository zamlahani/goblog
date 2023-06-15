package routes

import (
	"goblog/controllers"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App) {
	app.Post("/posts", controllers.CreatePost)
	//All routes related to users comes here
	app.Get("/posts", controllers.GetAllPosts)

	app.Get("/fib", controllers.Fib)
}
