package utils

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func GetEnv(key string) string {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file")

	}
	value, exists := os.LookupEnv(key)
	if !exists {
		return viper.GetString(key)
	}
	return value
}
