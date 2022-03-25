package user_controllers

import (
	ports "zetting/pkg/user/core/ports"
)

type userController struct {
	userService ports.UserService
}

func NewUserController(userService ports.UserService) ports.UserController {
	return &userController{
		userService,
	}
}
