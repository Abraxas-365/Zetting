package user_service

import ()

func (r *userService) CheckEmailExist(email string) bool {

	user, _ := r.userRepo.GetUserByEmail(email)
	if user == nil {
		return false
	}
	return true

}
