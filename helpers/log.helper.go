package helpers

import (
	"context"
	"myInventory/mongo"
)

var logCollection = mongo.GetClient().Database("logs").Collection("logs")

func SaveLog(log interface{}) error {
	_, err := logCollection.InsertOne(context.TODO(), log)
	return err
}
