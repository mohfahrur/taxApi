package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(db:3306)/db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Item{})
	db.AutoMigrate(structs.Tax{})
	return db
}
