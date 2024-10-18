package model

import "gorm.io/gorm"

type AssemblyOrderDetails struct {
	gorm.Model
	IdAccemblyOrder int     `gorm:"column:id_accembly_order" json:"id_accembly_order"`
	Articul         string  `gorm:"column:articul" json:"articul"`
	NameArticul     string  `gorm:"column:name_articul" json:"name_articul"`
	StrikeCode      string  `gorm:"column:strike_code" json:"strike_code"`
	Qty             float64 `gorm:"column:qty" json:"qty"`
	QtySbor         float64 `gorm:"column:qty_sbor" json:"qty_sbor"`
	Cena            float64 `gorm:"column:cena" json:"cena"`
	Discount        float64 `gorm:"column:discount" json:"discount"`
	SumArtucul      float64 `gorm:"column:sum_artucul" json:"sum_artucul"`
	Srok            string  `gorm:"column:srok" json:"Srok"`
	Partia          string  `gorm:"column:partia" json:"Partia"`
}

func (AssemblyOrderDetails) TableName() string {
	return "assemblyorder_details"
}
