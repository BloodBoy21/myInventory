package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description,omitempty"`
	Quantity    int                `json:"quantity" bson:"quantity,omitempty"`
	Price       float32            `json:"price" bson:"price,omitempty"`
	Available   bool               `json:"available" bson:"available"`
	InventoryID primitive.ObjectID `json:"inventory_id" bson:"inventory_id"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
}

type NewProduct struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description,omitempty"`
	Quantity    int     `json:"quantity" bson:"quantity,omitempty"`
	Price       float32 `json:"price" bson:"price,omitempty"`
}
