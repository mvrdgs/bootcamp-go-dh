package products

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

func NewRepository() Repository {
	return &repository{}
}
