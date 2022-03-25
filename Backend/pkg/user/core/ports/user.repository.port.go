package user_ports

import (
	models "zetting/pkg/user/core/models"
)

type UserRepository interface {
	CreateUser(user models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
	GetUsersByProfession(string) (models.Users, error)
}
