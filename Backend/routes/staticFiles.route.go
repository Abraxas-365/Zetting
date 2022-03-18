package routes

import "github.com/gofiber/fiber/v2"

func ServeStatic(app *fiber.App) {
	static := app.Group("/static")

	static.Static("/app_default_images", "./../Static/app_default_images", fiber.Static{
		Browse: false, // default: false
	})

}
