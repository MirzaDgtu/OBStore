package model

import "gorm.io/gorm"

type AssemblyOrderDetails struct {
	gorm.Model
	AssemblyOrderId int     `gorm:"column:assembly_order_id" json:"assembly_order_id"`
	Article         string  `gorm:"column:article" json:"article"`
	NameArticle     string  `gorm:"column:name_article" json:"name_article"`
	StrikeCode      string  `gorm:"column:strike_code" json:"strike_code"`
	Qty             float64 `gorm:"column:qty" json:"qty"`
	QtySbor         float64 `gorm:"column:qty_sbor" json:"qty_sbor"`
	Cena            float64 `gorm:"column:cena" json:"cena"`
	Discount        float64 `gorm:"column:discount" json:"discount"`
	SumArticle      float64 `gorm:"column:sum_article" json:"sum_article"`
	Srok            string  `gorm:"column:srok" json:"srok"`
	Partia          string  `gorm:"column:partia" json:"partia"`
	Mark            string  `gorm:"column:mark" json:"mark"`
}

func (AssemblyOrderDetails) TableName() string {
	return "assemblyorder_details"
}
