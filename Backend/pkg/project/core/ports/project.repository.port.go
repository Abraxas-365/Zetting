package project_ports

import (
	models "zetting/pkg/project/core/models"
)

type ProjectRepository interface {
	CreateProject(project *models.Project, userId interface{}) (interface{}, error)
	GetProjects(userId interface{}, document string, page int) (models.Projects, error)
}
