package user_handlers

import (
	"fmt"
	models "zetting/pkg/user/core/models"

	"github.com/gofiber/fiber/v2"
)

func (s *userController) UpdateUser(c *fiber.Ctx) error {
	fmt.Println("---UpdateUser Route---")
	updatedUser := new(models.User)
	if err := c.BodyParser(&updatedUser); err != nil {
		return fiber.ErrBadRequest
	}
	return nil

}
