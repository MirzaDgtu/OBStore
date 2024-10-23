package store

import (
	"obstore/internal/model"
)

type OrderRepository interface {
	Create(model.Order) (model.Order, error)
	Update(model.Order) (model.Order, error)
	All() ([]model.Order, error)
	ById(int) (model.Order, error)
	ByOrderUID(int) (model.Order, error)
	ByFolioNum(int) (model.Order, error)
	ByDateRange(string, string) ([]model.Order, error)
	ByDriver(string) ([]model.Order, error)
	ByAgent(string) ([]model.Order, error)
}
