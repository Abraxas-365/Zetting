package project_service

import models "zetting/pkg/project/core/models"

func (s *projectService) GetProjectByProjectId(projectId interface{}) (*models.Project, error) {
	project, err := s.projectRepo.GetProjectByProjectId(projectId)
	if err != nil {
		return nil, err
	}
	return project, nil

}
