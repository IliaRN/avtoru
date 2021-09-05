package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string
	Autos []Auto `gorm:"many2many:auto_categories;"`
}
