package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"myInventory/helpers"
	"time"
)

func GetInventoryReport(c *fiber.Ctx) error {
	location, _ := time.LoadLocation("America/Mexico_City")
	_from := c.Query("from")
	_to := c.Query("to")
	_inventoryId := c.Params("inventory_id")
	inventoryId, err := primitive.ObjectIDFromHex(_inventoryId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid inventory ID",
		})

	}
	invetory, err := helpers.GetInventoryById(inventoryId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Inventory not found",
		})
	}
	from := helpers.ParseDate(_from, location)
	to := helpers.ParseDate(_to, location)
	from = helpers.SetStartOfDay(from)
	to = helpers.SetEndOfDay(to)
	log.Printf("From: %v, To: %v", from, to)
	logs := helpers.GetInventoryLogsFromRange(inventoryId, from, to)
	fileName := helpers.GenerateProductLogReport(invetory.Name, logs)
	path := fmt.Sprintf("reports/%s", fileName)
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", "attachment; filename="+fileName)
	return c.SendFile(path)
}
