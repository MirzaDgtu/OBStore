package sqlstore

import "obstore/internal/model"

type AssemblyOrderRepository struct {
	store *Store
}

func (r *AssemblyOrderRepository) Create(u model.AssemblyOrder) (model.AssemblyOrder, error) {
	return u, r.store.db.Create(&u).Error
}

func (r *AssemblyOrderRepository) Update(u model.AssemblyOrder) (ao model.AssemblyOrder, err error) {
	return u, r.store.db.Save(&u).Error
}

func (r *AssemblyOrderRepository) Delete(id uint) error {
	return r.store.db.Delete(&model.AssemblyOrder{}, id).Error
}

func (r *AssemblyOrderRepository) ByID(id uint) (ao model.AssemblyOrder, err error) {
	return ao, r.store.db.Model(&model.AssemblyOrder{}).Preload("OrderDetails").First(&ao, id).Error
}

func (r *AssemblyOrderRepository) ByOrderID(orderid uint) (ao model.AssemblyOrder, err error) {
	return ao, r.store.db.Where("order_id=?", orderid).First(&ao).Error
}

func (r *AssemblyOrderRepository) ByUserID(userid uint) (aos []model.AssemblyOrder, err error) {
	return aos, r.store.db.Where("user_id=?", userid).Find(&aos).Error
}

func (r *AssemblyOrderRepository) All() (aos []model.AssemblyOrder, err error) {
	return aos, r.store.db.Find(&aos).Error
}

func (r *AssemblyOrderRepository) ByDateRange(dtStart, dtFinish string) (aos []model.AssemblyOrder, err error) {
	return aos, r.store.db.Where("date_doc between ? and ?", dtStart, dtFinish).Find(&aos).Error
}
