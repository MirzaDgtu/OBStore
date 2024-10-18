package model

import (
	"time"

	"gorm.io/gorm"
)

type AssemblyOrder struct {
	gorm.Model
	DateDoc      *time.Time             `gorm:"column:date_doc" json:"date_doc"`
	UserId       int                    `gorm:"column:user_id" json:"user_id"`
	StartAt      *time.Time             `gorm:"column:start_at" json:"start_at"`
	FinishAt     *time.Time             `gorm:"column:finish_at" json:"finish_at"`
	SumDoc       float32                `gorm:"column:sum_doc" json:"sum_doc"`
	IdSclad      int                    `gorm:"column:id_sclad" json:"id_sclad"`
	VidDoc       string                 `gorm:"column:vid_doc" json:"vid_doc"`
	IdOrder      int                    `gorm:"column:id_order" json:"id_order"`
	WeightDoc    float32                `gorm:"column:weight_doc" json:"weight_doc"`
	IdStatus     int                    `gorm:"column:id_status" json:"id_status"`
	Done         bool                   `gorm:"column:done" json:"done"`
	OrderDetails []AssemblyOrderDetails `gorm:"foreignKey:id_accemblyorder" json:"order_details"`
}

func (AssemblyOrder) TableName() string {
	return "assemblyorder"
}
