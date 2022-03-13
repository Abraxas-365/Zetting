package service

import (
	"fmt"
	m "mongoCrud/models"
	repository "mongoCrud/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FilterByFeatures(profession string, filter m.Filter) (m.Users, error) {
	fmt.Println("---GetByFeatures service ---")
	query := bson.D{primitive.E{Key: "profession", Value: bson.D{primitive.E{Key: "$all", Value: bson.A{profession}}}}}
	//filtro de edad
	if filter.AgeTop != 0 {
		query = append(query, primitive.E{Key: "features.age", Value: bson.D{primitive.E{Key: "$gte", Value: filter.AgeLow}}})
		query = append(query, primitive.E{Key: "features.age", Value: bson.D{primitive.E{Key: "$lte", Value: filter.AgeTop}}})
	}
	//filtro de estatura
	if filter.HeightTop != 0 {
		query = append(query, primitive.E{Key: "features.height", Value: bson.D{primitive.E{Key: "$gte", Value: filter.HeightLow}}})
		query = append(query, primitive.E{Key: "features.height", Value: bson.D{primitive.E{Key: "$lte", Value: filter.HeightTop}}})
	}

	//filtro de cuerpo
	if filter.Body != "" {
		query = append(query, primitive.E{Key: "features.body", Value: filter.Body})
	}
	//filtro de skin
	if filter.Skin != "" {
		query = append(query, primitive.E{Key: "features.skin", Value: filter.Skin})
	}
	//filtro de cabello tipo
	if filter.HairType != "" {
		query = append(query, primitive.E{Key: "features.hair_type", Value: filter.HairType})
	}
	//filtro de cabello tamano
	if filter.HairZise != "" {
		query = append(query, primitive.E{Key: "features.hair_zise", Value: filter.HairZise})
	}
	// filtro de barba
	if filter.FacialHair != "" {
		query = append(query, primitive.E{Key: "features.facial_hair", Value: filter.FacialHair})
	}
	//filtro de color cabello
	if filter.FacialHair != "" {
		query = append(query, primitive.E{Key: "features.hair_color", Value: filter.HairColor})
	}
	//filtro de color ojos
	if filter.EyeColor != "" {
		query = append(query, primitive.E{Key: "features.eye_color", Value: filter.EyeColor})
	}
	//filtro de abilidades
	if len(filter.Skills) != 0 {
		fmt.Println(filter.Skills)
		query = append(query, primitive.E{Key: "features.skills", Value: bson.D{primitive.E{Key: "$all", Value: filter.Skills}}})
	}

	u, _ := repository.GetByFeatures(query)
	return u, nil
}
