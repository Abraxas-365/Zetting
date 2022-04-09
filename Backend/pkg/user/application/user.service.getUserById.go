package user_service

import (
	models "zetting/pkg/user/core/models"
)

func (r *userService) GetUserById(userId interface{}) (*models.User, error) {
	user, err := r.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}
