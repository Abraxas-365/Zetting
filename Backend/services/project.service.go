package service

import (
	"fmt"
	m "mongoCrud/models"
	repository "mongoCrud/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProject(p *m.Proyecto, id string) (primitive.ObjectID, error) {
	//crear el proyecto en el Id
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	fmt.Println("crear el proyecto en el Id ", userId)
	// enviar a a bd
	projectId, err := repository.CreateProject(p, userId)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	if err := repository.AddProject(userId, projectId, "myprojects"); err != nil {
		return primitive.ObjectID{}, err
	}
	return projectId, nil

}

func GetMyProjects(id string) (m.Proyectos, error) {

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ps, err := repository.GetMyProjects(userId)
	if err != nil {
		return nil, err
	}
	return ps, nil

}

func GetProjectsWorkingOn(id string) (m.Proyectos, error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ps, err := repository.GetProjectsWorkingOn(userId)
	if err != nil {
		return nil, err
	}
	return ps, nil

}
