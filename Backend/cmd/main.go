package main

import (
	"fmt"
	"mongoCrud/routes"

	_ "mongoCrud/docs"

	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	routes.UsersRoute(app)
	routes.ProjectsRoute(app)
	fmt.Println("dd")
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
	app.Listen(":3000")
}
