package user_models

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
type Features struct {
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	// desde aqui es actores
	Height     int      `json:"height"`
	Body       string   `json:"body"`
	Skin       string   `json:"skin"`
	HairType   string   `json:"hair_type"`
	HairZise   string   `json:"hair_zise"`
	Skills     []string `json:"skills"`
	FacialHair string   `json:"facial_hair"`
	HairColor  string   `json:"hair_color"`
	EyeColor   string   `json:"eye_color"`
	// Aqui acaba actores
}

type Filter struct {
	AgeTop int `json:"age_top"`
	AgeLow int `json:"age_low"`
	// desde aqui es actores
	HeightTop  int      `json:"height_top"`
	HeightLow  int      `json:"height_low"`
	Body       string   `json:"body"`
	Skin       string   `json:"skin"`
	HairType   string   `json:"hair_type"`
	HairZise   string   `json:"hair_zise"`
	Skills     []string `json:"skills"`
	FacialHair string   `json:"facial_hair"`
	HairColor  string   `json:"hair_color"`
	EyeColor   string   `json:"eye_color"`
}
