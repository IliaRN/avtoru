package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	conn, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	err = db.Debug().AutoMigrate(&Account{}, &Announce{}, &Auto{}, &Brand{}, &Category{}, &Model{})
	if err != nil {
		panic("Can't connect to database")
	}
}

func GetDB() *gorm.DB {
	return db
}
