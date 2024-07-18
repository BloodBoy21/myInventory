package helpers

import (
	"myInventory/models"
	"strconv"
)

func ExistsStore(storeId string, userId uint) bool {
	var store models.Store
	if err := db.Where("name = ?", storeId).Where("user_id=?", userId).First(&store).Error; err != nil {
		return false
	}
	return true
}

func CanAccessStore(storeIdStr string, userId uint) bool {
	var store models.Store
	storeId, _ := strconv.Atoi(storeIdStr)
	if err := db.Where("id = ?", storeId).Where("user_id=?", userId).First(&store).Error; err != nil {
		return false
	}
	return true
}
