package sqlstore

import "obstore/internal/model"

type AssemblyOrderRepository struct {
	store Store
}

func (r *AssemblyOrderRepository) Create(u model.AssemblyOrder) (ao model.AssemblyOrder, err error) {
	return ao, r.store.db.Create(&u).Error
}

func (r *AssemblyOrderRepository) Update(u model.AssemblyOrder) (ao model.AssemblyOrder, err error) {
	return u, r.store.db.Save(&u).Error
}

func (r *AssemblyOrderRepository) Delete(uint) error {

}

func (r *AssemblyOrderRepository) GetByID(uint) (ao model.AssemblyOrder, err error) {

}

func (r *AssemblyOrderRepository) GetByOrderID(uint) (ao model.AssemblyOrder, err error) {

}

func (r *AssemblyOrderRepository) GetByUserID(uint) (aos []model.AssemblyOrder, err error) {

}
