package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	NameRole    string `json:"name_role"`
	Description string `json:"description" validation:"required"`
	Priority    int    `json:"priority"`
	User        []User `gorm:"many2many:user_roles;" json:"users"`
}

func (Role) TableName() string {
	return "roles"
}
