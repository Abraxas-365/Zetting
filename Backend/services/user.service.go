package service

import (
	"fmt"
	"mongoCrud/auth"
	"mongoCrud/models"
	repository "mongoCrud/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (*models.AuthUser, error) {

	//crear usurio
	if _, err := repository.CreateUser(user); err != nil {
		return nil, err
	}
	authUser, err := AuthUser(user.Contact.Email, user.Password)
	if err != nil {
		return nil, err
	}
	return authUser, nil
}

func GetUser(id string) (*models.User, error) {

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := repository.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func AuthUser(email string, password string) (*models.AuthUser, error) {
	fmt.Println("----AuthUser ----")

	authUser := new(models.AuthUser)

	user, token, err := auth.LoginUser(email, password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	authUser.Token = token
	authUser.User = *user
	return authUser, nil
}

func CheckEmailExist(email string) (*models.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	fmt.Println(user.FirstName)
	return user, nil
}

func UpdateUser(user models.User, uid string) error {
	if err := repository.UpdateUser(user, uid); err != nil {
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
