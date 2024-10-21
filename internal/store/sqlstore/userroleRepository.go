package sqlstore

import "obstore/internal/model"

type UserRoleRepository struct {
	store *Store
}

func (r *UserRoleRepository) Create(u model.UserRole) (model.UserRole, error) {
	return u, r.store.db.Create(&u).Error
}

func (r *UserRoleRepository) Update(model.UserRole) (model.UserRole, error) {
	return model.UserRole{}, nil
}

func (r *UserRoleRepository) DeleteByID(uint) error {
	return nil
}

func (r *UserRoleRepository) DeleteByUserID(uint) error {
	return nil
}

func (r *UserRoleRepository) ByID(ID uint) (ur model.UserRole, err error) {
	return ur, r.store.db.First(&ur, ID).Error
}

func (r *UserRoleRepository) ByUserID(userID uint) (userroles []model.UserRole, err error) {
	return userroles, r.store.db.Where("user_id").Find(&userroles).Error
}

func (r *UserRoleRepository) ByRoleID(roleID uint) (userroles []model.UserRole, err error) {
	return userroles, r.store.db.Where("role_id", roleID).Find(&userroles).Error
}

func (r *UserRoleRepository) All() (userroles []model.UserRole, err error) {
	return userroles, r.store.db.Find(&userroles).Error
}
