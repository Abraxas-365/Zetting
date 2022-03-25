package project_controllers

import (
	ports "zetting/pkg/project/core/ports"
)

type projectController struct {
	projectService ports.ProjectService
}

func NewProjectController(projectService ports.ProjectService) ports.ProjectController {
	return &projectController{
		projectService,
	}
}
