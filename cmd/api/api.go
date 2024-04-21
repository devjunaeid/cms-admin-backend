package api

import (
	"github.com/devjunaeid/cms-admin-backend/routes"
	"github.com/gofiber/fiber/v3"
)

type greeting struct {
	Message  string `json:"message"`
	Greeting string `json:"greeting"`
}

func CreateAPI() {
	app := fiber.New(fiber.Config{
		AppName: "Synafeia CMS Backend",
	})

	app.Get("/", func(c fiber.Ctx) error {
		c.Accepts("application/json")
		data := greeting{
			Message:  "Hello There",
			Greeting: "Welcome to Fiber!",
		}
		return c.JSON(data)
	})
	// User Route Group.
	userHandler := app.Group("/user")
	routes.CeateRoute(userHandler)
	app.Listen(":3000")
}
