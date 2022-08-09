package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/golangdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Koneksi Gagal Database")
	}

	db.AutoMigrate(&Person{})

	return db
}
