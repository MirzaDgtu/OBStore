package store

import "obstore/internal/model"

type UserRepository interface {
	Create(model.User) (model.User, error)
	Update(model.User) (model.User, error)
	SignInUser(username, password string) (model.User, error)
	SignOutUserById(int) error
}
