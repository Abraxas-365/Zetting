package workRequest_service

import models "zetting/pkg/workRequest/core/models"

func (r *workRequestService) GetWorkRequests(userId interface{}, page int, document string) (models.WorkRequests, error) {

	workRequests, err := r.projectRepo.GetWorkRequests(userId, page, document)
	if err != nil {
		return nil, err
	}
	return workRequests, nil

}
