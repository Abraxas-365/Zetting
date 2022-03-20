package routes

import (
	"fmt"
	"mongoCrud/auth"
	service "mongoCrud/services"

	"github.com/gofiber/fiber/v2"
)

func WorkRequestRoute(app *fiber.App) {

	project := app.Group("/api/projects")

	//agregar un trabajador al proyecto

	project.Post("/add-work-request", auth.JWTProtected(), func(c *fiber.Ctx) error {
		body := struct {
			ProjectId string `json:"project_id"`
			WorkerId  string `json:"worker_id"`
		}{}

		if err := c.BodyParser(&body); err != nil {
			fmt.Println(err)
			return fiber.ErrBadRequest
		}
		u, err := auth.ExtractTokenMetadata(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if err := service.AddWorkRequest(u.ID, body.ProjectId, body.WorkerId); err != nil {
			return err
		}
		return nil
	})

	project.Post("/respond-work-request", auth.JWTProtected(), func(c *fiber.Ctx) error {
		body := struct {
			WorkRequestId string `json:"work_request_id"`
			Accept        bool   `json:"accept"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			fmt.Println(err)
			return fiber.ErrBadRequest
		}
		u, err := auth.ExtractTokenMetadata(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if err := service.AcceptWorkRequest(body.WorkRequestId, u.ID, body.Accept); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusOK)

	})
	project.Get("/work-request", auth.JWTProtected(), func(c *fiber.Ctx) error {

		return nil
	})
}
