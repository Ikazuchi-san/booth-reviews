package controllers

import (
	"booth-reviews/models"
	"log"
	"net/http"
	"text/template"

	"gorm.io/gorm"
)

func AssetListHandler(w http.ResponseWriter, r *http.Request) {

	var db *gorm.DB
	var assets []models.Asset

	//fetch assets from the database
	result := db.Find(&assets) // improve this later
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// parse and render the html template
	tmpl := template.Must(template.ParseFiles("templates/asset_list.html"))
	err := tmpl.Execute(w, assets)
	if err != nil {
		log.Println("Template execution error: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
