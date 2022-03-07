package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetCollection(collection string) *mongo.Collection {
	// uri := "mongodb+srv://abraxas:kancerver0@cluster0.lbtx1.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		panic(err.Error())
	}

	return client.Database("Zetting").Collection(collection)
}
