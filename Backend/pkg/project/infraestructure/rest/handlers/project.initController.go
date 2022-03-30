package project_controllers

import (
	ports "zetting/pkg/project/core/ports"

	"github.com/gofiber/fiber/v2"
)

type ProjectController interface {
	CreateProject(c *fiber.Ctx) error
	GetMyProjects(c *fiber.Ctx) error
	GetProjectsWorkingOn(c *fiber.Ctx) error
}
type projectController struct {
	projectService ports.ProjectService
}

func NewProjectController(projectService ports.ProjectService) ProjectController {
	return &projectController{
		projectService,
	}
}
