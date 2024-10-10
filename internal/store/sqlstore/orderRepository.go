package sqlstore

import (
	"obstore/internal/model"
	"time"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(u model.Order) (order model.Order, err error) {
	err = r.store.db.Create(&u).Error

	return order, err
}

func (r *OrderRepository) DeleteById(id int) error {
	return nil
}

func (r *OrderRepository) Update(u model.Order) (order model.Order, err error) {
	return order, nil
}

func (r *OrderRepository) GetAll() (orders []model.Order, err error) {
	return orders, nil
}

func (r *OrderRepository) GetById(id int) (order model.Order, err error) {
	return order, nil
}

func (r *OrderRepository) GetByOrderUID(id int) (order model.Order, err error) {
	return order, nil
}

func (r *OrderRepository) GetByFolioNum(id int) (order model.Order, err error) {
	return order, nil
}

func (r *OrderRepository) GetByDateRange(dtStart, dtFinish time.Time) (orders []model.Order, err error) {
	return orders, nil
}

func (r *OrderRepository) GetByDriver(driver string) (orders []model.Order, err error) {
	return orders, nil
}

func (r *OrderRepository) GetByAgent(agent string) (orders []model.Order, err error) {
	return orders, nil
}
