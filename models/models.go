package models

import "gorm.io/gorm"

//type Model struct {
//	gorm.Model
//	Name       string
//	BrandID    uint
//	Brand      Brand
//	Autos      []Auto
//}
type Model struct {
	gorm.Model
	Name    string
	BrandID uint
	Brand   Brand
	Autos   []Auto
}
