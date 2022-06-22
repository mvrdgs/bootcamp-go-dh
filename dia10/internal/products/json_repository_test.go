package products

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Store struct {
	products []Product
	updated  bool
	err      error
}

func (s *Store) Read(data interface{}) error {
	s.updated = true
	products, _ := json.Marshal(s.products)
	if s.err != nil {
		return s.err
	}

	return json.Unmarshal(products, &data)
}

func (s *Store) Write(data interface{}) error {
	prod, _ := json.Marshal(data)
	err := json.Unmarshal(prod, &s.products)
	if err != nil {
		return err
	}

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

func TestService_GetAll(t *testing.T) {
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
	service := NewService(repo)

	result, _ := service.GetAll()

	assert.Equal(t, products, result)
}

func TestService_GetAllError(t *testing.T) {
	mockError := errors.New("could not get products")
	db := Store{
		products: []Product{},
		updated:  false,
		err:      mockError,
	}
	repo := NewRepository(&db)
	service := NewService(repo)

	_, err := service.GetAll()

	assert.Error(t, err)
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

func TestService_UpdateName(t *testing.T) {
	products := []Product{{
		ID:    1,
		Name:  "test",
		Type:  "test",
		Count: 1,
		Price: 1,
	}}
	db := Store{products: products, updated: false}
	repo := NewRepository(&db)
	service := NewService(repo)

	expected := Product{
		ID:    1,
		Name:  "updated",
		Type:  "test",
		Count: 1,
		Price: 1,
	}

	result, _ := service.UpdateName(1, "updated")
	assert.Equal(t, expected, result)
}

func TestService_UpdateNameError(t *testing.T) {
	products := []Product{{
		ID:    1,
		Name:  "test",
		Type:  "test",
		Count: 1,
		Price: 1,
	}}
	db := Store{products: products, updated: false}
	repo := NewRepository(&db)
	service := NewService(repo)

	_, err := service.UpdateName(2, "updated")

	assert.Error(t, err)
}

func TestService_Store(t *testing.T) {
	var products []Product
	expected := Product{
		ID:    1,
		Name:  "a",
		Type:  "a",
		Count: 1,
		Price: 1,
	}

	db := Store{products: products}
	repo := NewRepository(&db)
	service := NewService(repo)

	result, _ := service.Store("a", "a", 1, 1)

	assert.Equal(t, expected, result)
}

func TestService_Delete(t *testing.T) {
	products := []Product{{
		ID:    1,
		Name:  "a",
		Type:  "a",
		Count: 1,
		Price: 1,
	}}

	db := Store{products: products}
	repo := NewRepository(&db)
	service := NewService(repo)

	err := service.Delete(1)
	result, _ := service.GetAll()

	assert.Nil(t, err)
	assert.Len(t, result, 0)
}
