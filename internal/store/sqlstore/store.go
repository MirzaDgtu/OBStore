package sqlstore

import (
	"obstore/internal/store"

	"gorm.io/gorm"
)

type Store struct {
	db                        *gorm.DB
	userRepository            *UserRepository
	teamRepository            *TeamRepository
	orderRepository           *OrderRepository
	productRepository         *ProductRepository
	teamCompositionRepository *TeamCompositionRepository
	assemblyOrderRepository   *AssemblyOrderRepository
	warehouseRepository       *WarehouseRepository
}

func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Team() store.TeamRepository {
	if s.teamRepository != nil {
		return s.teamRepository
	}

	s.teamRepository = &TeamRepository{
		store: s,
	}

	return s.teamRepository
}

func (s *Store) Order() store.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		store: s,
	}

	return s.orderRepository
}

func (s *Store) Product() store.ProductRepository {
	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &ProductRepository{
		store: s,
	}

	return s.productRepository
}

func (s *Store) TeamComposition() store.TeamCompositionRepository {
	if s.teamCompositionRepository != nil {
		return s.teamCompositionRepository
	}

	s.teamCompositionRepository = &TeamCompositionRepository{
		store: s,
	}

	return s.teamCompositionRepository
}

func (s *Store) AssemblyOrder() store.AssemblyOrderRepository {
	if s.assemblyOrderRepository != nil {
		return s.assemblyOrderRepository
	}

	s.assemblyOrderRepository = &AssemblyOrderRepository{
		store: s,
	}

	return s.assemblyOrderRepository
}

func (s *Store) Warehouse() store.WarehouseRepository {
	if s.warehouseRepository != nil {
		return s.warehouseRepository
	}

	s.warehouseRepository = &WarehouseRepository{
		store: s,
	}

	return s.warehouseRepository
}
