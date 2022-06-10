package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/internal/products"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type Store struct {
	products []Product
	updated  bool
	err      error
}

func (s *Store) Read(data interface{}) error {
	s.updated = true
	prods, _ := json.Marshal(s.products)
	if s.err != nil {
		return s.err
	}

	return json.Unmarshal(prods, &data)
}

func (s *Store) Write(data interface{}) error {
	prod, _ := json.Marshal(data)
	err := json.Unmarshal(prod, &s.products)
	if err != nil {
		return err
	}

	return nil
}

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	db := Store{
		products: []Product{},
		updated:  false,
		err:      nil,
	}
	repo := products.NewRepository(&db)
	service := products.NewService(repo)
	p := NewProduct(service)
	r := gin.Default()

	r.GET("/products", p.GetAll())
	r.POST("/products", p.Store())

	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestProduct_GetAll(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestProduct_Store(t *testing.T) {
	r := createServer()
	body := `{"name": "test", "type": "ola", "count": 10, "price": 10.99}`
	req, rr := createRequestTest(http.MethodPost, "/products", body)

	r.ServeHTTP(rr, req)
	
	assert.Equal(t, 201, rr.Code)
}
