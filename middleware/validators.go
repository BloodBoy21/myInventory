package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myInventory/helpers"
	"myInventory/models"
	"strconv"
)

func CanAccessStore(c *fiber.Ctx) error {
	_storeId := c.Params("store_id")
	storeId, err := strconv.Atoi(_storeId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid store ID",
		})
	}
	userId := c.Locals("user").(models.User).ID
	if err, statusCode := helpers.CanAccessStore(uint(storeId), userId); err != nil {
		return c.Status(statusCode).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
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
	if err, statusCode := helpers.CanAccessInventory(_id, userId); err != nil {
		return c.Status(statusCode).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
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
	if err, statusCode := helpers.CanAccessProduct(_id, userId); err != nil {
		return c.Status(statusCode).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	c.Locals("productId", _id)
	return c.Next()
}
