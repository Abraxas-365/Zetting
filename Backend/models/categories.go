package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Actor struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId     string             `json:"user_id"`
	Height     int                `json:"height"`
	Body       string             `json:"body"`
	Skin       string             `json:"skin"`
	HairType   []string           `json:"hair_type"`
	Skills     []string           `json:"skills"`
	FacialHair string             `json:"facial_hair"`
	HairColor  string             `json:"hair_color"`
	EyeColor   string             `json:"eye_color"`
}

type Actors []*Actor
