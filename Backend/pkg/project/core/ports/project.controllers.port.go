package project_ports

import "github.com/gofiber/fiber/v2"

type ProjectController interface {
	CreateProject(c *fiber.Ctx) error
	GetMyProjects(c *fiber.Ctx) error
	GetProjectsWorkingOn(c *fiber.Ctx) error
}
