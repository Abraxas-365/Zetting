package service

import (
	m "mongoCrud/models"
	repository "mongoCrud/repositories"
)

func GetByProfession(profession string) (m.Users, error) {
	users, err := repository.GetByProfession(profession)
	if err != nil {
		return nil, err
	}
	return users, nil
}
