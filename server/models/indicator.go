package models

import "gorm.io/gorm"

type Indicator struct {
	gorm.Model
	Name string `gorm:"unique"`
	Type string
}
