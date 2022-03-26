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
	check := new(models.User)
	filter := bson.M{"contact.email": user.Contact.Email}
	if err := collection.FindOne(ctx, filter).Decode(&check); err != nil {
		user.Created = time.Now()
		user.Updated = time.Now()
		user.MyProjects = []primitive.ObjectID{}
		user.Projects = []primitive.ObjectID{}
		user.WorkRequests = []primitive.ObjectID{}
		user.Features.Skills = []string{}
		user.Verified = false
		result, err := collection.InsertOne(ctx, user)
		if err != nil {
			return nil, err
		}
		user.ID = result.InsertedID
		return &user, nil

	}
	return nil, ErrConflict
}
