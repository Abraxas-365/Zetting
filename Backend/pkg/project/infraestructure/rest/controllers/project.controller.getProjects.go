package project_controllers

import (
	"github.com/gofiber/fiber/v2"
	"zetting/auth"
)

func (s *projectController) GetMyProjects(c *fiber.Ctx) error {
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	myProjects, err := s.projectService.GetProjects(userTokenData.ID, "myprojects")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(myProjects)

}
func (s *projectController) GetProjectsWorkingOn(c *fiber.Ctx) error {
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	myProjects, err := s.projectService.GetProjects(userTokenData.ID, "myprojects")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(myProjects)

}
