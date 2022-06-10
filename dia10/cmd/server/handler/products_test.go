package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/internal/products"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/pkg/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	db := store.New(store.FileType, "./products_test.json")
	repo := products.NewRepository(db)
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
