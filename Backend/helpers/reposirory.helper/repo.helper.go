package repositoryhelper

import (
	"context"
	"fmt"
	"mongoCrud/database"
	m "mongoCrud/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionLugares = database.GetCollection("ubicaciones")
var collectionJuegos = database.GetCollection("document_juego")
var ctx = context.Background()

func GetRandomPlaces(tipo int, size int) (m.Ubicaciones, error) {
	var lugares m.Ubicaciones
	filter := bson.D{{"$match", bson.D{{"tipo", tipo}}}}
	random := bson.D{{"$sample", bson.D{{"size", size}}}}
	cur, err := collectionLugares.Aggregate(ctx, mongo.Pipeline{filter, random})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var lugar m.Ubicacion
		if err = cur.Decode(&lugar); err != nil {
			return nil, err
		}
		fmt.Println("lugaar", lugar)
		lugares = append(lugares, &lugar)
	}
	return lugares, nil

}

func GetRandomGame(tipo int, size int) (*m.Juego, error) {
	fmt.Println("-----get random game---")
	juego := new(m.Juego)
	filter := bson.D{{"$match", bson.D{{"tipo", tipo}}}}
	// random := bson.D{{"$sample", bson.D{{"size", size}}}}
	cur, err := collectionJuegos.Aggregate(ctx, mongo.Pipeline{filter})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		if err = cur.Decode(&juego); err != nil {
			fmt.Println(err)
			return nil, err
		}
		a := juego.Answer
		fmt.Println("respuesta", juego)
		fmt.Println(a)
	}
	return juego, nil

}
