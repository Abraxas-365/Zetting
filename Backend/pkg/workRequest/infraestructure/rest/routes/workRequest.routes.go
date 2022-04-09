package workRequest_routes

import (
	"zetting/auth"

	"github.com/gofiber/fiber/v2"
	handler "zetting/pkg/workRequest/infraestructure/rest/handlers"
)

func WorkRequestRoute(app *fiber.App, handler handler.WorkRequestHandler) {

	workRequest := app.Group("/api/work-request")

	/*Create a new project*/
	workRequest.Post("/new", auth.JWTProtected(), handler.CreateWorkRequest)
	/*Get all the work request of a project*/
	workRequest.Get("/project_id=:project_id/status=:status/page=:page/number=:number", auth.JWTProtected(), handler.GetWorkRequestsByProject)
	/*Get all the work request a user have resive*/
	workRequest.Get("/worker_id=:worker_id/status=:status/page=:page/number=:number", auth.JWTProtected(), handler.GetWorkRequestsByWorker)
	/*Answer a work request*/
	workRequest.Post("/answer", auth.JWTProtected(), handler.AnswerWorkRequest)
	/*Get work request by id*/
	workRequest.Get("/id=:id", auth.JWTProtected(), handler.GetWorkRequestsById)
}
