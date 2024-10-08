package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Article    string  `gorm:"column:article" json:"article"`
	NameArtic  string  `gorm:"column:nameartic" json:"name_artic"`
	Unit       string  `gorm:"column:unit" json:"unit"`
	Fasovka    float32 `gorm:"column:fasovka" json:"dasovka"`
	TipTovr    string  `gorm:"column:tiptovr" json:"tip_tovr"`
	Maker      string  `gorm:"column:maker" json:"maker"`
	Price      float32 `gorm:"column:price" json:"price"`
	StrikeCode string  `gorm:"column:strikecode" json:"strikecode"`
	Code       string  `gorm:"column:code" json:"code"`
}

func (Product) TableName() string {
	return "products"
}
