package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	conn, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Announce{}, &Auto{}, &Brand{}, &Category{}, &Model{})
}

func GetDB() *gorm.DB {
	return db
}
