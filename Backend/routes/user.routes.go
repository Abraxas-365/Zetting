package routes

import (
	"mongoCrud/auth"
	controller "mongoCrud/controllers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App) {

	users := app.Group("/api/users")
	/*Login user*/
	users.Post("/login", controller.LoginController)
	/*Register user*/
	users.Post("/register", controller.RegisterController)
	/*Check email exist*/
	users.Get("/:email", controller.CheckEmailExist)
	/*Get user*/
	users.Get("/", auth.JWTProtected(), controller.GetUser)
}
