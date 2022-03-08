package service

import (
	"fmt"
	m "mongoCrud/models"
	repository "mongoCrud/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProject(p *m.Proyecto, email string) error {
	//crear el proyecto en el correo correcto
	fmt.Println("crear el proyecto en el correo correct ", email)
	// enviar a a bd
	if err := repository.CreateProject(p, email); err != nil {
		return err
	}
	return nil

}

func GetMyProjects(email string) (m.Proyectos, error) {
	ps, err := repository.GetMyProjects(email)
	if err != nil {
		return nil, err
	}
	return ps, nil

}

func GetProjectsWorkingOn(email string) (m.Proyectos, error) {
	ps, err := repository.GetProjectsWorkingOn(email)
	if err != nil {
		return nil, err
	}
	return ps, nil

}
func AddWorker(email string, projectId string) error {
	// comprobar que el el proyect id le pertenece al usuario
	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return err
	}
	if err = repository.AddWorker(email, objID); err != nil {
		return err
	}
	return nil

}
