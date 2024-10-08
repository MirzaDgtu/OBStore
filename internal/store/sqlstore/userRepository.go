package sqlstore

import "obstore/internal/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u model.User) (model.User, error) {
	err := r.store.db.Create(&u).Error
	return u, err
}
