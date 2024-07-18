package helpers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myInventory/models"
	"myInventory/mongo"
)

var productCollection = mongo.GetClient().Database("inventory").Collection("products")

func ExistsProduct(productId primitive.ObjectID, c *fiber.Ctx) (models.Product, error) {
	var product models.Product
	err := productCollection.FindOne(c.Context(), bson.D{{"_id", productId}}).Decode(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func CanAccessProduct(productId primitive.ObjectID, userId uint) bool {
	var product models.Product
	err := productCollection.FindOne(context.TODO(), bson.D{{"_id", productId}}).Decode(&product)
	if err != nil {
		return false
	}
	inventoryId := product.InventoryID
	return CanAccessInventory(inventoryId, userId)
}
