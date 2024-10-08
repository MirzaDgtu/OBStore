package store

import "obstore/internal/model"

type ProductRepository interface {
	Create(model.Product) (model.Product, error)
	GetByArticle(string) (model.Product, error)
	GetByStrikeCode(string) ([]model.Product, error)
	GetByName(string) ([]model.Product, error)
	GetAll() ([]model.Product, error)
}
