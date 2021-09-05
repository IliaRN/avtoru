package models

import "gorm.io/gorm"

type Auto struct {
	gorm.Model
	Mileage    uint
	Year       uint
	ModelID    uint
	ModelItem  Model      `gorm:"foreignKey:ModelID"`
	Categories []Category `gorm:"many2many:auto_categories;"`
}

func (a *Auto) AddAuto() *Auto {
	GetDB().Create(a)
	return a
}
