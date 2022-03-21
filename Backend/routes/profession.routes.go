package routes

import (
	controller "mongoCrud/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProfessionRoute(app *fiber.App) {

	profession := app.Group("/api/profession")

	profession.Get("/:profession", controller.GetByProfession)

	profession.Post("/:profession/filter", controller.GetByProfession)

}
