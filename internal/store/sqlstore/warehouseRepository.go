package sqlstore

import "obstore/internal/model"

type WarehouseRepository struct {
	store *Store
}

func (r *WarehouseRepository) Create(u model.Warehouse) (model.Warehouse, error) {
	return u, r.store.db.Create(&u).Error
}

func (r *WarehouseRepository) Update(u model.Warehouse) (warehouse model.Warehouse, err error) {
	err = r.store.db.Model(&u).Updates(map[string]interface{}{
		"name_warehouse": u.NameWarehouse,
	}).Error

	if err != nil {
		return model.Warehouse{}, err
	}

	// Загружаем обновленный объект из базы данных
	var updatedWarehouse model.Warehouse
	if err := r.store.db.First(&updatedWarehouse, u.ID).Error; err != nil {
		return model.Warehouse{}, err
	}

	return updatedWarehouse, nil
}

func (r *WarehouseRepository) GetAll() (warehouses []model.Warehouse, err error) {
	return warehouses, r.store.db.Find(&warehouses).Error
}

func (r *WarehouseRepository) GetByID(id uint) (warehouse model.Warehouse, err error) {
	return warehouse, r.store.db.First(&warehouse, id).Error
}

func (r *WarehouseRepository) DeleteByID(id uint) error {
	return r.store.db.Where("id=?", id).Delete(&model.Warehouse{}).Error
}
