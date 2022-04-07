package project_ports

import (
	models "zetting/pkg/project/core/models"
)

type ProjectService interface {
	CreateProject(newProject *models.Project, userId interface{}) (interface{}, error)
	GetProjects(userId interface{}, document string, page int) (models.Projects, error)
	AddUserToProject(addUserData models.AddUserToProject, document string) error
	GetProjectByProjectId(projectId interface{}) (*models.Project, error)
}
