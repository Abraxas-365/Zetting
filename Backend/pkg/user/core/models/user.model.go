package user_models

import (
	"time"
)

type User struct {
	ID           interface{} `bson:"_id,omitempty" json:"id"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Password     string      `json:"password"`
	PerfilImage  string      `json:"perfil_image"`
	Contact      Contact     `json:"contact"`
	Profession   Profession  `json:"profession"`
	Features     Features    `json:"features"`
	Projects     interface{} `json:"projects"`
	MyProjects   interface{} `json:"my_projects"`
	WorkRequests interface{} `bson:"work_requests" json:"work_requests"`
	Verified     bool        `json:"verified"`
	Created      time.Time   `json:"created_at"`
	Updated      time.Time   `json:"updated_at,omitempty"`
}

type Users []*User
