package workRequest_service

import (
	"errors"
	ports "zetting/pkg/workRequest/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

type workRequestService struct {
	projectRepo ports.WorkRequestRepository
}

func NewWorkRequestService(workRequestRepo ports.WorkRequestRepository) ports.WorkRequestService {
	return &workRequestService{
		workRequestRepo,
	}

}
