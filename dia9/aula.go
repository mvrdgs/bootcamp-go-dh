package dia9

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type request struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Tipo       string  `json:"tipo"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

var products []request
var lastID int

func Salvar(c *gin.Context) {
	var req request

	token := c.GetHeader("token")
	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "token inválido",
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	lastID++
	req.ID = lastID

	products = append(products, req)
	c.JSON(http.StatusOK, req)
}

func Aula() {
	router := gin.Default()

	postRouter := router.Group("/produtos")
	postRouter.POST("/", Salvar)

	router.Run()
}
