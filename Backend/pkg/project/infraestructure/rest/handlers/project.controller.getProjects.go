package project_controllers

import (
	"strconv"
	"zetting/auth"

	"github.com/gofiber/fiber/v2"
)

func (s *projectController) GetMyProjects(c *fiber.Ctx) error {
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	page, _ := strconv.Atoi(c.Params("page"))
	myProjects, err := s.projectService.GetProjects(userTokenData.ID, "owner", page)
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

	page, _ := strconv.Atoi(c.Params("page"))
	myProjects, err := s.projectService.GetProjects(userTokenData.ID, "workers", page)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(myProjects)

}
