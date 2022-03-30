package user_handlers

import (
	"github.com/gofiber/fiber/v2"
)

// CheckUserExist.
// @CheckUserExist
// @Summary  check user exists.
// @Tags     Helper
// @Param    email  path  string  true  "Email"
// @Succes   200
// @Router   /users/{email} [get]

func (s *userHandler) CheckEmailExist(c *fiber.Ctx) error {
	email := c.Params("email")
	check := s.userService.CheckEmailExist(email)
	resp := struct {
		Exists bool `json:"exists"`
	}{}
	resp.Exists = check
	return c.Status(fiber.StatusOK).JSON(resp)

}
