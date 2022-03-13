package repository

import (
	"fmt"
	m "mongoCrud/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetByFeatures(query bson.D) (m.Users, error) {
	fmt.Println("---GetByFeatures---")
	var s m.Users
	cur, err := collectionUser.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &s); err != nil {
		panic(err)
	}
	fmt.Println(s)
	return s, nil
}
