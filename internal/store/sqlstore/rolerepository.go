package sqlstore

import "obstore/internal/model"

type RoleRepository struct {
	store *Store
}

func (r *RoleRepository) Create(u model.Role) (model.Role, error) {
	return u, r.store.db.Create(&u).Error
}

func (r *RoleRepository) Update(u model.Role) (model.Role, error) {
	return u, r.store.db.Where("id=?", u.ID).Update("name_role", u.NameRole).Error
}

func (r *RoleRepository) DeleteByID(id uint) error {
	return r.store.db.Where("id=?", id).Delete(&model.Role{}).Error
}

func (r *RoleRepository) ByID(id uint) (role model.Role, err error) {
	return role, r.store.db.First(&role, id).Error
}

func (r *RoleRepository) All() (roles []model.Role, err error) {
	return roles, r.store.db.Find(&roles).Error

}

func (r *RoleRepository) UsersByIdRole(id uint, roles *model.Role) error {
	err := r.store.db.Where("id = ?", id).Preload("Users").First(roles).Error
	if err != nil {
		return err
	}
	for i := range roles.Users {
		roles.Users[i].Pass = ""
		roles.Users[i].Token = ""
		roles.Users[i].RefreshToken = ""
	}

	return nil
}
