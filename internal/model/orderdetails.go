package model

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetails struct {
	gorm.Model
	OrderUID    int       `gorm:"column:orderuid" json:"order_uid"`
	Articul     string    `gorm:"column:articul" json:"articul"`
	NameArticul string    `gorm:"column:namearticul" json:"name_articul"`
	Qty         float64   `gorm:"column:qty" json:"qty"`
	QtySbor     float64   `gorm:"column:qtySbor" json:"qty_sbor"`
	Cena        float64   `gorm:"column:cena" json:"cena"`
	Discount    float64   `gorm:"column:discount" json:"discount"`
	SumArtucul  float64   `gorm:"column:sumartucul" json:"sum_artucul"`
	FinishAt    time.Time `gorm:"column:finishat" json:"finish_at"`
	Srok        string    `gorm:"column:Srok" json:"Srok"`
	Partia      string    `gorm:"column:Partia" json:"Partia"`
	Done        bool      `gorm:"column:done" json:"done"`
}

func (OrderDetails) TableName() string {
	return "orderdetails"
}
