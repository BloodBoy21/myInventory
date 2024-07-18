package routes

import (
	"github.com/gofiber/fiber/v2"
	"myInventory/controllers"
	"myInventory/middleware"
)

func InventoryRoutes(app fiber.Router) fiber.Router {
	inventoryRouter := app.Group("/inventory")
	inventoryRouter.Post("/:store_id", middleware.CanAccessStore, controllers.CreateInventory)
	inventoryRouter.Post("/:inventory_id/product", middleware.CanAccessInventory, controllers.AddToInventory)
	inventoryRouter.Get("/:inventory_id", middleware.CanAccessInventory, controllers.GetInventoryById)
	inventoryRouter.Get("/:inventory_id/products", middleware.CanAccessInventory, controllers.GetAllInventory)
	inventoryRouter.Get("/", controllers.GetAllInventories)
	inventoryRouter.Patch("/:inventory_id", middleware.CanAccessInventory, controllers.UpdateInventory)
	inventoryRouter.Delete("/:inventory_id", middleware.CanAccessInventory, controllers.DeleteInventory)
	return inventoryRouter
}
