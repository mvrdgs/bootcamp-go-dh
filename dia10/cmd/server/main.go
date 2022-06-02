package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/cmd/server/handler"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/internal/products"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/pkg/store"
	"log"
	"os"
)

func Main() {
	err := godotenv.Load("./dia10/cmd/server/.env")
	if err != nil {
		log.Fatalln(err.Error())
	}

	handler.TOKEN = os.Getenv("TOKEN")

	db := store.New(store.FileType, "./dia10/cmd/server/products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
