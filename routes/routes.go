package routes

import (
	"github.com/Ignacio07/micro-service-admin/controllers"
	"github.com/Ignacio07/micro-service-admin/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//Home route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	//Token validation Route (middleware example)
	app.Get("/validate", middleware.Validate)

	//auth routes
	app.Post("/login", controllers.Login)
	app.Post("/signup", controllers.SingUp)

	//user routes
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:userId", controllers.GetUser)
	app.Post("/users", controllers.CreateUser)
	app.Put("/users/:userId", controllers.UpdateUser)
	app.Delete("/users/:userId", controllers.DeleteUser)

	//test
	//app.Test("/testing",tests.TestCreateUser)

}
