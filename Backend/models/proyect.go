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
	Workers      []string             `json:"workers"`
	Color        string               `json:"color"`
	DateStarted  string               `json:"date_started"`
	DateFinished string               `json:"date_finished"`
	Created      time.Time            `json:"created_at"`
	Updated      time.Time            `json:"updated_at,omitempty"`
}

type Proyectos []*Proyecto
