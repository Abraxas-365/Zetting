package main

import (
	"fmt"
	_ "zetting/docs"

	project "zetting/pkg/project"
	user "zetting/pkg/user"
	workRequest "zetting/pkg/workRequest"

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
	workRequest.WorkRequestInit(app)
	fmt.Println("inicando en puerto 3000")
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "swagger/doc.json",
		DeepLinking: false,
	}))

	app.Listen(":3000")
}
