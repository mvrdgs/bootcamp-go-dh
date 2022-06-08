package products

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Store struct {
	products []Product
}

func (s Store) Read(data interface{}) error {
	return nil
}

func (s Store) Write(data interface{}) error {
	return nil
}

func (r repository) GetAllTest(t *testing.T) {
	products := []Product{
		{
			ID:    1,
			Name:  "a",
			Type:  "a",
			Count: 1,
			Price: 1,
		},
		{
			ID:    2,
			Name:  "b",
			Type:  "b",
			Count: 2,
			Price: 2,
		},
	}
	db := Store{products: products}
	repo := NewRepository(db)
	result, _ := repo.GetAll()

	assert.Equal(t, result, products)
}
