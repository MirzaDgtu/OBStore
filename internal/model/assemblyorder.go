package model

import (
	"gorm.io/gorm"
)

type AssemblyOrder struct {
	gorm.Model
	DateDoc      string                 `gorm:"column:date_doc" json:"date_doc"`
	UserId       int                    `gorm:"column:user_id" json:"user_id"`
	StartAt      string                 `gorm:"column:start_at" json:"start_at"`
	FinishAt     string                 `gorm:"column:finish_at" json:"finish_at"`
	SumDoc       float32                `gorm:"column:sum_doc" json:"sum_doc"`
	ScladId      int                    `gorm:"column:sclad_id" json:"sclad_id"`
	VidDoc       string                 `gorm:"column:vid_doc" json:"vid_doc"`
	OrderId      int                    `gorm:"column:order_id" json:"order_id"`
	WeightDoc    float32                `gorm:"column:weight_doc" json:"weight_doc"`
	StatusId     int                    `gorm:"column:status_id" json:"status_id"`
	Done         bool                   `gorm:"column:done" json:"done"`
	OrderDetails []AssemblyOrderDetails `gorm:"foreignKey:AssemblyOrderId" json:"order_details"`
}

func (AssemblyOrder) TableName() string {
	return "assemblyorder"
}
