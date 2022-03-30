package user_service

import (
	"fmt"
	models "zetting/pkg/user/core/models"
)

func (r *userService) UpdateUser(user *models.User) (*models.User, error) {

	// user, _ := r.userRepo.GetUserByEmail(email)
	updateQuery := make(map[string]interface{})

	switch true {
	case len(user.FirstName) > 0:
		updateQuery["first_name"] = user.FirstName

	case len(user.LastName) > 0:
		updateQuery["last_name"] = user.LastName

	case len(user.Password) > 0:
		/*funcion de crear password*/
		fmt.Println("funcion para encriptar el password")
		updateQuery["password"] = user.Password

	case len(user.PerfilImage) > 0:
		updateQuery["perfil_image"] = user.PerfilImage

	case len(user.Contact.Country) > 0:
		updateQuery["contact.country"] = user.Contact.Country

	case len(user.Contact.Email) > 0:
		updateQuery["contact.email"] = user.Contact.Email
	case len(user.Contact.Phone) > 0:
		updateQuery["contact.phone"] = user.Contact.Phone

	}
	err := r.userRepo.UpdateUser(updateQuery, user.ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
