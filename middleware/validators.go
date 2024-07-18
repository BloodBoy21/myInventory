package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myInventory/helpers"
	"myInventory/models"
)

func CanAccessStore(c *fiber.Ctx) error {
	storeId := c.Params("store_id")
	userId := c.Locals("user").(models.User).ID
	if !helpers.CanAccessStore(storeId, userId) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}
	return c.Next()
}

func CanAccessInventory(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID
	inventoryId := c.Params("inventory_id")
	_id, err := primitive.ObjectIDFromHex(inventoryId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid inventory ID",
		})
	}
	if !helpers.CanAccessInventory(_id, userId) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}
	c.Locals("inventoryId", _id)
	return c.Next()
}

func CasAccessProduct(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID
	productId := c.Params("product_id")
	_id, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid product ID",
		})
	}
	if !helpers.CanAccessProduct(_id, userId) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}
	c.Locals("productId", _id)
	return c.Next()
}
