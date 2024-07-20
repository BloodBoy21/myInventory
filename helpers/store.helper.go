package helpers

import (
	"github.com/gofiber/fiber/v2"
	"myInventory/models"
)

func ExistsStore(storeId string, userId uint) bool {
	var store models.Store
	if err := db.Where("name = ?", storeId).Where("user_id=?", userId).First(&store).Error; err != nil {
		return false
	}
	return true
}

func CanAccessStore(storeId uint, userId uint) (error, int) {
	var store models.Store
	if err := db.Where("id = ?", storeId).Where("user_id=?", userId).First(&store).Error; err != nil {
		return err, fiber.StatusUnauthorized
	}
	return nil, fiber.StatusOK
}
