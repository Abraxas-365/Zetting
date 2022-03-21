package service

import (
	"fmt"
	m "mongoCrud/models"
	repository "mongoCrud/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProject(newProject *m.Proyecto, uid string) (primitive.ObjectID, error) {
	//crear el proyecto en el uId
	userId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	fmt.Println("crear el proyecto en el uId ", userId)
	// enviar a a bd
	projectId, err := repository.CreateProject(newProject, userId)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	if err := repository.AddProject(userId, projectId, "myprojects"); err != nil {
		return primitive.ObjectID{}, err
	}
	return projectId, nil

}

func GetMyProjects(uid string) (m.Proyectos, error) {

	userId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return nil, err
	}
	myprojects, err := repository.GetMyProjects(userId)
	if err != nil {
		return nil, err
	}
	return myprojects, nil

}

func GetProjectsWorkingOn(uid string) (m.Proyectos, error) {
	userId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return nil, err
	}
	proyects, err := repository.GetProjectsWorkingOn(userId)
	if err != nil {
		return nil, err
	}
	return proyects, nil

}
