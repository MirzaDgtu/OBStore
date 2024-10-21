package store

import "obstore/internal/model"

type WarehouseRepository interface {
	Create(model.Warehouse) (model.Warehouse, error)
	Update(model.Warehouse) (model.Warehouse, error)
	DeleteByID(uint) error
	GetAll() ([]model.Warehouse, error)
	GetByID(uint) (model.Warehouse, error)
}
