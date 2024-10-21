package model

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	NameWarehouse string `gorm:"column:name_warehouse" json:"name_warehouse"`
}

func (Warehouse) TableName() string {
	return "warehouses"
}
