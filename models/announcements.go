package models

import (
	"gorm.io/gorm"
)

type Announce struct {
	AccountId   uint
	Name        string
	Description string
	gorm.Model
	CategoryId uint
	BrandId    uint
	AutoId     uint
	Price      float64
}

type UpdAnounce struct {
	Name        string
	Description string
}

func (a *Announce) AddAn() *Announce {
	GetDB().Create(a)
	return a
}

func (a *Announce) UpdAn(announce Announce) *Announce {
	GetDB().Model(a).Updates(announce)
	return a
}

func GetAnnById(u uint) *Announce {
	ann := &Announce{}
	GetDB().Table("announces").Where("id = ?", u).First(ann)
	return ann
}

func GetAnns() []Announce {
	announcements := []Announce{}
	GetDB().Find(&announcements)
	return announcements
}
func (a *Announce) DelAn(u uint) bool {
	//GetDB().Delete(a)
	//ann := &Announce{}
	ann := GetAnnById(u)
	GetDB().Delete(&Announce{}, u).First(ann) //.Table("announces").Where("id = ?", u).First(ann)
	return true

}
