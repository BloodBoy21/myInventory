package routes

import (
	"github.com/gofiber/fiber/v2"
	"myInventory/controllers"
	"myInventory/middleware"
)

func ProductRoutes(app fiber.Router) fiber.Router {
	productRouter := app.Group("/product")
	productRouter.Get("/:product_id", middleware.CasAccessProduct, controllers.GetProductById)
	productRouter.Patch("/:product_id", middleware.CasAccessProduct, controllers.UpdateProduct)
	productRouter.Delete("/:product_id", middleware.CasAccessProduct, controllers.DeleteProduct)
	return productRouter
}
