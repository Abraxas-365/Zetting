package repository

import (
	"fmt"
	"mongoCrud/database"
	m "mongoCrud/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionWorkRequest = database.GetCollection("WorkRequest")

func CreateWorkRequest(ownerId primitive.ObjectID, projectId primitive.ObjectID, workerId primitive.ObjectID) error {
	var id primitive.ObjectID
	wrq := new(m.WorkRequest)
	fmt.Println("---CreateWorkRequest---")
	check := bson.M{}
	filter := bson.M{"project_id": projectId, "worker_id": workerId}
	cur, err := collectionWorkRequest.Find(ctx, filter)
	if err != nil {
		return fiber.ErrConflict
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&check); err != nil {
			return err
		}
	}
	fmt.Println("tamano", len(check))
	if len(check) < 1 {

		wrq.OwnerId = ownerId
		wrq.ProjectId = projectId
		wrq.WorkerId = workerId
		wrq.Created = time.Now()
		wrq.Updated = time.Now()
		fmt.Println(wrq.Created)
		result, err := collectionWorkRequest.InsertOne(ctx, wrq)
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

func AcceptWorkRequest(workRequestId primitive.ObjectID, userId primitive.ObjectID) error {
	workRequest := new(m.WorkRequest)
	//encontrar el work request que coincidan con ambos params cambiar estado
	filter := bson.D{primitive.E{Key: "_id", Value: workRequestId}, primitive.E{Key: "worker_id", Value: userId}}
	if err := collectionWorkRequest.FindOne(ctx, filter).Decode(&workRequest); err != nil {
	}

	fmt.Println(workRequest)
	//buscar el projecto y anadir al paramtro usuario en worker
	filter = bson.D{primitive.E{Key: "_id", Value: workRequest.ProjectId}}
	update := bson.M{
		"$push": bson.M{
			"workers": workRequest.WorkerId,
		},
	}
	if _, err := collectionProject.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	//anadir el proyect id proyectos del usuario

	filter = bson.D{primitive.E{Key: "_id", Value: workRequest.WorkerId}}
	update = bson.M{
		"$push": bson.M{
			"projects": workRequest.ProjectId,
		},
	}
	if _, err := collectionUser.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	//eliminar el workRequest del usuario

	filter = bson.D{primitive.E{Key: "_id", Value: workRequest.WorkerId}}
	update = bson.M{
		"$pull": bson.M{
			"work_requests": workRequest.ID,
		},
	}
	if _, err := collectionUser.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

// func GetWorkRequest(workRequestId primitive.ObjectID) error {

// 	matchStage := bson.D{primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: userId}}}}
// 	lookupStage := bson.D{primitive.E{Key: "$lookup", Value: bson.D{primitive.E{Key: "from", Value: "Projects"}, {Key: "localField", Value: "projects"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "projects"}}}}
// 	unwindStage := bson.D{primitive.E{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$projects"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
// 	replaceRoot := bson.D{primitive.E{Key: "$replaceRoot", Value: bson.D{primitive.E{Key: "newRoot", Value: "$projects"}}}}

// 	cur, err := collectionUser.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage, replaceRoot})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = cur.All(ctx, &ps); err != nil {
// 		return nil, err
// 	}

// 	return ps, nil

// }
