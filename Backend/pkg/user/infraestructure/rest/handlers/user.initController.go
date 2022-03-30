package user_handlers

import (
	ports "zetting/pkg/user/core/ports"
)

import "github.com/gofiber/fiber/v2"

type UserController interface {
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	CheckEmailExist(c *fiber.Ctx) error
	GetUsersByProfession(c *fiber.Ctx) error
	UploadProfileImage(c *fiber.Ctx) error
}
type userController struct {
	userService ports.UserService
}

func NewUserController(userService ports.UserService) UserController {
	return &userController{
		userService,
	}
}
