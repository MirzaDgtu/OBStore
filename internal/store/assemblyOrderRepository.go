package store

import "obstore/internal/model"

type AssemblyOrderRepository interface {
	Create(model.AssemblyOrder) (model.AssemblyOrder, error)
	Update(model.AssemblyOrder) (model.AssemblyOrder, error)
	DeleteByID(id uint)
}
