package workRequest_models

import (
	"time"
)

//TODO validate status
//Status{
// "P"=pending
// "A"=accepted
// "D"=denied
// }
type WorkRequest struct {
	ID        interface{} `bson:"_id,omitempty" json:"id"`
	OwnerId   interface{} `bson:"owner_id" json:"owner_id"`
	ProjectId interface{} `bson:"project_id" json:"project_id"`
	WorkerId  interface{} `bson:"worker_id" json:"worker_id"`
	Status    string      `bson:"acepted" json:"acepted"`
	Created   time.Time   `json:"created_at"`
	Updated   time.Time   `json:"updated_at,omitempty"`
}

type WorkRequests []*WorkRequest
