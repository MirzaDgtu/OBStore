package store

type Store interface {
	User() UserRepository
	Team() TeamRepository
	Order() OrderRepository
	Product() ProductRepository
	TeamComposition() TeamCompositionRepository
}
