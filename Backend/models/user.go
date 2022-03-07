package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                      primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	FirstName               string               `json:"first_name"`
	LastName                string               `json:"last_name"`
	Gender                  string               `json:"gender"`
	Password                string               `json:"password"`
	Email                   string               `json:"email"`
	Phone                   string               `json:"phone"`
	Country                 string               `json:"country"`
	IdentifierDocumentation string               `json:"identifier_document"`
	Profesion               string               `json:"profesion"`
	Projects                []primitive.ObjectID `json:"projects"`
	MyProjects              []primitive.ObjectID `json:"my_projects"`
	Verified                bool                 `json:"verified"`
	Created                 time.Time            `json:"created_at"`
	Updated                 time.Time            `json:"updated_at,omitempty"`
}

type AuthUser struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
type Users []*User
