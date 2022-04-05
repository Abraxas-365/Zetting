package user_handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetUser.
// @GetUser
// @Summary   get user.
// @Tags    User
// @Success  200
// @Security  ApiKeyAuth
// @Router    /users/ [get]

func (s *userHandler) GetUserById(c *fiber.Ctx) error {
	userId := c.Params("id")
	fmt.Println(userId)
	user, err := s.userService.GetUserById(userId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	fmt.Println(user.FirstName)
	return c.Status(fiber.StatusOK).JSON(user)

}
