package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
	"myInventory/helpers"
	"myInventory/models"
	"myInventory/mongo"
	"strconv"
	"time"
)

var productCollection = mongo.GetClient().Database("inventory").Collection("products")

var inventoryCollection = mongo.GetClient().Database("inventory").Collection("inventories")

func AddToInventory(c *fiber.Ctx) error {
	_inventoryId := c.Params("inventory_id")
	inventoryId, err := primitive.ObjectIDFromHex(_inventoryId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid inventory ID",
		})
	}
	var payload models.NewProduct
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
		})
	}
	newProduct := models.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Quantity:    payload.Quantity,
		Price:       payload.Price,
		InventoryID: inventoryId,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
	product, err := productCollection.InsertOne(c.Context(), newProduct)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	newProduct.ID = product.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    newProduct,
	})
}

func GetAllInventory(c *fiber.Ctx) error {
	inventoryId := c.Locals("inventoryId").(primitive.ObjectID)
	_limit := c.Query("limit", "20")
	_page := c.Query("page", "1")
	limit, _ := strconv.Atoi(_limit)
	page, _ := strconv.Atoi(_page)
	var products []models.Product
	cursor, err := productCollection.Find(c.Context(), bson.D{{"inventory_id", inventoryId}}, options.Find().SetLimit(int64(limit)), options.Find().SetSkip(int64((page-1)*limit)))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	if err := cursor.All(c.Context(), &products); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	count, err := productCollection.CountDocuments(c.Context(), bson.D{{"inventory_id", inventoryId}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	pages := math.Ceil(float64(count) / float64(limit))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    products,
		"pages":   pages,
	})
}

func GetProductById(c *fiber.Ctx) error {
	_id := c.Locals("productId").(primitive.ObjectID)
	var product models.Product
	err := productCollection.FindOne(c.Context(), bson.D{{"_id", _id}}).Decode(&product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	_id := c.Locals("productId").(primitive.ObjectID)
	_, err := helpers.ExistsProduct(_id, c)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product not found",
		})
	}
	var payload models.NewProduct
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
		})
	}
	update := bson.D{
		{
			"$set",
			bson.D{{
				"name", payload.Name,
			}, {
				"description", payload.Description,
			}, {
				"quantity", payload.Quantity,
			}, {
				"price", payload.Price,
			}, {
				"updated_at", time.Now().Format(time.RFC3339),
			}},
		},
	}
	updatedProduct, err := productCollection.UpdateOne(c.Context(), bson.D{{"_id", _id}}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    updatedProduct,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	_id := c.Locals("productId").(primitive.ObjectID)
	product, err := helpers.ExistsProduct(_id, c)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product not found",
		})
	}
	_, err = productCollection.DeleteOne(c.Context(), bson.D{{"_id", _id}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"product_id": product.ID,
		},
	})
}

func CreateInventory(c *fiber.Ctx) error {
	storeId := c.Params("store_id")
	var payload models.NewInventory
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
		})
	}
	inventory := models.Inventory{
		Name:        payload.Name,
		Description: payload.Description,
		Tags:        payload.Tags,
		StoreID:     storeId,
	}
	inserted, err := inventoryCollection.InsertOne(c.Context(), inventory)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	inventory.ID = inserted.InsertedID.(primitive.ObjectID).Hex()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    inventory,
	})
}

func GetAllInventories(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID
	var stores []models.Store
	var inventories []models.Inventory
	db.Where("user_id = ?", userId).Find(&stores)
	limitQuery := c.Query("limit", "20")
	skipQuery := c.Query("page", "1")
	limit, _ := strconv.ParseInt(limitQuery, 10, 64)
	skip, _ := strconv.ParseInt(skipQuery, 10, 64)
	storeIds := make([]string, len(stores))
	for i, store := range stores {
		storeIds[i] = strconv.FormatUint(uint64(store.ID), 10)
	}
	log.Println(storeIds)
	opts := options.Find().SetLimit(limit).SetSkip((skip - 1) * limit)
	filter := bson.D{{"storeId", bson.D{
		{"$in", storeIds},
	}},
	}
	log.Println(filter)
	cursor, err := inventoryCollection.Find(c.Context(), filter, opts)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})

	}
	if err := cursor.All(c.Context(), &inventories); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})

	}
	count, err := inventoryCollection.CountDocuments(c.Context(), bson.D{{"store_id", userId}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	pages := math.Ceil(float64(count) / float64(limit))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    inventories,
		"pages":   pages,
	})
}

func GetInventoryById(c *fiber.Ctx) error {
	inventoryId := c.Locals("inventoryId").(primitive.ObjectID)
	var inventory models.Inventory
	err := inventoryCollection.FindOne(c.Context(), bson.D{{"_id", inventoryId}}).Decode(&inventory)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    inventory,
	})
}

func UpdateInventory(c *fiber.Ctx) error {
	var payload models.NewInventory
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
		})
	}
	var inventory models.Inventory
	update := bson.D{
		{
			"$set",
			bson.D{{
				"name", payload.Name,
			}, {
				"description", payload.Description,
			}, {
				"tags", payload.Tags,
			}, {
				"updated_at", time.Now().Format(time.RFC3339),
			},
			},
		},
	}
	inventoryId := c.Locals("inventoryId").(primitive.ObjectID)
	err := inventoryCollection.FindOneAndUpdate(c.Context(), bson.D{{"_id", inventoryId}}, update).Decode(&inventory)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    inventory,
	})
}

func DeleteInventory(c *fiber.Ctx) error {
	inventoryId := c.Locals("inventoryId").(primitive.ObjectID)
	_, err := inventoryCollection.DeleteOne(c.Context(), bson.D{{"_id", inventoryId}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": fiber.Map{
			"id": inventoryId,
		},
	})

}
