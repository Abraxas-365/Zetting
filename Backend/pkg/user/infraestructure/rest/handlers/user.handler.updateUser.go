package user_handlers

import (
	"fmt"
	"zetting/auth"
	models "zetting/pkg/user/core/models"

	"github.com/gofiber/fiber/v2"
)

func (s *userHandler) UpdateUser(c *fiber.Ctx) error {
	fmt.Println("---UpdateUser Route---")
	updatedUser := new(models.User)
	if err := c.BodyParser(&updatedUser); err != nil {
		return fiber.ErrBadRequest
	}
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := s.userService.UpdateUser(updatedUser, userTokenData.ID); err != nil {
		return fiber.ErrBadRequest
	}
	return c.SendStatus(fiber.StatusOK)

}
