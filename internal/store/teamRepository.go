package store

import (
	"errors"
	"obstore/internal/model"
)

type TeamRepository interface {
	Create(model.Team) (model.Team, error)
	Update(model.Team) (model.Team, error)
	DeleteById(int) error
	GetById(int) (model.Team, error)
	GetAll() ([]model.Team, error)
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")
