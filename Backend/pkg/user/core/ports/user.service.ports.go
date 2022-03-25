package user_ports

import (
	models "zetting/pkg/user/core/models"
)

type UserService interface {
	CreateUser(user models.User) (*models.User, error)
	LoginUser(email string, password string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
	CheckEmailExist(email string) bool
	GetUsersByProfession(string) (models.Users, error)
}
