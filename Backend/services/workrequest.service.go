package service

import (
	"fmt"
	m "mongoCrud/models"
	repository "mongoCrud/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddWorkRequest(owner string, project string, worker string) error {
	fmt.Println("owner", owner)
	fmt.Println("project", project)
	fmt.Println("worker", worker)
	projectId, err := primitive.ObjectIDFromHex(project)
	if err != nil {
		return err
	}
	ownerId, err := primitive.ObjectIDFromHex(owner)
	if err != nil {
		return err
	}
	workerId, err := primitive.ObjectIDFromHex(worker)
	if err != nil {
		return err
	}

	if err := repository.CreateWorkRequest(ownerId, projectId, workerId); err != nil {
		return err
	}

	return nil

}

func AcceptWorkRequest(workRequest string, user string, accept bool) error {

	workRequestId, err := primitive.ObjectIDFromHex(workRequest)
	if err != nil {
		return err
	}
	userId, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return err
	}
	repository.AcceptWorkRequest(workRequestId, userId)
	return nil
}
func GetWorkRequest(workRequest string) (*m.WorkRequest, error) {

	return nil, nil
}
