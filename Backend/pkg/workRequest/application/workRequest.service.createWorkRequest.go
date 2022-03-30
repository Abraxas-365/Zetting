package workRequest_service

import models "zetting/pkg/workRequest/core/models"

func (r *workRequestService) CreateWorkRequest(newWorkRequest models.WorkRequest) error {

	if err := r.projectRepo.CreateWorkRequest(newWorkRequest); err != nil {
		return err
	}
	return nil

}
