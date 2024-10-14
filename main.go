package main

import (
	"booth-reviews/controllers"
	"booth-reviews/models"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	databaseSetup()
	http.HandleFunc("/", controllers.AssetListHandler)
}

func databaseSetup() {
	db, err := gorm.Open(sqlite.Open("boothreviews.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	//Auto-migrate the schema
	db.AutoMigrate(&models.Asset{}, &models.Review{}, &models.Rating{})
}
