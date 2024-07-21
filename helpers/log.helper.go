package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"log"
	"myInventory/models"
	"myInventory/mongo"
	"time"
)

var logCollection = mongo.GetClient().Database("logs").Collection("logs")

func SaveLog(log interface{}) error {
	_, err := logCollection.InsertOne(context.TODO(), log)
	return err
}

func GetInventoryLogsFromRange(inventoryId primitive.ObjectID, from, to time.Time) []models.ProductLog {
	var logs []models.ProductLog
	filter := bson.M{
		"inventoryId": inventoryId,
		"date": bson.M{
			"$gte": from,
			"$lte": to,
		},
	}
	log.Println(filter)
	cursor, err := logCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Error finding logs:", err)
		return logs
	}
	defer func(cursor *mongo2.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println("Error closing cursor:", err)
		}
	}(cursor, context.TODO()) // Ensure cursor is closed after use

	if err := cursor.All(context.TODO(), &logs); err != nil {
		log.Println("Error retrieving logs:", err)
		return logs
	}

	for i := range logs {
		var payload map[string]interface{}
		bsonBytes, err := bson.Marshal(logs[i].Payload)
		if err != nil {
			log.Println("Error marshaling payload:", err)
			continue
		}
		err = bson.Unmarshal(bsonBytes, &payload)
		if err != nil {
			log.Println("Error unmarshaling payload:", err)
			continue
		}

		logs[i].Payload = payload
	}
	return logs
}
