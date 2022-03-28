package project_models

import (
	"time"
)

type Project struct {
	ID           interface{} `bson:"_id,omitempty" json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Propietarios interface{} `json:"propietarios"`
	Workers      interface{} `json:"workers"`
	WorkRequests interface{} `bson:"work_requests" json:"work_requests"`
	Color        string      `json:"color"`
	DateStarted  string      `json:"date_started"`
	DateFinished string      `json:"date_finished"`
	Created      time.Time   `json:"created_at"`
	Updated      time.Time   `json:"updated_at,omitempty"`
}

type Projects []*Project
