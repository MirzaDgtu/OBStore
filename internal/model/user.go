package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string `gorm:"column:firstname" json:"firstname" validate:"required"`
	LastName     string `gorm:"column:lastname" json:"lastname" validate:"required"`
	Email        string `gorm:"column:email" json:"email"`
	Pass         string `gorm:"column:pass" json:"pass" validate:"required"`
	LoggedIn     bool   `gorm:"column:loggedin" json:"loggedin"`
	Inn          string `gorm:"column:inn" json:"inn"`
	Token        string `gorm:"column:token" json:"token"`
	RefreshToken string `gorm:"column:refresh_token" json:"refresh_token"`
	Restore      bool   `gorm:"column:restore" json:"restore"`
	Blocked      bool   `gorm:"column:blocked" json:"blocked"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Teams        []Team `gorm:"many2many:teamcompositions;" json:"teams"`
	Roles        []Role `gorm:"many2many:user_roles;" json:"user_roles"`
}

func (User) TableName() string {
	return "users"
}
