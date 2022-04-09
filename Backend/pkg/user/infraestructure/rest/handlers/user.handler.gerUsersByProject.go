package user_handlers

import "github.com/gofiber/fiber/v2"

func (h *userHandler) GetUsersByProject(c *fiber.Ctx) error {
	projectId := c.Params(":project_id")

	users, err := h.userService.GetUsersByProject(projectId, "owner")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
