package store

import "obstore/internal/model"

type RoleRepository interface {
	Create(model.Role) (model.Role, error)
	Update(model.Role) (model.Role, error)
	DeleteByID(uint) error
	ByID(uint) (model.Role, error)
	All() ([]model.Role, error)
	UsersByIdRole(uint, *model.Role) error
}
