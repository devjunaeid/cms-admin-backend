package routes

import (
	"log"

	"github.com/devjunaeid/cms-admin-backend/models"
	"github.com/devjunaeid/cms-admin-backend/types"
	"github.com/devjunaeid/cms-admin-backend/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type AuthRoute struct {
	router fiber.Router
	db     *gorm.DB
}

func InitAuth(r fiber.Router, database *gorm.DB) *AuthRoute {
	return &AuthRoute{
		router: r,
		db:     database,
	}
}

func (ar *AuthRoute) CreateRoute() {
	ar.router.Post("/register", ar.registerUser)
	ar.router.Post("/login", ar.loginUser)
}

func (ar *AuthRoute) registerUser(c fiber.Ctx) error {
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

	// Check If the user already registerd.
	var user models.Users
	req := ar.db.First(&user, "email=?", payload.Email)
	if req.Error == nil {
		res := utils.CreateErrorRes("Failed to register user, already Registered!!", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	}

	// Insert new register request to the database.
	dbPayload := models.Users{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
	dbRes := ar.db.Create(&dbPayload)

	if dbRes.Error != nil {
		// Register Failed(DB error) Response.
		res := utils.CreateErrorRes("Failed to register user!", fiber.ErrInternalServerError.Code)
		return c.JSON(res)
	}

	// Register Successful Response.
	res := utils.CreateSuccessRes("Register Successfull", fiber.StatusCreated)
	return c.JSON(res)
}

func (ar *AuthRoute) loginUser(c fiber.Ctx) error {
	if c.Request().Body() == nil {
		res := utils.CreateErrorRes("Failed To log-in", fiber.ErrBadRequest.Code)
		c.JSON(res)
	}

	payload := new(types.LoginPayload)
	c.Bind().Body(payload)

	err := validator.Struct(payload)
	if err != nil {
		res := utils.CreateErrorRes("Failed to login", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	}

	var user types.DbLoginReqResponse
	dbQ := ar.db.Model(&models.Users{}).First(&user).Where("email=?", payload.Email)
	if dbQ.RowsAffected < 1 {
		res := utils.CreateErrorRes("Failed to login", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	} else {
		isPasswordCorrect := utils.CheckPasswordHash(payload.Password, user.Password)
		if !isPasswordCorrect {
			res := utils.CreateErrorRes("Failed to login", fiber.ErrBadRequest.Code)
			return c.JSON(res)
		} else {
			return c.JSON(user)
		}
	}
}
