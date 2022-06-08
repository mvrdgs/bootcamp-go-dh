package products

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Store struct {
	products []Product
	updated  bool
}

func (s *Store) Read(data interface{}) error {
	s.updated = true
	products, _ := json.Marshal(s.products)
	return json.Unmarshal(products, &data)
}

func (s *Store) Write(data interface{}) error {
	return nil
}

func TestRepository_GetAll(t *testing.T) {
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
	repo := NewRepository(&db)
	result, _ := repo.GetAll()

	assert.Equal(t, products, result)
}

func TestRepository_UpdateName(t *testing.T) {
	products := []Product{{
		ID:    1,
		Name:  "test",
		Type:  "test",
		Count: 1,
		Price: 1,
	}}
	db := Store{products: products, updated: false}
	repo := NewRepository(&db)
	expected := Product{
		ID:    1,
		Name:  "updated",
		Type:  "test",
		Count: 1,
		Price: 1,
	}

	result, _ := repo.UpdateName(1, "updated")

	assert.True(t, db.updated)
	assert.Equal(t, expected, result)
}
