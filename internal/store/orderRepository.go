package store

import (
	"obstore/internal/model"
	"time"
)

type OrderRepository interface {
	Create(model.Order) (model.Order, error)
	DeleteById(int) error
	Update(model.Order) (model.Order, error)
	GetAll() ([]model.Order, error)
	GetById(int) (model.Order, error)
	GetByOrderUID(int) (model.Order, error)
	GetByFolioNum(int) (model.Order, error)
	GetByDateRange(time.Time, time.Time) ([]model.Order, error)
	GetByDriver(string) ([]model.Order, error)
	GetByAgent(string) ([]model.Order, error)
}
