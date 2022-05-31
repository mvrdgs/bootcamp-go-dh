package products

type Product struct {
	ID            int     `json:"id,omitempty"`
	Nome          string  `json:"nome,omitempty" binding:"required"`
	Cor           string  `json:"cor,omitempty" binding:"required"`
	Preco         float64 `json:"preco,omitempty" binding:"required"`
	Estoque       int     `json:"estoque,omitempty" binding:"required"`
	Codigo        string  `json:"codigo,omitempty" binding:"required"`
	Publicacao    bool    `json:"publicacao,omitempty" binding:"required"`
	DataDeCriacao string  `json:"data_de_criacao,omitempty" binding:"required"`
}

var products []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id, estoque int, nome, cor, codigo, dataCriacao string, preco float64, publicacao bool) (Product, error)
	LastID() (int, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) Store(id, estoque int, nome, cor, codigo, dataCriacao string, preco float64, publicacao bool) (Product, error) {
	product := Product{id, nome, cor, preco, estoque, codigo, publicacao, dataCriacao}
	products = append(products, product)
	lastID = product.ID
	return product, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func NewRepository() Repository {
	return &repository{}
}
