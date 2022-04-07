package workRequest_service

import models "zetting/pkg/workRequest/core/models"

func (r *workRequestService) CreateWorkRequest(newWorkRequest models.WorkRequest) (*models.WorkRequest, error) {

	WorkRequest, err := r.projectRepo.CreateWorkRequest(newWorkRequest)
	if err != nil {
		return nil, err
	}
	/*TODO: add adapter to send the id to the rabbitMQ*/

	return WorkRequest, nil

}
