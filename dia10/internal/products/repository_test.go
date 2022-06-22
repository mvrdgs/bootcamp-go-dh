package products

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/internal/products/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mockProduct = models.Product{
		ID:    1,
		Name:  "a",
		Type:  "a",
		Count: 1,
		Price: 1,
	}

	mockProduct2 = models.Product{
		ID:    2,
		Name:  "b",
		Type:  "b",
		Count: 2,
		Price: 2,
	}

	mockList = []models.Product{
		mockProduct,
		mockProduct2,
	}
)

func TestRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "type", "count", "price"}).
		AddRow(mockProduct.ID, mockProduct.Name, mockProduct.Type, mockProduct.Count, mockProduct.Price).
		AddRow(mockProduct2.ID, mockProduct2.Name, mockProduct2.Type, mockProduct2.Count, mockProduct2.Price)

	query := "SELECT id, name, type, count, price FROM products"

	mock.ExpectQuery(query).WillReturnRows(rows)

	productsRepo := NewRepository(db)
	result, err := productsRepo.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, result[0], mockProduct)
	assert.Equal(t, result[1], mockProduct2)
}
