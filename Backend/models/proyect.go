package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Proyecto struct {
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

type Proyectos []*Proyecto

type WorkRequest struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProjectId primitive.ObjectID `bson:"project_id" json:"project_id"`
	OwnerId   primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	WorkerId  primitive.ObjectID `bson:"worker_id" json:"Worker_id"`
	Acepted   bool               `bson:"acepted" json:"acepted"`
	Created   time.Time          `json:"created_at"`
	Updated   time.Time          `json:"updated_at,omitempty"`
}

type WorkRequests []*WorkRequest
