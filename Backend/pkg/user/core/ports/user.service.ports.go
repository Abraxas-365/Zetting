package user_ports

import (
	models "zetting/pkg/user/core/models"
)

type UserService interface {
	UpdateUser(userDataToUpdate *models.User, userId interface{}) error
	CreateUser(user models.User) (*models.User, error)
	LoginUser(email string, password string) (*models.AuthUser, error)
	GetUserById(userId interface{}) (*models.User, error)
	CheckEmailExist(email string) bool
	GetUsersByProfession(profession string, page int) (models.Users, error)
	GetUsersByProject(projectId interface{}, document string) (models.Users, error)
}
