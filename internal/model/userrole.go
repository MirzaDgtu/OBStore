package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	RoleID uint `json:"role_id"`
	UserID uint `json:"user_id"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
