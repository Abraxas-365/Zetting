package project_controllers

import (
	ports "zetting/pkg/project/core/ports"

	"github.com/gofiber/fiber/v2"
)

type ProjectHandler interface {
	CreateProject(c *fiber.Ctx) error
	GetMyProjects(c *fiber.Ctx) error
	GetProjectsWorkingOn(c *fiber.Ctx) error
}
type projectHandler struct {
	projectService ports.ProjectService
}

func NewProjectHandler(projectService ports.ProjectService) ProjectHandler {
	return &projectHandler{
		projectService,
	}
}
