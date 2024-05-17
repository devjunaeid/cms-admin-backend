package routes

import (
	"fmt"
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
	ur.db.AutoMigrate(&models.Users{})

	// Get User.
	ur.route.Get("/users", ur.getUsers)
	ur.route.Get("/user/:id", ur.getUser)

	// Update User Role.
	ur.route.Patch("/user/:id", ur.patchUsers)
	// Delete User.
	ur.route.Delete("/user/:id", ur.deleteUser)
}

// Get All users.
func (ur *UserRoute) getUsers(c fiber.Ctx) error {
	// Get All user from database.
	var users []types.UserResponse
	dbRes := ur.db.Model(&models.Users{}).Find(&users)

	if dbRes.Error != nil {
		res := utils.CreateErrorRes("Failed to get users!!", fiber.ErrInternalServerError.Code)
		return c.JSON(res)
	}

	return c.JSON(users)
}

// Get Single User.
func (ur *UserRoute) getUser(c fiber.Ctx) error {
	userID := c.Params("id")

	var user types.UserResponse
	dbRes := ur.db.Model(&models.Users{}).Where("id=?", userID).First(&user)

	if dbRes.Error != nil {
		res := utils.CreateErrorRes("No user found!", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	}

	return c.JSON(user)
}

// Delete User.
func (ur *UserRoute) deleteUser(c fiber.Ctx) error {
	userID := c.Params("id")
	dbRes := ur.db.Where("id=?", userID).Delete(&models.Users{})
	fmt.Println(dbRes.Error)
	if dbRes.RowsAffected < 1 {
		res := utils.CreateErrorRes("No record found to delete!!", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	} else if dbRes.Error != nil {
		res := utils.CreateErrorRes("Faild to delete!!", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	}

	res := utils.CreateSuccessRes("Deleted!!", fiber.StatusAccepted)
	return c.JSON(res)
}

// Patch User.
func (ur *UserRoute) patchUsers(c fiber.Ctx) error {
	userID := c.Params("id")
	if c.Request().Body() == nil {
		res := utils.CreateErrorRes("Bad Request", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	}

	payload := new(types.UpdateUserRolePayload)
	c.Bind().Body(payload)
	err := validator.Struct(payload)
	if err != nil {
		log.Printf("Error updating user, error: %v", err.Error())
		res := utils.CreateErrorRes("Failed to update user", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	}

	// Check If the user already registerd.
	var user models.Users
	dbRes := ur.db.Model(&user).Where("id=?", userID).Update("role", payload.Role)
	if dbRes.RowsAffected < 1 {
		res := utils.CreateErrorRes("No record found to update!!", fiber.ErrBadRequest.Code)
		return c.JSON(res)
	} else if dbRes.Error != nil {
		res := utils.CreateErrorRes("Faild to update!!", fiber.ErrInternalServerError.Code)
		return c.JSON(res)
	}

	res := utils.CreateSuccessRes("Role updated!", fiber.StatusCreated)
	return c.JSON(res)
}
