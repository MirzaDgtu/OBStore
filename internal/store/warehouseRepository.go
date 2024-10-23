package store

import "obstore/internal/model"

type WarehouseRepository interface {
	Create(model.Warehouse) (model.Warehouse, error)
	Update(model.Warehouse) (model.Warehouse, error)
	DeleteByID(uint) error
	All() ([]model.Warehouse, error)
	ByID(uint) (model.Warehouse, error)
}
