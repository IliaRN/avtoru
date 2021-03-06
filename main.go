package main

import (
	"avtoru/app"
	"avtoru/controllers"
	"avtoru/models"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	app.AddRoute("/api/user/new", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodPost: controllers.CreateAccount,
	})
	app.AddRoute("/api/user/login", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodPost: controllers.Authenticate,
	})
	app.AddRoute("/api/user", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodGet: controllers.GetAccount,
	})
	app.AddRoute("/api/announcement", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodDelete: controllers.DelAn,
		http.MethodPut:    controllers.UpdAn,
		http.MethodPost:   controllers.AddAn,
		http.MethodGet:    controllers.GetAnnById,
	})
	app.AddRoute("/api/announcements/", map[string]func(w http.ResponseWriter, r *http.Request){
		http.MethodGet: controllers.GetAnns,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	db := models.GetDB()

	var brands = []string{"BMW", "Lada", "Hyundai", "Audi"}
	for _, brand := range brands {
		db.Create(&models.Brand{Name: brand})
	}

	var categories = []string{"Coupe", "Hatchback", "Crossover", "Sedan"}
	for _, category := range categories {
		db.Create(&models.Category{Name: category})
	}

	var modelItems = map[string][]string{"BMW": {"X4", "X5", "X6", "M5"},
		"Lada":    {"Vesta", "Xray", "Granta", "Largus"},
		"Hyundai": {"Solaris", "SantaFe", "Sonata", "Accent"},
		"Audi":    {"Q8", "A5", "R7", "TT"}}

	for brand, modRan := range modelItems {
		brandMod := &models.Brand{}
		db.First(&brandMod, "name = ?", brand)

		for _, model := range modRan {
			db.Create(&models.Model{BrandID: brandMod.ID, Name: model})
		}
	}

	fmt.Println(port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
