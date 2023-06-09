package main

import (
	"log"
	"middleware/database"
	"middleware/router"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error loading env file:", err)
	}

	log.Println("Env successfully loaded")
}

func main() {
	database.StartDB()
	router.StartApp().Run(":" + viper.GetString("PORT"))
}
