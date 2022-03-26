package project_ports

import (
	models "zetting/pkg/project/core/models"
)

type ProjectRepository interface {
	CreateProject(project *models.Project, userId interface{}) (interface{}, error)
	GetProjects(userId interface{}, document string) (models.Projects, error)
	AddProject(userId interface{}, projectId interface{}, campo string) error
}
