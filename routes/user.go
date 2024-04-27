package routes

import (
	"log"

	"github.com/devjunaeid/cms-admin-backend/models"
	"github.com/devjunaeid/cms-admin-backend/types"
	"github.com/devjunaeid/cms-admin-backend/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var validator = utils.InitValidator()

type UserRoute struct {
	route fiber.Router
	db    *gorm.DB
}

// Init User Route.
func InitUser(r fiber.Router, db *gorm.DB) *UserRoute {
	return &UserRoute{
		route: r,
		db:    db,
	}
}

func (ur *UserRoute) CreateRoute() {
	// User Table Migration.
	ur.db.AutoMigrate(&models.User{})

	// Register User.
	ur.route.Get("/", func(c fiber.Ctx) error {
		return c.SendString("From User Routing Group")
	})
	ur.route.Post("/register", ur.registerUser)
}

func (ur *UserRoute) registerUser(c fiber.Ctx) error {
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
