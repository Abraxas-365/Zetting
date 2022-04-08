package workRequest_service

import models "zetting/pkg/workRequest/core/models"

func (r *workRequestService) GetWorkRequest(workRequestId interface{}) (*models.WorkRequest, error) {

	workRequest, err := r.projectRepo.GetWorkRequest(workRequestId)
	if err != nil {
		return nil, err
	}
	return workRequest, nil

}
