package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myInventory/models"
	"myInventory/mongo"
)

var inventoryCollection = mongo.GetClient().Database("inventory").Collection("inventories")

func CanAccessInventory(inventoryId primitive.ObjectID, userId uint) bool {
	var inventory models.Inventory
	err := inventoryCollection.FindOne(context.TODO(), bson.D{{"_id", inventoryId}}).Decode(&inventory)
	if err != nil {
		return false
	}
	return CanAccessStore(inventory.StoreID, userId)
}
