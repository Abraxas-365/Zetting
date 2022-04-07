package notification_ports

import (
	models "zetting/pkg/notification/core/models"
)

type NotificationRepository interface {
	CreateNotification(newNotification models.Notification) error
	GetNotifications(userId interface{}, page int) (models.Notifications, error)
	// UpdateNotification(notification models.Notification) error
}
