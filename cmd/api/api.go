package api

import (
	"github.com/gofiber/fiber/v3"
)

func CreateAPI() {
	app := fiber.New(fiber.Config{
		AppName: "Synafeia CMS Backend",
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, Juaneid!")
	})

	app.Listen(":3000")
}
