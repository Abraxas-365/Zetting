package workRequest_ports

import (
	models "zetting/pkg/workRequest/core/models"
)

type WorkRequestService interface {
	CreateWorkRequest(workRequest models.WorkRequest) error
	GetWorkRequests(id string, page int, document string) (models.WorkRequests, error)
	// GetWorkRequestByWorker(workerId interface{}) (models.WorkRequests, error)
	// GetWorkRequestByOwner(ownerId interface{}) (models.WorkRequests, error)
}
