package main

import (
	"fmt"
	"mongoCrud/routes"

	_ "mongoCrud/docs"

	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title                       Zetting APi
// @version                     1.0
// @description                 API for Zetting
// @termsOfService              http://swagger.io/terms/
// @contact.name                API Support
// @contact.email               fiber@swagger.io
// @license.name                Apache 2.0
// @license.url                 http://www.apache.org/licenses/LICENSE-2.0.html
// @host                        localhost:3000
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath                    /api

func main() {

	app := fiber.New()
	app.Use(logger.New())
	//Routes.
	routes.UsersRoute(app) //User routes
	routes.ProjectsRoute(app)
	routes.ProfessionRoute(app)
	routes.ServeStatic(app)
	fmt.Println("inicando en puerto 3000")
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "swagger/doc.json",
		DeepLinking: false,
	}))
	app.Listen(":3000")
}
