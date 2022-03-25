package user_repository

import (
	"context"
	models "zetting/pkg/user/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) GetUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user models.User
	filter := bson.M{"_id": userId}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil

}
