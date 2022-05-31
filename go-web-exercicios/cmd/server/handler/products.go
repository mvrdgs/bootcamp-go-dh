package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mvrdgs/bootcamp-go-dh/go-web-exercicios/internal/products"
	"net/http"
)

type Request struct {
	Nome          string  `json:"nome" binding:"required"`
	Cor           string  `json:"cor" binding:"required"`
	Preco         float64 `json:"preco" binding:"required"`
	Estoque       int     `json:"estoque" binding:"required"`
	Codigo        string  `json:"codigo" binding:"required"`
	Publicacao    bool    `json:"publicacao" binding:"required"`
	DataDeCriacao string  `json:"data_de_criacao" binding:"required"`
}

type Product struct {
	service products.Service
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "token super secreto" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, products)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "token super secreto" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := p.service.Store(req.Estoque, req.Nome, req.Cor, req.Cor, req.DataDeCriacao, req.Preco, req.Publicacao)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, product)
	}
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}
