package user_routes

import (
	"zetting/auth"
	ports "zetting/pkg/user/core/ports"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, controller ports.UserController) {

	users := app.Group("/api/users")
	/*Login user*/
	users.Post("/login", controller.LoginUser)
	/*Register user*/
	users.Post("/register", controller.CreateUser)
	/*Check email exist*/
	users.Get("/:email", controller.CheckEmailExist)
	/*Get user*/
	users.Get("/", auth.JWTProtected(), controller.GetUserById)
	/*Get users by profession*/
	users.Get("/:profession", auth.JWTProtected(), controller.GetUsersByProfession)

}
