package workRequest_ports

import (
	models "zetting/pkg/workRequest/core/models"
)

type WorkRequestRepository interface {
	CreateWorkRequest(workRequest models.WorkRequest) (*models.WorkRequest, error)
	GetWorkRequests(id string, page int, document string) (models.WorkRequests, error)
	AnswerWorkRequest(workRequest models.WorkRequest) error
}
