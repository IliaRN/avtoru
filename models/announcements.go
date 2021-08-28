package models

import "gorm.io/gorm"

type Announce struct {
	gorm.Model
	AccountID   uint
	Account     Account
	Name        string
	Description string
	AutoID      uint
	Auto        Auto
	Price       float64
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
	//GetDB().Table("announces").Where("id = ?", u).First(ann)
	GetDB().Preload("Account").Find(ann, "id = ?", u)
	return ann
}

func GetAnns() []Announce {
	announcements := []Announce{}
	GetDB().Preload("Account").Find(&announcements)
	return announcements
}
func (a *Announce) DelAn(u uint) bool {
	//GetDB().Delete(a)
	//ann := &Announce{}
	ann := GetAnnById(u)
	GetDB().Delete(&Announce{}, u).First(ann) //.Table("announces").Where("id = ?", u).First(ann)
	return true

}
