package controllers

import (
	"github.com/gofiber/fiber/v2"
	"myInventory/helpers"
	"myInventory/models"
)

func CreateStore(c *fiber.Ctx) error {
	var storeIn models.StoreIn
	if err := c.BodyParser(&storeIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input",
		})
	}
	exits := helpers.ExistsStore(storeIn.Name, c.Locals("user").(models.User).ID)
	if exits {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Store already exists",
		})
	}
	store := models.Store{
		Name:        storeIn.Name,
		Description: storeIn.Description,
		Location:    storeIn.Location,
		UserID:      c.Locals("user").(models.User).ID,
	}
	if err := db.Create(&store).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create store",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": false,
		"data": fiber.Map{
			"storeId": store.ID,
		},
	})
}

func GetStoreById(c *fiber.Ctx) error {
	var store models.Store
	storeId := c.Params("store_id")
	if err := db.Where("id = ?", storeId).First(&store).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Store not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    store,
	})
}

func UpdateStore(c *fiber.Ctx) error {
	var storeIn models.StoreIn
	if err := c.BodyParser(&storeIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input",
		})
	}
	storeId := c.Params("store_id")
	var store models.Store
	if err := db.Where("id = ?", storeId).First(&store).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Store not found",
		})
	}
	store.Name = storeIn.Name
	store.Description = storeIn.Description
	store.Location = storeIn.Location
	if err := db.Save(&store).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update store",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    store,
	})
}

func DeleteStore(c *fiber.Ctx) error {
	var store models.Store
	storeId := c.Params("store_id")
	if err := db.Where("id = ?", storeId).First(&store).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Store not found",
		})
	}
	if err := db.Delete(&store).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete store",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    fiber.Map{"storeId": store.ID},
	})
}
