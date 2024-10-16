package sqlstore

import (
	"obstore/internal/model"
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

	return orders, r.store.db.Preload("OrderDetails").Find(&orders).Error
}

func (r *OrderRepository) GetById(id int) (order model.Order, err error) {
	return order, r.store.db.Where("orders.id=?", id).Preload("OrderDetails").First(&order).Error
}

func (r *OrderRepository) GetByOrderUID(uid int) (order model.Order, err error) {
	return order, r.store.db.Where("orderuid=?", uid).Preload("OrderDetails").Find(&order).Error
}

func (r *OrderRepository) GetByFolioNum(folioNum int) (order model.Order, err error) {
	return order, r.store.db.Where("foliouid = ?", folioNum).Preload("OrderDetails").First(&order).Error
}

func (r *OrderRepository) GetByDateRange(dtStart, dtFinish string) (orders []model.Order, err error) {
	return orders, r.store.db.Where("foliodate BETWEEN ? AND ?", dtStart, dtFinish).
		Preload("OrderDetails").Find(&orders).Error
}

func (r *OrderRepository) GetByDriver(driver string) (orders []model.Order, err error) {
	return orders, r.store.db.Where("orders.Driver LIKE ?", "%"+driver+"%").Preload("OrderDetails").Find(&orders).Error
}

func (r *OrderRepository) GetByAgent(agent string) (orders []model.Order, err error) {
	return orders, r.store.db.Where("orders.Agent LIKE ?", "%"+agent+"%").Preload("OrderDetails").Find(&orders).Error
}
