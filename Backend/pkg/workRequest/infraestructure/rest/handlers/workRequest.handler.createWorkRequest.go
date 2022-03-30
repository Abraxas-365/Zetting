package workRequest_handlers

import (
	"zetting/auth"
	models "zetting/pkg/workRequest/core/models"

	"github.com/gofiber/fiber/v2"
)

func (s *workRequestHandler) CreateWorkRequest(c *fiber.Ctx) error {
	newWorkrequestData := new(models.WorkRequest)
	if err := c.BodyParser(&newWorkrequestData); err != nil {
		return fiber.ErrBadRequest
	}
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	newWorkrequestData.OwnerId = userTokenData.ID
	if err := s.userService.CreateWorkRequest(*newWorkrequestData); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)

}
