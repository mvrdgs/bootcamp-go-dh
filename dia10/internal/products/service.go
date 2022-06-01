package products

import "log"

type Service interface {
	GetAll() ([]Product, error)
	Store(name, tipo string, count int, price float64) (Product, error)
	Update(id int, name, tipo string, count int, price float64) (Product, error)
}
type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(name, tipo string, count int, price float64) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Store(lastID, name, tipo, count, price)
	log.Println(product)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, name, tipo string, count int, price float64) (Product, error) {
	return s.repository.Update(id, name, tipo, count, price)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
