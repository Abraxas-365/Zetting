package repository

import (
	"fmt"
	"mongoCrud/database"
	m "mongoCrud/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionFeatures = database.GetCollection("Features")

func CreateFeatures(a *m.Features) (primitive.ObjectID, error) {

	result, err := collectionFeatures.InsertOne(ctx, a)
	id := result.InsertedID.(primitive.ObjectID)
	fmt.Println(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id, nil
}
