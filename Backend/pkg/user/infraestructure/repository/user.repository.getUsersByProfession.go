package user_repository

import (
	"context"
	"fmt"
	models "zetting/pkg/user/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) GetUsersByProfession(profession string) (models.Users, error) {
	fmt.Println("--dd", profession)
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	var users models.Users
	// filter := bson.M{"profession": bson.M{"$all": bson.A{profession}}} esto es por si se vuelve un array
	filter := bson.D{primitive.E{Key: "profession.name", Value: profession}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil

}
