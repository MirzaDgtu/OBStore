package store

import (
	"obstore/internal/model"
)

type OrderRepository interface {
	Create(model.Order) (model.Order, error)
	Update(model.Order) (model.Order, error)
	GetAll() ([]model.Order, error)
	GetById(int) (model.Order, error)
	GetByOrderUID(int) (model.Order, error)
	GetByFolioNum(int) (model.Order, error)
	GetByDateRange(string, string) ([]model.Order, error)
	GetByDriver(string) ([]model.Order, error)
	GetByAgent(string) ([]model.Order, error)
}
