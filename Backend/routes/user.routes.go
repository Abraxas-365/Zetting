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
		fmt.Println("login")
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

		user := new(m.User)
		if err := c.BodyParser(user); err != nil {
			a := c.JSON(err)
			fmt.Println(a)
			return fiber.ErrBadRequest
		}
		fmt.Println(user.FirstName)
		fmt.Println("new user", user)

		authUser, err := user_service.CreateUser(*user)
		if err != nil {
			fmt.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(authUser)

	})
	/**/
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
	/**/
	users.Delete("/:nikname", func(c *fiber.Ctx) error {
		nikname := c.Params("nikname")
		if err := user_service.DeleteUser(nikname); err != nil {
			return c.Status(500).JSON(err)
		}
		return c.JSON("msg: eliminado correctamente")

	})
	/**/
	// users.Put("/:nikname", func(c *fiber.Ctx) error {
	// 	nikname := c.Params("nikname")
	// 	if err := user_service.DeleteUser(nikname); err != nil {
	// 		return c.Status(500).JSON(err)
	// 	}
	// 	return c.JSON("msg: modificado correctamente")

	// })
	// users.Put("/score", func(c *fiber.Ctx) error {
	// 	body := struct {
	// 		ID    string `json:"id"`
	// 		Score int    `json:"score"`
	// 	}{}
	// 	if err := c.BodyParser(&body); err != nil {
	// 		fmt.Println(err.Error())
	// 		return fiber.ErrBadRequest
	// 	}
	// 	fmt.Println(body.ID)

	// 	newScore, err := user_service.UpdateUserScore(body.ID, body.Score)
	// 	if err != nil {
	// 		return c.Status(500).JSON(err.Error())
	// 	}
	// 	body.Score = newScore
	// 	return c.JSON(body)
	// })

}
