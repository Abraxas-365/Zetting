package repository

import (
	"fmt"
	"mongoCrud/database"
	m "mongoCrud/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionProject = database.GetCollection("Projects")

func CreateProject(p *m.Proyecto, userId primitive.ObjectID) (primitive.ObjectID, error) {
	var id primitive.ObjectID
	fmt.Println(p.Name)
	check := bson.M{}
	filter := bson.M{"name": p.Name, "propietarios": bson.A{userId}}
	cur, err := collectionProject.Find(ctx, filter)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&check); err != nil {
			return primitive.ObjectID{}, err
		}
	}
	if len(check) < 1 {

		p.Created = time.Now()
		p.Updated = time.Now()
		p.Propietarios = []primitive.ObjectID{userId}
		p.Workers = []string{}
		fmt.Println(p.Created)
		result, err := collectionProject.InsertOne(ctx, p)
		id = result.InsertedID.(primitive.ObjectID)

		if err != nil {
			return primitive.ObjectID{}, err
		}
	} else {

		return primitive.ObjectID{}, fiber.ErrConflict
	}
	fmt.Println("el id del projecto es:", id)
	return id, nil
}

func GetMyProjects(userId primitive.ObjectID) (m.Proyectos, error) {
	var ps m.Proyectos

	//matchea el email
	matchStage := bson.D{primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: userId}}}}
	//buca y trae todo los id en proyectos
	lookupStage := bson.D{primitive.E{Key: "$lookup", Value: bson.D{primitive.E{Key: "from", Value: "Projects"}, {Key: "localField", Value: "myprojects"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "myprojects"}}}}
	unwindStage := bson.D{primitive.E{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$myprojects"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	//reemplaza todo el objeeto por solo el campo quq queremos
	replaceRoot := bson.D{primitive.E{Key: "$replaceRoot", Value: bson.D{primitive.E{Key: "newRoot", Value: "$myprojects"}}}}

	cur, err := collectionUser.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage, replaceRoot})
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &ps); err != nil {
		panic(err)
	}

	return ps, nil
}
func GetProjectsWorkingOn(userId primitive.ObjectID) (m.Proyectos, error) {
	var ps m.Proyectos

	matchStage := bson.D{primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: userId}}}}
	lookupStage := bson.D{primitive.E{Key: "$lookup", Value: bson.D{primitive.E{Key: "from", Value: "Projects"}, {Key: "localField", Value: "projects"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "projects"}}}}
	unwindStage := bson.D{primitive.E{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$projects"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	replaceRoot := bson.D{primitive.E{Key: "$replaceRoot", Value: bson.D{primitive.E{Key: "newRoot", Value: "$projects"}}}}

	cur, err := collectionUser.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage, replaceRoot})
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &ps); err != nil {
		panic(err)
	}

	return ps, nil
}

func AddWorker(userId primitive.ObjectID, projectId primitive.ObjectID) error {
	filter := bson.M{"_id": projectId}
	update := bson.M{
		"$push": bson.M{
			"workers": userId,
		},
	}

	if _, err := collectionUser.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil

}
