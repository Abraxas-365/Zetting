package workRequest_ports

import (
	models "zetting/pkg/workRequest/core/models"
)

type WorkRequestRepository interface {
	CreateWorkRequest(workRequest models.WorkRequest) error
	// GetWorkRequestByProject(projectId interface{}) (models.WorkRequests, error)
	// GetWorkRequestByWorker(workerId interface{}) (models.WorkRequests, error)
	// GetWorkRequestByOwner(ownerId interface{}) (models.WorkRequests, error)
}
