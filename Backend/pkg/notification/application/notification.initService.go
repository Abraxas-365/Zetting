package notification_service

import (
	"errors"
	ports "zetting/pkg/notification/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type notificationService struct {
	notificationRepo ports.NotificationService
}

func NewNotificationService(notificationRepo ports.NotificationRepository) ports.NotificationService {
	return &notificationService{
		notificationRepo,
	}

}
