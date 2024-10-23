package store

import (
	"obstore/internal/model"
)

type UserRepository interface {
	Create(model.User) (model.User, error)
	Update(model.User) (model.User, error)
	SignInUser(username, password string) (model.User, error)
	SignOutUserById(int) error
	ChangePassword(int, string) error
	All() ([]model.User, error)
	UpdateToken(uint, string) error
	UserFromID(float64) (model.User, error)
	SetTemporaryPassword(string) (string, error)
}
