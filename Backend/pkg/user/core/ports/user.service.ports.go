package user_ports

import (
	models "zetting/pkg/user/core/models"
)

type UserService interface {
	UpdateUser(user *models.User) (*models.User, error)
	CreateUser(user models.User) (*models.User, error)
	LoginUser(email string, password string) (*models.AuthUser, error)
	GetUserById(userId interface{}) (*models.User, error)
	CheckEmailExist(email string) bool
	GetUsersByProfession(string) (models.Users, error)
}
