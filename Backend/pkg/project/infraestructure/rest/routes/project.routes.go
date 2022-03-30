package project_routes

import (
	"zetting/auth"

	"github.com/gofiber/fiber/v2"
	handler "zetting/pkg/project/infraestructure/rest/handlers"
)

func ProjectsRoute(app *fiber.App, controller handler.ProjectController) {

	project := app.Group("/api/projects")

	/*Create a new project*/
	project.Post("/new", auth.JWTProtected(), controller.CreateProject)

	/*Get my projects*/
	project.Get("/myprojects", auth.JWTProtected(), controller.GetMyProjects)

	/*get projects im in*/
	project.Get("/projects", auth.JWTProtected(), controller.GetProjectsWorkingOn)

}
