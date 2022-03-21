package controller

import (
	"mongoCrud/auth"
	"mongoCrud/models"
	service "mongoCrud/services"

	"github.com/gofiber/fiber/v2"
)

// Create Project.
// @CreateProject
// @Summary   crear project.
// @Tags      Project
// @Accept    json
// @Produce   json
// @Success   200
// @Param     project  body  models.ProyectCreator  true  "Project"
// @Security  ApiKeyAuth
// @Router    /project/new [post]
func CreateProject(c *fiber.Ctx) error {
	newProject := new(models.Proyecto)
	if err := c.BodyParser(newProject); err != nil {
		return fiber.ErrBadRequest
	}
	tokenData, err := auth.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	newProjectId, err := service.CreateProject(newProject, tokenData.ID)
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "correcto", "pid": newProjectId})
}

// GetMyProjects.
// @GetMyProjects
// @Summary   get my projects.
// @Tags      Project
// @Success   200
// @Security  ApiKeyAuth
// @Router    /project/myprojects [get]
func GetMyProjects(c *fiber.Ctx) error {

	tokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	myProjects, err := service.GetMyProjects(tokenData.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(myProjects)
}

// GetProjectsImIn.
// @GetProjectsImIn
// @Summary   get projects in my .
// @Tags      Project
// @Success   200
// @Security  ApiKeyAuth
// @Router    /project/projects [get]
func GetProjectsImIn(c *fiber.Ctx) error {
	tokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	GetProjectsImIn, err := service.GetProjectsWorkingOn(tokenData.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(GetProjectsImIn)

}
