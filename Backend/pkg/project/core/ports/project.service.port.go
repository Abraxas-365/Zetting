package project_ports

import (
	models "zetting/pkg/project/core/models"
)

type ProjectService interface {
	CreateProject(newProject *models.Project, userId interface{}) (interface{}, error)
	GetProjects(userId interface{}, document string) (models.Projects, error)
}
