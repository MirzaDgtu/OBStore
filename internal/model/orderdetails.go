package model

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetails struct {
	gorm.Model
	OrderId     int        `gorm:"column:orderid" json:"order_id"`
	OrderUID    int        `gorm:"column:orderuid" json:"order_uid"`
	Articul     string     `gorm:"column:articul" json:"articul"`
	NameArticul string     `gorm:"column:nameArticul" json:"name_articul"`
	Qty         float64    `gorm:"column:qty" json:"qty"`
	QtySbor     float64    `gorm:"column:qtySbor" json:"qty_sbor"`
	Cena        float64    `gorm:"column:cena" json:"cena"`
	Discount    float64    `gorm:"column:discount" json:"discount"`
	SumArtucul  float64    `gorm:"column:sumartucul" json:"sum_artucul"`
	FinishAt    *time.Time `gorm:"column:finishat" json:"finish_at"`
	Srok        *time.Time `gorm:"column:srok" json:"Srok"`
	Partia      string     `gorm:"column:partia" json:"Partia"`
	Done        bool       `gorm:"column:done" json:"done"`
}

func (OrderDetails) TableName() string {
	return "orderdetails"
}
