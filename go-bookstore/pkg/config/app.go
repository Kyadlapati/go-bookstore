package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "kranthi:bodapadu%2803@tcp(localhost:3306)/bookstore")

	if err != nil {
		panic(err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
