package user_routes

import (
	"zetting/auth"
	handler "zetting/pkg/user/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, handler handler.UserHandler) {
	/*SERVE*/
	static := app.Group("/static")
	static.Static("/app_default_images", "./static/app_default_images", fiber.Static{
		Browse: false, // default: false
	})
	users := app.Group("/api/users")
	/*Login user*/
	users.Post("/login", handler.LoginUser)
	/*Register user*/
	users.Post("/register", handler.CreateUser)
	/*Check email exist*/
	users.Get("/:email", handler.CheckEmailExist)
	/*Update user*/
	users.Put("/update", auth.JWTProtected(), handler.UpdateUser)
	/*Get user*/
	users.Get("/", auth.JWTProtected(), handler.GetUserById)
	/*Get users by profession*/
	users.Get("/profession/:profession/:page", auth.JWTProtected(), handler.GetUsersByProfession)
	/*Upload file image*/
	users.Put("/upload/perfil_image", auth.JWTProtected(), handler.UploadProfileImage)

}
