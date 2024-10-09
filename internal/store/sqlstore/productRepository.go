package sqlstore

import "obstore/internal/model"

type ProductRepository struct {
	store *Store
}

func (r *ProductRepository) Create(u model.Product) (model.Product, error) {
	return u, r.store.db.Create(&u).Error
}

func (r *ProductRepository) GetAll() (products []model.Product, err error) {
	return products, r.store.db.Table("products").Select("*").Scan(&products).Error
}

func (r *ProductRepository) GetByArticle(articul string) (product model.Product, err error) {
	return product, r.store.db.Table("products").
		Select("*").Where("article=?", articul).
		Scan(&product).Error
}

func (r *ProductRepository) GetByStrikeCode(strikecode string) (products []model.Product, err error) {
	return products, r.store.db.Table("products").
		Select("*").Where("strikecode=?", strikecode).
		Scan(&products).Error
}

func (r *ProductRepository) GetByName(nameArtic string) (products []model.Product, err error) {
	return products, r.store.db.Table("products").
		Select("*").Where("nameArtic LIKE '%?%'", nameArtic).
		Scan(&products).Error
}
