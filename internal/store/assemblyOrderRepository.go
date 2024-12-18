package store

import "obstore/internal/model"

type AssemblyOrderRepository interface {
	Create(model.AssemblyOrder) (model.AssemblyOrder, error)
	Update(model.AssemblyOrder) (model.AssemblyOrder, error)
	Delete(uint) error
	All() ([]model.AssemblyOrder, error)
	ByDateRange(dtStart, dtFinish string) ([]model.AssemblyOrder, error)
	ByID(uint) (model.AssemblyOrder, error)
	ByOrderID(uint) (model.AssemblyOrder, error)
	ByUserID(uint) ([]model.AssemblyOrder, error)
}
