package models

type Inventory struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description,omitempty"`
	Tags        []string `json:"tags" bson:"tags"`
	StoreID     uint     `json:"storeId" bson:"storeId"`
}

type NewInventory struct {
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description,omitempty"`
	Tags        []string `json:"tags" bson:"tags,omitempty"`
}
