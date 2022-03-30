package user_handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *userController) GetUsersByProfession(c *fiber.Ctx) error {

	profession := c.Params("profession")
	page, _ := strconv.Atoi(c.Params("page"))
	users, err := s.userService.GetUsersByProfession(profession, page)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.Status(fiber.StatusOK).JSON(users)

}
