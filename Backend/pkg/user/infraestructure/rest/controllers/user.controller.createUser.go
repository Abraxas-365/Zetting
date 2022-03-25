package user_controllers

import (
	"fmt"
	"zetting/auth"
	models "zetting/pkg/user/core/models"

	"github.com/gofiber/fiber/v2"
)

// Register.
// @Register
// @Summary  register.
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Param    register  body      models.UserRegistration  true  "Register"
// @Success  200    {object}  models.AuthUser
// @Router   /users/register [post]

func (s *userController) CreateUser(c *fiber.Ctx) error {
	fmt.Println("---Register Route---")
	userRegisterData := new(models.User)
	if err := c.BodyParser(&userRegisterData); err != nil {
		return fiber.ErrBadRequest
	}

	user, err := s.userService.CreateUser(*userRegisterData)
	stringObjectID := user.ID.Hex()
	token, err := auth.GereteToken(user.Contact.Email, stringObjectID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(models.AuthUser{User: *user, Token: token})

}
