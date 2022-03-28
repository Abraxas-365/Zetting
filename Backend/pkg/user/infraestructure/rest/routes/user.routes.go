package user_routes

import (
	"zetting/auth"
	handler "zetting/pkg/user/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, controller handler.UserController) {
	/*SERVE*/
	static := app.Group("/static")
	static.Static("/app_default_images", "./static/app_default_images", fiber.Static{
		Browse: false, // default: false
	})
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
	users.Get("/profession/:profession", auth.JWTProtected(), controller.GetUsersByProfession)
	/*Upload file image*/
	users.Put("/upload/perfil", auth.JWTProtected(), controller.UploadProfileImage)

}
