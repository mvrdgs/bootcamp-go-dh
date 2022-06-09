package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mvrdgs/bootcamp-go-dh/go-web-exercicios/cmd/server/handler"
	"github.com/mvrdgs/bootcamp-go-dh/go-web-exercicios/internal/products"
	"log"
)

func Main() {
	rep := products.NewRepository()
	service := products.NewService(rep)
	product := handler.NewProduct(service)

	router := gin.Default()
	productsGroup := router.Group("/produtos")
	productsGroup.GET("/", product.GetAll())
	productsGroup.POST("/", product.Store())
	err := router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
