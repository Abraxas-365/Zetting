package user_ports

import (
	models "zetting/pkg/user/core/models"
)

type UserRepository interface {
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(query interface{}, userId interface{}) error
	CheckEmailExist(email string) (*models.User, error)
	GetUserById(userId interface{}) (*models.User, error)
	GetUsersByProfession(string) (models.Users, error)
}
