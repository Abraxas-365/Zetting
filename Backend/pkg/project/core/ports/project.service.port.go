package project_ports

import (
	models "zetting/pkg/project/core/models"
)

type ProjectService interface {
	CreateProject(newProject *models.Project, userId string) (string, error)
	GetProjects(userId string, document string) (models.Projects, error)
}
