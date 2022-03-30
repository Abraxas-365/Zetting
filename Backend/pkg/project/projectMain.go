package project

import (
	service "zetting/pkg/project/application"
	repo "zetting/pkg/project/infraestructure/repository"
	handlers "zetting/pkg/project/infraestructure/rest/handlers"
	routes "zetting/pkg/project/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
)

func ProjectInit(app *fiber.App) {

	repo, _ := repo.NewMongoRepository("mongodb://localhost:27017", "Zetting", 10, "Projects")
	service := service.NewProjectService(repo)
	handlers := handlers.NewProjectHandler(service)
	routes.ProjectsRoute(app, handlers) //User routes
}
