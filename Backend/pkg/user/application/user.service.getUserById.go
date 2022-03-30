package user_service

import (
	models "zetting/pkg/user/core/models"
)

func (r *userService) GetUserById(userId interface{}) (*models.User, error) {
	return r.userRepo.GetUserById(userId)

}
