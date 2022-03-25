package project

import (
	service "zetting/pkg/project/application"
	repo "zetting/pkg/project/infraestructure/repository"
	controller "zetting/pkg/project/infraestructure/rest/controllers"
	routes "zetting/pkg/project/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
)

func ProjectInit(app *fiber.App) {

	repo, _ := repo.NewMongoRepository("mongodb://localhost:27017", "Zetting", 10, "Projects")
	service := service.NewProjectService(repo)
	controller := controller.NewProjectController(service)
	routes.ProjectsRoute(app, controller) //User routes
}
