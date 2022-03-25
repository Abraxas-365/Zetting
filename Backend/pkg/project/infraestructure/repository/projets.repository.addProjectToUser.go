package project_repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) AddProject(userId string, projectId string, campo string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("Users")

	userObjecId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	projectObjecId, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return err
	}
	fmt.Println("obj", userObjecId)
	fmt.Println("obj", projectObjecId)

	filter := bson.M{"_id": userObjecId}
	update := bson.M{
		"$push": bson.M{
			campo: projectObjecId,
		},
	}

	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil

}
