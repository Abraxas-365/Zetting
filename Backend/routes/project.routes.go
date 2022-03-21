package routes

import (
	"github.com/gofiber/fiber/v2"

	"mongoCrud/auth"
	controller "mongoCrud/controllers"
)

func ProjectsRoute(app *fiber.App) {

	project := app.Group("/api/projects")

	/*Create a new project*/
	project.Post("/new", auth.JWTProtected(), controller.CreateProject)

	/*Get my projects*/
	project.Get("/myprojects", auth.JWTProtected(), controller.GetMyProjects)

	/*get projects im in*/
	project.Get("/projects", auth.JWTProtected(), controller.GetProjectsImIn)

}
