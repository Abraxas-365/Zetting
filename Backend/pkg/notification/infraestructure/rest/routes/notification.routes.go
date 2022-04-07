package notification_routes

import (
	"zetting/auth"

	"github.com/gofiber/fiber/v2"
	handler "zetting/pkg/notification/infraestructure/rest/handlers"
)

func NotificationRoute(app *fiber.App, handler handler.NotificationsHandler) {

	notification := app.Group("/api/notification")

	/*Create a new notification*/
	notification.Post("/new/type=:type", auth.JWTProtected(), handler.CreateNotification)
	notification.Get("/page=:page", auth.JWTProtected(), handler.GetNotifications)

}
