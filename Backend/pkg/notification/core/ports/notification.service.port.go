package notification_ports

import (
	models "zetting/pkg/notification/core/models"
)

type NotificationService interface {
	CreateNotification(newNotification models.Notification) error
	GetNotifications(userId interface{}, page int) (models.Notifications, error)
}
