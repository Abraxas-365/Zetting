package user_service

import (
	models "zetting/pkg/user/core/models"
)

func (r *userService) LoginUser(email string, password string) (*models.User, error) {

	switch true {
	case email == "":
		return nil, ErrEmptyParams
	case password == "":
		return nil, ErrEmptyParams
	}

	user, _ := r.userRepo.GetUserByEmail(email)

	switch true {
	case user == nil:
		return nil, ErrUnauthorized
	case user.Password != password:
		return nil, ErrUserNotFound
	}

	return user, nil
}
