package user_service

import (
	models "zetting/pkg/user/core/models"
)

func (r *userService) GetUserById(id string) (*models.User, error) {
	return r.userRepo.GetUserById(id)

}
