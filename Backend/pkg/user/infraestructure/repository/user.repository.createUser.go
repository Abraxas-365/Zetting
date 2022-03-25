package user_repository

import (
	"context"
	"time"
	models "zetting/pkg/user/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) CreateUser(user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	var id primitive.ObjectID
	check := bson.M{}
	filter := bson.M{"email": user.Contact.Email}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&check); err != nil {
			return nil, err
		}
	}
	if len(check) < 1 {
		user.Created = time.Now()
		user.Updated = time.Now()
		user.Projects = []primitive.ObjectID{}
		user.MyProjects = []primitive.ObjectID{}
		user.WorkRequests = []primitive.ObjectID{}
		user.Features.Skills = []string{}
		user.Verified = false
		result, err := collection.InsertOne(ctx, user)
		id = result.InsertedID.(primitive.ObjectID)
		if err != nil {
			return nil, err
		}
		user.ID = id
	} else {

		return nil, ErrConflict
	}

	return &user, nil

}
