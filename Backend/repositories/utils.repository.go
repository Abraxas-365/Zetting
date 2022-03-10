package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//funcion para agregaar un proyecccto a cualquera de los 2 campos o proximos campos ejemplo :"myproject" o "projects"
func AddProject(userId primitive.ObjectID, projectId primitive.ObjectID, campo string) error {

	filter := bson.M{"_id": userId}
	update := bson.M{
		"$push": bson.M{
			campo: projectId,
		},
	}

	if _, err := collectionUser.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil

}
