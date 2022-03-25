package user_service

import models "zetting/pkg/user/core/models"

func (r *userService) GetUsersByProfession(profession string) (models.Users, error) {
	users, err := r.userRepo.GetUsersByProfession(profession)
	if err != nil {
		return nil, err
	}
	return users, nil
}
