package routes

import (
	"github.com/gofiber/fiber/v2"
	"myInventory/controllers"
	"myInventory/middleware"
)

func ReportsRoutes(app fiber.Router) fiber.Router {
	reportsRouter := app.Group("/reports")
	reportsRouter.Get("/inventory/:inventory_id", middleware.CanAccessInventory, controllers.GetInventoryReport)
	return reportsRouter
}
