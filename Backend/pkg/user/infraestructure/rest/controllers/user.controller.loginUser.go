package user_controllers

import (
	"fmt"
	"zetting/auth"
	models "zetting/pkg/user/core/models"

	"github.com/gofiber/fiber/v2"
)

// Login.
// @Login
// @Summary  authentication.
// @Tags  Auth
// @Accept   json
// @Produce  json
// @Param    login  body  models.UserLogin  true  "Login"
// @Success   200   {object}  models.AuthUser
// @Router   /users/login [post]

func (s *userController) LoginUser(c *fiber.Ctx) error {
	fmt.Println("---Register Route---")
	userLoginData := new(models.UserLogin)
	if err := c.BodyParser(&userLoginData); err != nil {
		return fiber.ErrBadRequest
	}

	userRegisterData := new(models.User)
	if err := c.BodyParser(&userRegisterData); err != nil {
		return fiber.ErrBadRequest
	}

	user, err := s.userService.LoginUser(userLoginData.Email, userLoginData.Password)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString(err.Error())
	}
	token, err := auth.GereteToken(user.Contact.Email, user.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(models.AuthUser{User: *user, Token: token})

}
