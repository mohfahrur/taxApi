package structs

import "github.com/jinzhu/gorm"

type Tax struct {
	gorm.Model
	Code int
}
