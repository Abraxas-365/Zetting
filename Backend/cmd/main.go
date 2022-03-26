package main

import (
	"fmt"
	_ "zetting/docs"

	project "zetting/pkg/project"
	static "zetting/pkg/static/infraestructure/rest/routes"
	user "zetting/pkg/user"

	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	//Routes.
	user.UserInit(app)
	project.ProjectInit(app)
	static.ServeStatic(app)
	fmt.Println("inicando en puerto 3000")
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "swagger/doc.json",
		DeepLinking: false,
	}))

	app.Listen(":3000")
}
