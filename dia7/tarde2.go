package dia7

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id            int     `json:"id,omitempty"`
	Nome          string  `json:"nome,omitempty"`
	Cor           string  `json:"cor,omitempty"`
	Preco         float64 `json:"preco,omitempty"`
	Estoque       int     `json:"estoque,omitempty"`
	Codigo        string  `json:"codigo,omitempty"`
	Publicacao    bool    `json:"publicacao,omitempty"`
	DataDeCriacao string  `json:"data_de_criacao,omitempty"`
}

func ReadFile(c *gin.Context) Products {
	file, _ := ioutil.ReadFile("./dia7/products.json")

	data := Products{}

	err := json.Unmarshal(file, &data.Products)
	if err != nil {
		fmt.Println(err)
	}

	query := c.Request.URL.Query()
	log.Println(query)

	if len(query) == 0 {
		return data
	}

	filteredData := Products{make([]Product, 0)}

	// falta fazer filtro
	filteredData = data

	return filteredData
}

func getAll(c *gin.Context) {
	data := ReadFile(c)
	c.JSON(http.StatusOK, data)
}

func getById(c *gin.Context) {
	data := ReadFile(c)
	products := make(map[int]Product)

	for _, product := range data.Products {
		products[product.Id] = product
	}

	param, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "id precisa ser um número")
	}

	product, ok := products[param]

	if ok {
		c.JSON(http.StatusOK, product)
	} else {
		c.JSON(http.StatusNotFound, "produto não encontrado")
	}
}

func Tarde2() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Maurício",
		})
	})

	group := router.Group("/products")

	{
		group.GET("/", getAll)
		group.GET("/:id", getById)
	}

	router.Run()
}
