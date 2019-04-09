package structs

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name  string
	Price int
}
