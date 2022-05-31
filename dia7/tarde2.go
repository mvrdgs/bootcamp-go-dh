package dia7

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Products struct {
	Products []Product `json:"products"`
}

var products Products

type Product struct {
	Id            int     `json:"id,omitempty"`
	Nome          string  `json:"nome,omitempty" binding:"required"`
	Cor           string  `json:"cor,omitempty" binding:"required"`
	Preco         float64 `json:"preco,omitempty" binding:"required"`
	Estoque       int     `json:"estoque,omitempty" binding:"required"`
	Codigo        string  `json:"codigo,omitempty" binding:"required"`
	Publicacao    bool    `json:"publicacao,omitempty" binding:"required"`
	DataDeCriacao string  `json:"data_de_criacao,omitempty" binding:"required"`
}

func ReadFile() Products {
	file, _ := ioutil.ReadFile("./dia7/products.json")

	data := Products{}

	err := json.Unmarshal(file, &data.Products)
	if err != nil {
		fmt.Println(err)
	}

	return data
}

func getAll(c *gin.Context) {
	data := products.Products
	c.JSON(http.StatusOK, data)
}

func getById(c *gin.Context) {
	data := make(map[int]Product)

	for _, product := range products.Products {
		data[product.Id] = product
	}

	param, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "id precisa ser um número")
	}

	product, ok := data[param]

	if ok {
		c.JSON(http.StatusOK, product)
	} else {
		c.JSON(http.StatusNotFound, "produto não encontrado")
	}
}

func createProduct(c *gin.Context) {
	var product Product

	token := c.GetHeader("token")
	if token != "token super seguro" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "token inválido",
		})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product.Id = len(products.Products) + 1
	products.Products = append(products.Products, product)

	c.JSON(http.StatusOK, product)
}

func Tarde2() {
	router := gin.Default()
	products = ReadFile()

	group := router.Group("/products")

	{
		group.GET("/", getAll)
		group.GET("/:id", getById)
		group.POST("/", createProduct)
	}

	router.Run()
}
