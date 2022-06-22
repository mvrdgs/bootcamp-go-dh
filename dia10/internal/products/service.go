package products

import "github.com/mvrdgs/bootcamp-go-dh/dia10/internal/products/models"

type Service interface {
	GetAll() ([]models.Product, error)
	Store(name, tipo string, count int, price float64) (models.Product, error)
	Update(id int, name, tipo string, count int, price float64) (models.Product, error)
	UpdateName(id int, name string) (models.Product, error)
	Delete(id int) error
}
type service struct {
	repository Repository
}

func (s *service) GetAll() ([]models.Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(name, tipo string, count int, price float64) (models.Product, error) {
	product := models.Product{
		Name:  name,
		Type:  tipo,
		Count: count,
		Price: price,
	}

	product, err := s.repository.Store(product)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, name, tipo string, count int, price float64) (models.Product, error) {
	product := models.Product{
		ID:    id,
		Name:  name,
		Type:  tipo,
		Count: count,
		Price: price,
	}
	return s.repository.Update(product)
}

func (s *service) UpdateName(id int, name string) (models.Product, error) {
	product := models.Product{ID: id, Name: name}
	return s.repository.UpdateName(product)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
