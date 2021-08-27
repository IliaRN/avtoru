package models

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name       string
	BrandID    uint
	Brand      Brand
	CategoryID uint
	Category   Category
	Autos      []Auto
}
