package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description,omitempty"`
	Location    string `json:"location" bson:"location"`
	UserID      uint   `json:"user_id" bson:"user_id"`
}

type StoreIn struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description,omitempty"`
	Location    string `json:"location" bson:"location,omitempty"`
}
