package user_controllers

import "github.com/gofiber/fiber/v2"

func (s *userController) GetUsersByProfession(c *fiber.Ctx) error {

	profession := c.Params("profession")
	users, err := s.userService.GetUsersByProfession(profession)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.Status(fiber.StatusOK).JSON(users)

}
