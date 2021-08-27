package seeds

import (
	"avtoru/models"
)

func main() {
	db := models.GetDB()

	var brands = []string{"BMW", "Lada", "Hyundai", "Audi"}
	for _, brand := range brands {
		db.Create(&models.Brand{Name: brand})
	}

	var categories = []string{"Coupe", "Hatchback", "Crossover", "Sedan"}
	for _, category := range categories {
		db.Create(&models.Category{Name: category})
	}

	var colors = []string{"Red", "White", "Black"}
	for _, color := range colors {
		db.Create(&models.Color{Name: color})
	}
}
