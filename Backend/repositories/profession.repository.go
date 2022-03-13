package repository

import (
	"fmt"
	m "mongoCrud/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetByProfession(profession string) (m.Users, error) {
	fmt.Println("---GetByProfession---")
	var users m.Users
	filter := bson.M{"profession": bson.M{"$all": bson.A{profession}}}
	cur, err := collectionUser.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
