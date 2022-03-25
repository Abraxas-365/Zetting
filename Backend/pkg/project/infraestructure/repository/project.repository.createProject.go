package project_repository

import (
	"context"
	"fmt"
	"time"
	models "zetting/pkg/project/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) CreateProject(newProject *models.Project, userId string) (string, error) {
	fmt.Println("--CreateProjectRepo--", newProject.Name, userId)
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	userIdObjectId, err := primitive.ObjectIDFromHex(userId)
	check := bson.M{}
	filter := bson.M{"name": newProject.Name, "propietarios": bson.A{userIdObjectId}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return "", err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&check); err != nil {
			return "", err
		}
	}
	fmt.Println("tamano", len(check))
	if len(check) >= 1 {

		return "", ErrConflict
	}
	newProject.Created = time.Now()
	newProject.Updated = time.Now()
	newProject.Propietarios = []primitive.ObjectID{userIdObjectId}
	newProject.Workers = []primitive.ObjectID{}
	newProject.WorkRequests = []primitive.ObjectID{}
	result, err := collection.InsertOne(ctx, newProject)
	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, err
}
