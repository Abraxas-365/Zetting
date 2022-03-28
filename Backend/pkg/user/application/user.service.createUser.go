package user_service

import (
	"fmt"
	models "zetting/pkg/user/core/models"
)

func (r *userService) CreateUser(user models.User) (*models.User, error) {
	fmt.Println("---CreateUserService ---")

	/*Crear el usuario*/
	newUser, err := r.userRepo.CreateUser(user)
	if err != nil {

		return nil, err
	}

	/*Crear carpeta con el nombre del usuario*/

	return newUser, nil

}
