package routes

import (
	"mongoCrud/auth"
	controller "mongoCrud/controllers"

	"github.com/gofiber/fiber/v2"
)

func WorkRequestRoute(app *fiber.App) {

	project := app.Group("/api/projects")

	project.Post("/add-work-request", auth.JWTProtected(), controller.CreateWorkRequest)

	project.Post("/respond-work-request", auth.JWTProtected(), controller.RespondWorkRequest)
	project.Get("/work-request", auth.JWTProtected(), func(c *fiber.Ctx) error {

		return nil
	})
}
