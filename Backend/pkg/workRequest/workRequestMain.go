package workrequest

import (
	service "zetting/pkg/workRequest/application"
	repo "zetting/pkg/workRequest/infraestructure/repository"
	handlers "zetting/pkg/workRequest/infraestructure/rest/handlers"
	routes "zetting/pkg/workRequest/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
)

func WorkRequestInit(app *fiber.App) {

	repo, _ := repo.NewMongoRepository("mongodb://localhost:27017", "Zetting", 10, "WorkRequests")
	service := service.NewWorkRequestService(repo)
	handlers := handlers.NewWorkRequestHandler(service)
	routes.ProjectsRoute(app, handlers) //User routes
}
