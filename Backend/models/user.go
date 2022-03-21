package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	FirstName    string               `json:"first_name"`
	LastName     string               `json:"last_name"`
	Password     string               `json:"password"`
	PerfilImage  string               `json:"perfil_image"`
	Contact      UserContact          `json:"contact"`
	Profession   Profession           `json:"profession"`
	Features     Features             `json:"features"`
	Projects     []primitive.ObjectID `json:"projects"`
	MyProjects   []primitive.ObjectID `json:"my_projects"`
	WorkRequests []primitive.ObjectID `bson:"work_requests" json:"work_requests"`
	Verified     bool                 `json:"verified"`
	Created      time.Time            `json:"created_at"`
	Updated      time.Time            `json:"updated_at,omitempty"`
}

type UserContact struct {
	Email                   string `json:"email"`
	Phone                   string `json:"phone"`
	Country                 string `json:"country"`
	IdentifierDocumentation string `json:"identifier_document"`
}

type Profession struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type AuthUser struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
type Users []*User

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegistration struct {
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Password    string      `json:"password"`
	PerfilImage string      `json:"perfil_image"`
	Contact     UserContact `json:"contact"`
	Profession  Profession  `json:"profession"`
}
