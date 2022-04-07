package notification_service

import models "zetting/pkg/notification/core/models"

func (s *notificationService) GetNotifications(userId interface{}, page int) (models.Notifications, error) {
	notifications, err := s.notificationRepo.GetNotifications(userId, page)
	if err != nil {
		return nil, err
	}
	return notifications, nil

}
