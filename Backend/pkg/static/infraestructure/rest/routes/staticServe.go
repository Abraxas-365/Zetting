package static_routes

import (
	"zetting/auth"
	controller "zetting/pkg/static/infraestructure/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func ServeStatic(app *fiber.App) {
	static := app.Group("/static")

	static.Static("/app_default_images", "./../Static/app_default_images", fiber.Static{
		Browse: true, // default: false
	})
	static.Post("/upload/perfil", auth.JWTProtected(), controller.HandleUploadProfileImage)

}
