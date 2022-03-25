package user_service

import (
	"errors"
	ports "zetting/pkg/user/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
	ErrEmptyParams  = errors.New("Empty parameters")
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUsertService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		userRepo,
	}
}
