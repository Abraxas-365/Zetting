package user_ports

import "github.com/gofiber/fiber/v2"

type UserController interface {
	CreateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	CheckEmailExist(c *fiber.Ctx) error
	GetUsersByProfession(c *fiber.Ctx) error
}
