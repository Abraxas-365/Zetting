package notification

import (
	service "zetting/pkg/notification/application"
	repo "zetting/pkg/notification/infraestructure/repository"
	handlers "zetting/pkg/notification/infraestructure/rest/handlers"
	routes "zetting/pkg/notification/infraestructure/rest/routes"

	"github.com/gofiber/fiber/v2"
)

func NotificationInit(app *fiber.App) {

	repo, _ := repo.NewMongoRepository("mongodb://localhost:27017", "Zetting", 10, "Notifications")
	service := service.NewNotificationService(repo)
	handlers := handlers.NewNotificationHandler(service)
	routes.NotificationRoute(app, handlers)
}
