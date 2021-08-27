package models

import "gorm.io/gorm"

type Auto struct {
	gorm.Model
	BrandID   uint
	Brand     Brand
	ColorID   uint
	Color     Color
	ModelID   uint
	ModelItem Model `gorm:"foreignKey:ModelID"`
}
