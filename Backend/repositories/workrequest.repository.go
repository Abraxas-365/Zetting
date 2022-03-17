package repository

import (
	"fmt"
	m "mongoCrud/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateWorkRequest(ownerId primitive.ObjectID, projectId primitive.ObjectID, workerId primitive.ObjectID) error {
	var id primitive.ObjectID
	wrq := new(m.WorkRequest)
	fmt.Println("---CreateWorkRequest---")
	check := bson.M{}
	filter := bson.M{"projectId": projectId, "workerId": workerId}
	cur, err := collectionProject.Find(ctx, filter)
	if err != nil {
		return err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&check); err != nil {
			return err
		}
	}
	if len(check) < 1 {

		wrq.OwnerId = ownerId
		wrq.ProjectId = projectId
		wrq.WorkerId = workerId
		wrq.Created = time.Now()
		wrq.Updated = time.Now()
		fmt.Println(wrq.Created)
		result, err := collectionProject.InsertOne(ctx, wrq)
		id = result.InsertedID.(primitive.ObjectID)

		if err != nil {
			return err
		}
	} else {

		return fiber.ErrConflict
	}

	//agregar el work request al rpyecto
	filter = bson.M{"_id": projectId}
	update := bson.M{
		"$push": bson.M{
			"work_requests": id,
		},
	}

	if _, err := collectionProject.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	//agregar el work request al usuario
	filter = bson.M{"_id": workerId}
	update = bson.M{
		"$push": bson.M{
			"work_requests": id,
		},
	}

	if _, err := collectionUser.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil

}
