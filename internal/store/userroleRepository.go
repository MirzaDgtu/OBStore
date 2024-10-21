package store

import "obstore/internal/model"

type UserRoleRepository interface {
	Create(model.UserRole) (model.UserRole, error)
	Update(model.UserRole) (model.UserRole, error)
	DeleteByID(uint) error
	DeleteByUserID(uint) error
	ByID(uint) (model.UserRole, error)
	ByUserID(uint) ([]model.UserRole, error)
	ByRoleID(uint) ([]model.UserRole, error)
	All() ([]model.UserRole, error)
}
