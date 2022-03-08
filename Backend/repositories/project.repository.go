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

func CreateProject(p *m.Proyecto, email string) error {

	fmt.Println(p.Name)
	check := bson.M{}
	filter := bson.M{"name": p.Name, "propietarios": bson.A{email}}
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

		p.Created = time.Now()
		p.Updated = time.Now()
		p.Propietarios = []string{email}
		p.Workers = []string{}
		fmt.Println(p.Created)
		result, err := collectionProject.InsertOne(ctx, p)
		id := result.InsertedID.(primitive.ObjectID)

		if err != nil {
			return err
		}
		AddProject(email, id)
	} else {

		return fiber.ErrConflict
	}
	fmt.Println("nuevo usuario", p)
	return nil
}

func GetMyProjects(email string) (m.Proyectos, error) {
	var ps m.Proyectos

	matchStage := bson.D{{"$match", bson.D{{"email", email}}}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "Projects"}, {"localField", "myprojects"}, {"foreignField", "_id"}, {"as", "myprojects"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$myprojects"}, {"preserveNullAndEmptyArrays", false}}}}
	replaceRoot := bson.D{{"$replaceRoot", bson.D{{"newRoot", "$myprojects"}}}}

	cur, err := collectionUser.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage, replaceRoot})
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &ps); err != nil {
		panic(err)
	}

	return ps, nil
}
func GetProjectsWorkingOn(email string) (m.Proyectos, error) {
	var ps m.Proyectos

	filter := bson.M{"workers": bson.A{email}}
	cur, err := collectionProject.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var p m.Proyecto
		if err = cur.Decode(&p); err != nil {
			return nil, err
		}
		ps = append(ps, &p)
	}
	return ps, nil
}

//en version v2 cambiar a buscar por usurario
func AddWorker(email string, projectId primitive.ObjectID) error {
	filter := bson.M{"_id": projectId}
	update := bson.M{
		"$push": bson.M{
			"workers": email,
		},
	}

	if _, err := collectionUser.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil

}
