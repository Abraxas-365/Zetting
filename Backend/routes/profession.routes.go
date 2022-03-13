package routes

import (
	"fmt"
	m "mongoCrud/models"
	service "mongoCrud/services"

	"github.com/gofiber/fiber/v2"
)

func ProfessionRoute(app *fiber.App) {

	profession := app.Group("/api/profession")

	profession.Get("/:profession", func(c *fiber.Ctx) error {

		profession := c.Params("profession")
		users, err := service.GetByProfession(profession)
		if err != nil {
			return fiber.ErrNotFound
		}
		return c.Status(fiber.StatusOK).JSON(users)

	})

	profession.Post("/filter", func(c *fiber.Ctx) error {
		body := struct {
			Profession string   `json:"profession"`
			Filter     m.Filter `json:"filter"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return fiber.ErrInternalServerError
		}

		fmt.Println(body.Filter)
		u, _ := service.FilterByFeatures(body.Profession, body.Filter)
		return c.Status(fiber.StatusOK).JSON(u)
	})

}
