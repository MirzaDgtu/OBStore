package store

import "obstore/internal/model"

type ProductRepository interface {
	Create(model.Product) (model.Product, error)
	ByArticle(string) (model.Product, error)
	ByStrikeCode(string) ([]model.Product, error)
	ByName(string) ([]model.Product, error)
	All() ([]model.Product, error)
	UpdateStrikeCode(int, string) (model.Product, error)
}
