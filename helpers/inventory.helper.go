package helpers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"myInventory/models"
	"myInventory/mongo"
)

var inventoryCollection = mongo.GetClient().Database("inventory").Collection("inventories")

func CanAccessInventory(inventoryId primitive.ObjectID, userId uint) (error, int) {
	var inventory models.Inventory
	err := inventoryCollection.FindOne(context.TODO(), bson.D{{"_id", inventoryId}}).Decode(&inventory)
	if err != nil {
		log.Println(err)
		return err, fiber.StatusNotFound
	}
	return CanAccessStore(inventory.StoreID, userId)
}

func GetStoreIdFromInventory(inventoryId primitive.ObjectID) (uint, error) {
	var inventory models.Inventory
	err := inventoryCollection.FindOne(context.TODO(), bson.D{{"_id", inventoryId}}).Decode(&inventory)
	if err != nil {
		return 0, err
	}
	return inventory.StoreID, nil
}

func GetInventoryById(_id primitive.ObjectID) (models.Inventory, error) {
	var inventory models.Inventory
	err := inventoryCollection.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&inventory)
	if err != nil {
		log.Println(err)
		return inventory, err
	}
	return inventory, nil
}
