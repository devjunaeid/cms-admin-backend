package routes

import (
	"log"

	"github.com/devjunaeid/cms-admin-backend/types"
	"github.com/devjunaeid/cms-admin-backend/utils"
	"github.com/gofiber/fiber/v3"
)

var validator = utils.InitValidator()

func CeateRoute(route fiber.Router) {
	// Register User.
	route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("From User Routing Group")
	})
	route.Post("/register", registerUser)
}

var registerUser = func(c fiber.Ctx) error {
	// Check if the body is empty
	if c.Request().Body() == nil {
		res := utils.CreateErrorRes("Bad Request", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	}
	payload := new(types.RegisterPayload)
	c.Bind().Body(payload)
	err := validator.Struct(payload)

	if err != nil {
		log.Printf("Error registering user, error: %v", err.Error())
		res := utils.CreateErrorRes("Failed to register user", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	}
	return c.JSON(payload)
}
