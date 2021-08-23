package models

import (
	"gorm.io/gorm"
)


type Announce struct {
	AccountId uint
	Name string
	Description string
	gorm.Model
	CategoryId uint
	BrandId uint
	AutoId uint
	Price float64
}

func (a *Announce) AddAn() *Announce {
	GetDB().Create(a)
	return a
}

func (a *Announce) DelAn() {
	GetDB().Delete(a)
}

func (a *Announce) UpdAn(announce Announce) *Announce {
	GetDB().Model(a).Updates(announce)
	return a
}

func GetAnn() []Announce {
	announcements := []Announce{}
	GetDB().Find(&announcements)
	return announcements
}