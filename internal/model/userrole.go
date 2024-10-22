package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	RoleID uint `gorm:"column:role_id" json:"role_id"`
	UserID uint `gorm:"column:user_id" json:"user_id"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
