package notification_handlers

import (
	"strconv"
	"zetting/auth"

	"github.com/gofiber/fiber/v2"
)

func (h *notificationHandler) GetNotifications(c *fiber.Ctx) error {

	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	page, _ := strconv.Atoi(c.Params("page"))
	myNotifications, err := h.notificationService.GetNotifications(userTokenData.ID, page)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(myNotifications)
}
