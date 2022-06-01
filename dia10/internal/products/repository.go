package products

import (
	"fmt"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/pkg/store"
)

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

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	err := r.db.Read(&ps)
	if err != nil {
		return ps, err
	}
	return ps, nil
}

func (r *repository) Store(id int, name, tipo string, count int, price float64) (Product, error) {
	//p := Product{id, name, tipo, count, price}
	//ps = append(ps, p)
	//lastID = p.ID
	//return p, nil
	var ps []Product
	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}
	p := Product{id, name, tipo, count, price}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil
}

func (r *repository) Update(id int, name, tipo string, count int, price float64) (Product, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return Product{}, err
	}

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

	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
