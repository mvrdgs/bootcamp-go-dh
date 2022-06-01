package products

import "fmt"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

var ps []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, tipo string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, tipo string, count int, price float64) (Product, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) Store(id int, name, tipo string, count int, price float64) (Product, error) {
	p := Product{id, name, tipo, count, price}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, name, tipo string, count int, price float64) (Product, error) {
	p := Product{
		Name:  name,
		Type:  tipo,
		Count: count,
		Price: price,
	}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Produto %d não encontrado", id)
	}
	return p, nil
}

func NewRepository() Repository {
	return &repository{}
}
