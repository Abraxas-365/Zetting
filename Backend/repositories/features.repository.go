package repository

import (
	"fmt"
	m "mongoCrud/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetByFeatures(query bson.D) (m.Users, error) {
	fmt.Println("---GetByFeatures---")
	var u m.Users
	cur, err := collectionUser.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &u); err != nil {
		panic(err)
	}
	fmt.Println(u)
	return u, nil
}
