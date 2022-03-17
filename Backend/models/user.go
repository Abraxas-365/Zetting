package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                      primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	FirstName               string               `json:"first_name"`
	LastName                string               `json:"last_name"`
	Password                string               `json:"password"`
	Email                   string               `json:"email"`
	Phone                   string               `json:"phone"`
	Country                 string               `json:"country"`
	IdentifierDocumentation string               `json:"identifier_document"`
	Profession              string               `json:"professions"`
	Features                Features             `json:"features"`
	Projects                []primitive.ObjectID `json:"projects"`
	MyProjects              []primitive.ObjectID `json:"my_projects"`
	WorkRequests            []primitive.ObjectID `bson:"work_requests" json:"work_requests"`
	Verified                bool                 `json:"verified"`
	Created                 time.Time            `json:"created_at"`
	Updated                 time.Time            `json:"updated_at,omitempty"`
}

//EJEMPLO
// 	"user": {
// 		"id": "622b6c3516e5db00e3acaf2d",
// 		"first_name": "lis",
// 		"last_name": "miranda",
// 		"password": "kancerver0",
// 		"email": "teta8@gmail.com",
// 		"phone": "932222",
// 		"country": "PE",
// 		"identifier_document": "71135626",
// 		"professions": [
// 			"abogado",
// 			"dd"
// 		],
// 		"features": {
// 			"gender": "",
// 			"age": 23,
// 			"height": 0,
// 			"body": "",
// 			"skin": "",
// 			"hair_type": null,
// 			"skills": null,
// 			"facial_hair": "",
// 			"hair_color": "",
// 			"eye_color": ""
// 		},
// 		"projects": [],
// 		"my_projects": [],
// 		"verified": false,
// 		"created_at": "2022-03-11T15:35:17.4Z",
// 		"updated_at": "2022-03-11T15:35:17.4Z"
// 	}

type AuthUser struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
type Users []*User
