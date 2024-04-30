package api

import (
	"github.com/devjunaeid/cms-admin-backend/routes"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type greeting struct {
	Message  string `json:"message"`
	Greeting string `json:"greeting"`
}

func CreateAPI(db *gorm.DB) {
	app := fiber.New(fiber.Config{
		AppName: "Synafeia CMS Backend",
	})

	app.Get("/", func(c fiber.Ctx) error {
		c.Accepts("application/json")
		data := greeting{
			Message:  "This for admin personel only!!",
			Greeting: "Welcome to Synafeia-backend!",
		}
		return c.JSON(data)
	})
	// Route Group.
	userHandler := app.Group("/v1")
	userHandler.Get("/", func(c fiber.Ctx) error {
		return c.SendString("API Version 1.0!!!")
	})

	// Auth Routes.
	authRoute := routes.InitAuth(userHandler, db)
	authRoute.CreateRoute()

	// User Route
	userRoute := routes.InitUser(userHandler, db)
	userRoute.CreateRoute()

	app.Listen(":3000")
}
