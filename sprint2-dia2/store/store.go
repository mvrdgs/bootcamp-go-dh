package store

type Store struct {
	Products []Product
}

type Product struct {
	Name     string
	Quantity int
}

type Repository interface {
	GetAll() []Product
}

func (s *Store) GetAll() []Product {
	return s.Products
}
