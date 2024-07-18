package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"myInventory/controllers"
	"myInventory/mongo"
	"myInventory/routes"
	"myInventory/sql"
)

func main() {

	mongo.Connect()
	sql.Init()
	app := fiber.New()
	api := app.Group("/api/v1")
	app.Post("/login", controllers.LoginUser)
	app.Post("/signup", controllers.SingUpUser)
	routes.ApiRoutes(api)
	log.Fatal(app.Listen(":3000"))
}
