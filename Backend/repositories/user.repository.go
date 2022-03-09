package repository

import (
	"context"
	"fmt"
	"mongoCrud/database"
	m "mongoCrud/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionUser = database.GetCollection("Users")
var ctx = context.Background()

func CreateUser(user m.User) error {

	fmt.Println(user.Email)
	check := bson.M{}
	filter := bson.M{"email": user.Email}
	cur, err := collectionUser.Find(ctx, filter)
	if err != nil {
		return err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&check); err != nil {
			return err
		}
	}
	if len(check) < 1 {

		user.Created = time.Now()
		user.Updated = time.Now()
		user.Projects = []primitive.ObjectID{}
		user.MyProjects = []primitive.ObjectID{}
		user.Verified = false
		fmt.Println(user.Created)
		if _, err := collectionUser.InsertOne(ctx, user); err != nil {
			return err
		}
	} else {

		return fiber.ErrConflict
	}
	fmt.Println("nuevo usuario", user)
	return nil
}
func GetUsers() (m.Users, error) {
	var users m.Users

	filter := bson.D{}
	cur, err := collectionUser.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user m.User
		if err = cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
func GetUserByEmail(email string) (*m.User, error) {
	var user m.User
	filter := bson.M{"email": email}
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
func CheckUserExist(email string) (bool, error) {
	var user m.User
	filter := bson.M{"email": email}
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return false, err
	}
	return true, nil
}
func UpdateUser(user m.User, id string) error {

	var err error
	filter := bson.M{"nikname": id}
	update := bson.M{
		"$set": bson.M{
			"name":       user.FirstName,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	}

	if _, err = collectionUser.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}
func DeleteUser(id string) error {

	filter := bson.M{"nikname": id}

	if _, err := collectionUser.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil

}

//funcion para agregaar un proyecccto a cualquera de los 2 campos o proximos campos ejemplo :"myproject" o "projects"
func AddProject(email string, projectId primitive.ObjectID, campo string) error {

	filter := bson.M{"email": email}
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
