package project_service

import (
	models "zetting/pkg/project/core/models"
)

func (s *projectService) CreateProject(newProject *models.Project, userId string) (string, error) {

	projectId, err := s.projectRepo.CreateProject(newProject, userId)
	if err != nil {
		return "", err
	}

	if err := s.projectRepo.AddProject(userId, projectId, "myprojects"); err != nil {
		return "", err
	}
	return projectId, nil

}
