package store

import "obstore/internal/model"

type AssemblyOrderRepository interface {
	Create(model.AssemblyOrder) (model.AssemblyOrder, error)
	Update(model.AssemblyOrder) (model.AssemblyOrder, error)
	Delete(uint) error
	GetByID(uint) (model.AssemblyOrder, error)
	GetByOrderID(uint) (model.AssemblyOrder, error)
	GetByUserID(uint) ([]model.AssemblyOrder, error)
}
