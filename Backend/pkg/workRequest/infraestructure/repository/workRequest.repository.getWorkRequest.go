package workRequest_repository

import (
	"context"
	models "zetting/pkg/workRequest/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mongoRepository) GetWorkRequests(id string, page int, document string) (models.WorkRequests, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	ObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var workRequests models.WorkRequests
	options := options.Find()
	options.SetLimit(20)
	options.SetSkip((int64(page) - 1) * 20)
	filter := bson.D{primitive.E{Key: document, Value: ObjectId}}
	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &workRequests); err != nil {
		return nil, err
	}
	return workRequests, nil
}
