package models

import "gorm.io/gorm"

type Indicator struct {
	gorm.Model
	Name   string  `gorm:"column:name"`
	Ticker string  `gorm:"column:ticker"`                                // e.g) "DXY"
	Value  float64 `gorm:"column:value;type:decimal(10,2);default:0.00"` // "100.21"
	Type   string  `gorm:"column:type;default:Fiat"`                     // "Fiat" | "Crypto" | "ETF"
	Domain string  `gorm:"column:domain"`
}
