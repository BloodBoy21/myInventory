package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"myInventory/utils"
)

var client *mongo.Client

func Connect() *mongo.Client {
	uri := utils.GetEnv("MONGO_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}

func GetClient() *mongo.Client {
	if client == nil {
		log.Println("No client found. Connecting to MongoDB...")
		return Connect()
	}
	return client
}
