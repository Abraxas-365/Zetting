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

func CreateUser(user m.User) (primitive.ObjectID, error) {
	var id primitive.ObjectID
	fmt.Println(user.Contact.Email)
	check := bson.M{}
	filter := bson.M{"email": user.Contact.Email}
	cur, err := collectionUser.Find(ctx, filter)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&check); err != nil {
			return primitive.NewObjectID(), err
		}
	}
	if len(check) < 1 {

		user.Created = time.Now()
		user.Updated = time.Now()
		user.Projects = []primitive.ObjectID{}
		user.MyProjects = []primitive.ObjectID{}
		user.WorkRequests = []primitive.ObjectID{}
		user.Features.Skills = []string{}
		user.Verified = false
		fmt.Println(user.Created)
		result, err := collectionUser.InsertOne(ctx, user)
		id = result.InsertedID.(primitive.ObjectID)
		if err != nil {
			return primitive.ObjectID{}, err
		}
	} else {

		return primitive.ObjectID{}, fiber.ErrConflict
	}

	fmt.Println("nuevo usuario", user)
	fmt.Println("id", id)
	return id, nil

}
func GetUserByEmail(email string) (*m.User, error) {
	fmt.Println("---GetUserByEmail---", email)
	var user m.User
	filter := bson.M{"contact.email": email}
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(uid primitive.ObjectID) (*m.User, error) {

	fmt.Println("---GetUserById---")
	var user m.User
	filter := bson.M{"_id": uid}
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckUserExist(email string) (bool, error) {
	var user m.User
	filter := bson.M{"contact.email": email}
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
			"email":      user.Contact.Email,
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
