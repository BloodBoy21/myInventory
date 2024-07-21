package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductLog struct {
	ID          string             `json:"id" bson:"_id,omitempty"`
	Action      string             `json:"action" bson:"action"`
	Type        string             `json:"type" bson:"type"`
	Payload     interface{}        `json:"payload" bson:"payload"`
	Date        primitive.DateTime `json:"date" bson:"date"`
	StoreId     uint               `json:"storeId" bson:"storeId"`
	InventoryId primitive.ObjectID `json:"inventoryId" bson:"inventoryId"`
}
