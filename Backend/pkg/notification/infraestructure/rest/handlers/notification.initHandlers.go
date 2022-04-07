package notification_handlers

import (
	ports "zetting/pkg/notification/core/ports"

	"github.com/gofiber/fiber/v2"
)

type NotificationsHandler interface {
	CreateNotification(c *fiber.Ctx) error
	GetNotifications(c *fiber.Ctx) error
}
type notificationHandler struct {
	notificationService ports.NotificationService
}

func NewNotificationHandler(notificationService ports.NotificationService) NotificationsHandler {
	return &notificationHandler{
		notificationService,
	}
}
