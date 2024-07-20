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

var productCollection = mongo.GetClient().Database("inventory").Collection("products")

func ExistsProduct(productId primitive.ObjectID, c *fiber.Ctx) (models.Product, error) {
	var product models.Product
	err := productCollection.FindOne(c.Context(), bson.D{{"_id", productId}}).Decode(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func CanAccessProduct(productId primitive.ObjectID, userId uint) (error, int) {
	var product models.Product
	err := productCollection.FindOne(context.TODO(), bson.D{{"_id", productId}}).Decode(&product)
	if err != nil {
		return err, fiber.StatusNotFound
	}
	inventoryId := product.InventoryID
	return CanAccessInventory(inventoryId, userId)
}

func GetProductById(_id primitive.ObjectID) (models.Product, error) {
	var product models.Product
	err := productCollection.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&product)
	if err != nil {
		log.Println(err)
		return product, err
	}
	return product, nil
}

func UpdateProductDataStruct(newData models.NewProduct, productId primitive.ObjectID) (bson.D, error) {
	var newProduct models.Product
	err := productCollection.FindOne(context.TODO(), bson.D{{"_id", productId}}).Decode(&newProduct)
	if err != nil {
		return nil, err
	}

	updateKeys := []string{"name", "description", "quantity", "price"}
	stringValidator := func(value interface{}) bool { return value != "" }
	intValidator := func(value interface{}) bool { return value != 0 }
	floatValidator := func(value interface{}) bool {
		_value, _ := value.(float64)
		return _value != 0.0
	}
	keysValidators := map[string]func(interface{}) bool{
		"name":        stringValidator,
		"description": stringValidator,
		"quantity":    intValidator,
		"price":       floatValidator,
	}
	newDataMap := map[string]interface{}{
		"name":        newData.Name,
		"description": newData.Description,
		"quantity":    newData.Quantity,
		"price":       newData.Price,
	}

	var update bson.D
	for _, key := range updateKeys {
		modelValue := newProduct.GetField(key)
		newValue := newDataMap[key]
		log.Printf("Key:%v , Model value: %v, New value: %v", key, modelValue, newValue)
		if !keysValidators[key](newValue) {
			continue
		}
		update = append(update, bson.E{Key: key, Value: newValue})
	}

	return update, nil
}
