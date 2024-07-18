package sql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"myInventory/models"
	"myInventory/utils"
)

var client *gorm.DB

func Init() *gorm.DB {
	uri := utils.GetEnv("DB_URI")
	if uri == "" {
		log.Fatal("Set your 'DB_URI' environment variable")
	}
	pg := postgres.Open(uri)
	var err error                               // Change here to avoid redeclaration
	client, err = gorm.Open(pg, &gorm.Config{}) // Fixed redeclaration
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL!")
	//loadModels() // Corrected defer usage
	return client
}

func GetClient() *gorm.DB {
	if client == nil {
		log.Println("No client found. Connecting to PostgreSQL...")
		return Init()
	}
	return client
}

func loadModels() {
	err := client.AutoMigrate(&models.User{}, &models.Store{})
	if err != nil {
		log.Fatal("Error loading models:", err) // Added error detail
	}
}
