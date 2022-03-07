package service

import (
	"fmt"
	"mongoCrud/auth"
	m "mongoCrud/models"
	repository "mongoCrud/repositories"
)

func CreateUser(u m.User) (*m.AuthUser, error) {

	if err := repository.CreateUser(u); err != nil {
		return nil, err
	}
	authUser, err := AuthUser(u.Email, u.Password)
	if err != nil {
		return nil, err
	}
	return authUser, nil
}

func AuthUser(email string, password string) (*m.AuthUser, error) {

	authUser := new(m.AuthUser)

	u, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	t, err := auth.LoginUser(email, password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	authUser.Token = t
	authUser.User = *u
	return authUser, nil

}

func GetUsers() (m.Users, error) {
	users, err := repository.GetUsers()
	if err != nil {

		return nil, err
	}
	return users, nil
}

func CheckUserExist(email string) (*m.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	fmt.Println(user.FirstName)
	return user, nil
}

func UpdateUser(user m.User, id string) error {
	if err := repository.UpdateUser(user, id); err != nil {
		return err
	}

	return nil

}

func DeleteUser(id string) error {
	if err := repository.DeleteUser(id); err != nil {
		return err
	}
	return nil

}
