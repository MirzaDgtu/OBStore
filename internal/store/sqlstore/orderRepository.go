package sqlstore

import (
	"obstore/internal/model"
	"time"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(u model.Order) (model.Order, error) {
	err := r.store.db.Create(&u).Error

	return u, err
}

func (r *OrderRepository) Update(u model.Order) (order model.Order, err error) {
	return order, nil
}

func (r *OrderRepository) GetAll() (orders []model.Order, err error) {

	return orders, r.store.db.Table("orders").Select("*").
		Joins("orderdetails on orderdetails.orderuid = orders.orderuid").
		Scan(&orders).Error
}

func (r *OrderRepository) GetById(id int) (order model.Order, err error) {
	return order, r.store.db.Table("orders").Select("*").
		Joins("orderdetails on orderdetails.orderuid = orders.orderuid").
		Where("id = ?", id).
		Scan(&order).Error
}

func (r *OrderRepository) GetByOrderUID(uid int) (order model.Order, err error) {
	return order, r.store.db.Table("orders").Select("*").
		Joins("orderdetails on orderdetails.orderuid = orders.orderuid").
		Where("orderuid = ?", uid).
		Scan(&order).Error
}

func (r *OrderRepository) GetByFolioNum(folioNum int) (order model.Order, err error) {
	return order, r.store.db.Table("orders").Select("*").
		Joins("orderdetails on orderdetails.orderuid = orders.orderuid").
		Where("foliouid = ?", folioNum).
		Scan(&order).Error
}

func (r *OrderRepository) GetByDateRange(dtStart, dtFinish time.Time) (orders []model.Order, err error) {
	return orders, r.store.db.Table("orders").Select("*").
		Joins("orderdetails on orderdetails.orderuid = orders.orderuid").
		Where("foliodate BETWEEN ? AND ?", dtStart, dtFinish).
		Scan(&orders).Error
}

func (r *OrderRepository) GetByDriver(driver string) (orders []model.Order, err error) {
	return orders, r.store.db.Table("orders").Select("*").
		Joins("orderdetails on orderdetails.orderuid = orders.orderuid").
		Where("driver = ?", driver).
		Scan(&orders).Error
}

func (r *OrderRepository) GetByAgent(agent string) (orders []model.Order, err error) {
	return orders, r.store.db.Table("orders").Select("*").
		Joins("orderdetails on orderdetails.orderuid = orders.orderuid").
		Where("agent = ?", agent).
		Scan(&orders).Error
}
