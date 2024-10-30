package model

import (
	"gorm.io/gorm"
)

type OrderDetails struct {
	gorm.Model
	OrderId     int     `gorm:"column:orderid" json:"order_id"`
	OrderUID    int     `gorm:"column:order_uid" json:"order_uid"`
	Article     string  `gorm:"column:article" json:"article"`
	NameArticle string  `gorm:"column:name_article" json:"name_article"`
	Strikecode  string  `gorm:"column:strikecode" json:"strikecode"`
	Qty         float64 `gorm:"column:qty" json:"qty"`
	Cena        float64 `gorm:"column:cena" json:"cena"`
	Discount    float64 `gorm:"column:discount" json:"discount"`
	ArticleSum  float64 `gorm:"column:article_sum" json:"article_sum"`
	FinishAt    string  `gorm:"column:finish_at" json:"finish_at"`
	Srok        string  `gorm:"column:srok" json:"srok"`
	Partia      string  `gorm:"column:partia" json:"partia"`
	Marka       string  `gorm:"column:marka" json:"marka"`
	Done        bool    `gorm:"column:done" json:"done"`
}

func (OrderDetails) TableName() string {
	return "orderdetails"
}
