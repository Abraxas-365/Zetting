package project_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	Propietarios []primitive.ObjectID `json:"propietarios"`
	Workers      []primitive.ObjectID `json:"workers"`
	WorkRequests []primitive.ObjectID `bson:"work_requests" json:"work_requests"`
	Color        string               `json:"color"`
	DateStarted  string               `json:"date_started"`
	DateFinished string               `json:"date_finished"`
	Created      time.Time            `json:"created_at"`
	Updated      time.Time            `json:"updated_at,omitempty"`
}

type Projects []*Project

type ProyectCreator struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Color        string `json:"color"`
	DateStarted  string `json:"date_started"`
	DateFinished string `json:"date_finished"`
}
