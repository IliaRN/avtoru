package models

import (
	"avtoru/helpers"
	"gorm.io/gorm"
)

type Announce struct {
	gorm.Model
	AccountID   uint
	Account     Account
	Name        string
	Description string
	AutoID      uint
	Auto        Auto
	Price       float64 `json:"Price,float,omitempty"`
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
	GetDB().Preload("Auto").Find(ann, "id = ?", u)
	return ann
}

func GetAnns(fs helpers.FilterStruct) []Announce {
	var announcements []Announce

	query := GetDB().Preload("Auto.ModelItem.Brand").Preload("Auto.Categories").Model(&Announce{})
	flag := false
	if len(fs.Models) != 0 {
		query.Joins("LEFT JOIN `autos` ON `autos`.`id` = `announces`.`auto_id`")
		query.Where("`autos`.`model_id` IN ?", fs.Models)
		flag = true
	}
	if len(fs.Categories) != 0 {

		if !flag {
			query.Joins("LEFT JOIN `autos` ON `autos`.`id` = `announces`.`auto_id`")

			flag = true
		}
		query.Joins("LEFT JOIN `auto_categories` ON `auto_categories`.`auto_id` = `autos`. `id` ")
		query.Where("`auto_categories`.`category_id` IN ?", fs.Categories)
	}

	if len(fs.Brands) != 0 {
		if !flag {
			query.Joins("LEFT JOIN `autos` ON `autos`.`id` = `announces`.`auto_id`")
			flag = true
		}
		query.Joins("LEFT JOIN `models` ON `models`.`id` = `autos`.`model_id`")
		query.Joins("LEFT JOIN `brands` ON `brands`.`id` = `models`.`brand_id`")
		query.Where("`brands`.`id` IN ?", fs.Brands)
	}

	query.Find(&announcements)
	return announcements
}
func (a *Announce) DelAn(u uint) bool {
	ann := GetAnnById(u)
	GetDB().Delete(&Announce{}, u).First(ann) //.Table("announces").Where("id = ?", u).First(ann)
	return true

}
