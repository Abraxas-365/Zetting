package project_service

import (
	models "zetting/pkg/project/core/models"
)

func (s *projectService) GetProjects(userId interface{}, document string) (models.Projects, error) {
	projects, err := s.projectRepo.GetProjects(userId, document)
	if err != nil {
		return nil, err
	}
	return projects, nil

}
