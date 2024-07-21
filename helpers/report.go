package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myInventory/models"
	"strconv"
	"time"
)

// formatLogEntry is a helper function to format a log entry into the Excel file
func formatLogEntry(inventory, rowIndex string, data []byte, log models.ProductLog, productId string, f *excelize.File) error {
	dateString := time.Unix(int64(log.Date)/1000, 0).Format("2006-01-02 15:04:05")
	return f.SetSheetRow("Sheet1", rowIndex, &[]interface{}{
		inventory,
		log.Action,
		log.Type,
		productId,
		string(data), // Convert byte slice to string
		dateString,
	})
}

// getProductID extracts the product ID from the log payload
func getProductID(payload map[string]interface{}, key string) string {
	product := payload[key].(map[string]interface{})
	return product["_id"].(primitive.ObjectID).Hex()
}

// GenerateProductLogReport generates an Excel report for product logs
func GenerateProductLogReport(inventoryName string, data []models.ProductLog) string {
	f := excelize.NewFile()
	fileName := fmt.Sprintf("%s_product_log_%d.xlsx", inventoryName, time.Now().Unix())

	// Save the file after function execution
	defer func() {
		path := fmt.Sprintf("reports/%s", fileName)
		if err := f.SaveAs(path); err != nil {
			panic(err)
		}
	}()

	// Set header row
	err := f.SetSheetRow("Sheet1", "A1", &[]interface{}{"Inventory", "Action", "Type", "ProductId", "Data", "Date"})
	if err != nil {
		panic(err)
	}

	// Process each log entry
	for i, log := range data {
		row := i + 2
		rowIndex := "A" + strconv.Itoa(row)
		jsonData, err := json.Marshal(log.Payload)
		if err != nil {
			panic(err)
		}

		// Determine product ID based on log action
		var productId string
		if log.Action == "update" {
			productId = getProductID(log.Payload.(map[string]interface{}), "old")
		} else {
			productId = log.Payload.(map[string]interface{})["_id"].(primitive.ObjectID).Hex()
		}

		// Format the log entry
		err = formatLogEntry(inventoryName, rowIndex, jsonData, log, productId, f)
		if err != nil {
			panic(err)
		}
	}

	return fileName
}
