package dia7

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//type Products struct {
//	Products []Product `json:"products"`
//}

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

func ReadFile() []Product {
	file, _ := ioutil.ReadFile("./dia7/products.json")

	data := make([]Product, 0)

	fmt.Println(string(file))

	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println(err)
	}

	return data
}

func Tarde2() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Maurício",
		})
	})

	data := ReadFile()

	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
	})

	router.Run()
}
