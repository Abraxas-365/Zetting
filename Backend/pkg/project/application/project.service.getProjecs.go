package project_service

import (
	models "zetting/pkg/project/core/models"
)

func (s *projectService) GetProjects(userId interface{}, document string, page int) (models.Projects, error) {
	projects, err := s.projectRepo.GetProjects(userId, document, page)
	if err != nil {
		return nil, err
	}
	return projects, nil

}
