package routes

import (
	"github.com/gofiber/fiber/v2"
	"myInventory/middleware"
)

func ApiRoutes(app fiber.Router) {
	app.Use(
		middleware.IsAuthenticated,
	)
	inventoryRouter := InventoryRoutes(app)
	ProductRoutes(inventoryRouter)
	StoreRoutes(app)
	ReportsRoutes(app)

}
