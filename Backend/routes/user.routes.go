package routes

import (
	"fmt"
	m "mongoCrud/models"
	user_service "mongoCrud/services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App) {

	users := app.Group("/api/users")

	users.Post("/login", func(c *fiber.Ctx) error {
		fmt.Println("---login route----")
		body := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			a := c.JSON(err)
			fmt.Println(a)
			return fiber.ErrBadRequest
		}

		fmt.Println("El email", body.Email)
		authUser, err := user_service.AuthUser(body.Email, body.Password)
		if err != nil {
			if err == fiber.ErrUnauthorized {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Bad Credentials"})
			}
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(authUser)

	})

	users.Post("/register", func(c *fiber.Ctx) error {
		fmt.Println("---Register Route---")
		body := struct {
			User m.User `json:"user"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			a := c.JSON(err)
			fmt.Println(a)
			return fiber.ErrBadRequest
		}
		fmt.Println(body)

		authUser, err := user_service.CreateUser(*&body.User)
		if err != nil {
			fmt.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(authUser)

	})

	users.Get("/:email", func(c *fiber.Ctx) error {
		email := c.Params("email")
		user, err := user_service.CheckUserExist(email)
		body := struct {
			Exists bool `json:"exists"`
		}{}

		if err != nil {
			fmt.Println(err.Error())
			if err.Error() == "mongo: no documents in result" {

				return c.JSON(body)
			}
			return c.SendStatus(fiber.StatusRequestTimeout)
		}
		return c.JSON(user)

	})

}
