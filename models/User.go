package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Stores   []Store
}

type UserIn struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
