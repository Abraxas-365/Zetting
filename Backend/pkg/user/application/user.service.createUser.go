package user_service

import (
	"fmt"
	models "zetting/pkg/user/core/models"
)

func (r *userService) CreateUser(user models.User) (*models.User, error) {
	fmt.Println("---CreateUserService ---")

	_, err := r.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	newUser, err := r.LoginUser(user.Contact.Email, user.Password)
	if err != nil {
		return nil, err
	}

	return newUser, nil

}
