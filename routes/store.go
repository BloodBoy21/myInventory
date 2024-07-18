package routes

import (
	"github.com/gofiber/fiber/v2"
	"myInventory/controllers"
	"myInventory/middleware"
)

func StoreRoutes(app fiber.Router) fiber.Router {
	storesRouter := app.Group("/store")
	storesRouter.Post("/", controllers.CreateStore)
	storesRouter.Get("/:store_id", middleware.CanAccessStore, controllers.GetStoreById)
	storesRouter.Patch("/:store_id", middleware.CanAccessStore, controllers.UpdateStore)
	storesRouter.Delete("/:store_id", middleware.CanAccessStore, controllers.DeleteStore)
	return storesRouter
}
