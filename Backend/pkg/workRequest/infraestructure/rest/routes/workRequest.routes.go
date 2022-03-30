package workRequest_routes

import (
	"zetting/auth"

	"github.com/gofiber/fiber/v2"
	handler "zetting/pkg/workRequest/infraestructure/rest/handlers"
)

func ProjectsRoute(app *fiber.App, handler handler.WorkRequestHandler) {

	project := app.Group("/api/work-request")

	/*Create a new project*/
	project.Post("/new", auth.JWTProtected(), handler.CreateWorkRequest)
	/*Get all the work request of a project*/
	project.Get("/project_id=\\:project_id", auth.JWTProtected(), handler.CreateWorkRequest)
	/*Get all the work request a user have resive*/
	project.Get("/worker_id=\\:worker_id", auth.JWTProtected(), handler.CreateWorkRequest)

}
