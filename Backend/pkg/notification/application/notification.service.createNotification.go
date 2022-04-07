package notification_service

import models "zetting/pkg/notification/core/models"

func (s *notificationService) CreateNotification(newNotification models.Notification) error {

	switch {
	case newNotification.Type == "new-work-request":
		newNotification.Message = "Work request from"

	}

	if err := s.notificationRepo.CreateNotification(newNotification); err != nil {
		return err
	}
	return nil

}
