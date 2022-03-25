package project_ports

import (
	models "zetting/pkg/project/core/models"
)

type ProjectRepository interface {
	CreateProject(project *models.Project, userId string) (string, error)
	GetProjects(userId string, document string) (models.Projects, error)
	AddProject(userId string, projectId string, campo string) error
}
