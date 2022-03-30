package user_models

import (
	"time"
)

type User struct {
	ID           interface{} `bson:"_id,omitempty" json:"id"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name,omitempty"`
	Password     string      `json:"password,omitempty"`
	PerfilImage  string      `json:"perfil_image,omitempty"`
	Contact      Contact     `json:"contact,omitempty"`
	Profession   Profession  `json:"profession,omitempty"`
	Features     Features    `json:"features,omitempty"`
	Projects     interface{} `json:"projects,omitempty"`
	MyProjects   interface{} `json:"my_projects,omitempty"`
	WorkRequests interface{} `bson:"work_requests" json:"work_requests,omitempty"`
	Verified     bool        `json:"verified,omitempty"`
	Created      time.Time   `json:"created_at,omitempty"`
	Updated      time.Time   `json:"updated_at,omitempty"`
}

type Users []*User
