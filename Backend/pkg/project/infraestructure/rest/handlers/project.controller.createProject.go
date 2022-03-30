package project_controllers

import (
	"github.com/gofiber/fiber/v2"
	"zetting/auth"
	models "zetting/pkg/project/core/models"
)

func (s *projectHandler) CreateProject(c *fiber.Ctx) error {
	userTokenData, err := auth.ExtractTokenMetadata(c)
	createProjectData := new(models.Project)
	if err := c.BodyParser(&createProjectData); err != nil {
		return fiber.ErrBadRequest
	}
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	newProjectId, err := s.projectService.CreateProject(createProjectData, userTokenData.ID)
	if err != nil {
		return c.SendStatus(fiber.ErrConflict.Code)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "correcto", "pid": newProjectId})

}
