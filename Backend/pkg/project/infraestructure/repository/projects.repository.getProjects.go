package project_repository

import (
	"context"
	models "zetting/pkg/project/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *mongoRepository) GetProjects(userId string, document string) (models.Projects, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("Users")

	var projects models.Projects
	userObjecId, err := primitive.ObjectIDFromHex(userId)

	//matchea el email
	matchStage := bson.D{primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: userObjecId}}}}
	//buca y trae todo los id en projects
	lookupStage := bson.D{primitive.E{Key: "$lookup", Value: bson.D{primitive.E{Key: "from", Value: "Projects"}, {Key: "localField", Value: "myprojects"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: document}}}}
	unwindStage := bson.D{primitive.E{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$" + document}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	//reemplaza todo el objeeto por solo el campo quq queremos
	replaceRoot := bson.D{primitive.E{Key: "$replaceRoot", Value: bson.D{primitive.E{Key: "newRoot", Value: "$" + document}}}}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage, replaceRoot})
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
