package repository

import (
	"fmt"
	m "mongoCrud/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetByProfession(profession string) (m.Users, error) {
	fmt.Println("---GetByProfession---")
	var users m.Users
	// filter := bson.M{"profession": bson.M{"$all": bson.A{profession}}} esto es por si se vuelve un array
	filter := bson.D{primitive.E{Key: "profession.name", Value: profession}}
	cur, err := collectionUser.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
