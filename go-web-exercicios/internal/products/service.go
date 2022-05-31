package products

type Service interface {
	GetAll() ([]Product, error)
	Store(estoque int, nome, cor, codigo, dataCriacao string, preco float64, publicacao bool) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Store(estoque int, nome, cor, codigo, dataCriacao string, preco float64, publicacao bool) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Store(lastID, estoque, nome, cor, codigo, dataCriacao, preco, publicacao)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
