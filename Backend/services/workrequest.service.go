package service

import (
	repository "mongoCrud/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddWorkRequest(owner string, project string, worker string) error {
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
