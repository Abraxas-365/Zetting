package user

import (
	"github.com/gofiber/fiber/v2"
	service "zetting/pkg/user/application"
	repo "zetting/pkg/user/infraestructure/repository"
	handler "zetting/pkg/user/infraestructure/rest/handlers"
	routes "zetting/pkg/user/infraestructure/rest/routes"
)

func UserInit(app *fiber.App) {

	repo, _ := repo.NewMongoRepository("mongodb://localhost:27017", "Zetting", 10, "Users")
	service := service.NewUsertService(repo)
	handler := handler.NewUserHandler(service)
	routes.UsersRoute(app, handler) //User routes
}
