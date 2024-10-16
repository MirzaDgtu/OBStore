package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"column:firstname" json:"firstname" validate:"required"`
	LastName  string `gorm:"column:lastname" json:"lastname" validate:"required"`
	Email     string `gorm:"column:email" json:"email" validate:"required"`
	Pass      string `gorm:"column:pass" json:"pass" validate:"required"`
	LoggedIn  bool   `gorm:"column:loggedin" json:"loggedin"`
	Inn       string `gorm:"column:inn" json:"inn"`
	Teams     []Team `gorm:"many2many:teamcompositions;" json:"teams"`
}

func (User) TableName() string {
	return "users"
}
