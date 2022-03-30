package user_handlers

import (
	"zetting/auth"

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
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	user, err := s.userService.GetUserById(userTokenData.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusNetworkAuthenticationRequired)
	}
	return c.Status(fiber.StatusOK).JSON(user)

}
