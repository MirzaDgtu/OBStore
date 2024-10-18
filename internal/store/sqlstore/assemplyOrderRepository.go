package sqlstore

import "obstore/internal/model"

type AssemblyOrderRepository struct {
	store *Store
}

func (r *AssemblyOrderRepository) Create(u model.AssemblyOrder) (ao model.AssemblyOrder, err error) {
	return ao, r.store.db.Create(&u).Error
}

func (r *AssemblyOrderRepository) Update(u model.AssemblyOrder) (ao model.AssemblyOrder, err error) {
	return u, r.store.db.Save(&u).Error
}

func (r *AssemblyOrderRepository) Delete(id uint) error {
	return r.store.db.Delete(&model.AssemblyOrder{}, id).Error
}

func (r *AssemblyOrderRepository) GetByID(id uint) (ao model.AssemblyOrder, err error) {
	//return ao, r.store.db.Where("assemblyorder.id = ?", id).Preload("assemblyorder_details").First(&ao).Error
	return ao, r.store.db.Model(&model.AssemblyOrder{}).Preload("assemblyorder_details").First(&ao, id).Error
}

func (r *AssemblyOrderRepository) GetByOrderID(orderid uint) (ao model.AssemblyOrder, err error) {
	return ao, r.store.db.Where("order_id=?", orderid).First(&ao).Error
}

func (r *AssemblyOrderRepository) GetByUserID(userid uint) (aos []model.AssemblyOrder, err error) {
	return aos, r.store.db.Where("user_id=?", userid).Find(&aos).Error
}
