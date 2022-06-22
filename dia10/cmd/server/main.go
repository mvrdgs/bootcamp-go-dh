package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/cmd/server/handler"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/internal/products"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/pkg/jsonStore"
	"github.com/mvrdgs/bootcamp-go-dh/dia10/pkg/web"
	"github.com/mvrdgs/bootcamp-go-dh/docs"
	"github.com/swaggo/files"
	ginSwager "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func respondWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, web.NewResponse(code, nil, message))
}

func TokenAuthMiddleware() gin.HandlerFunc {
	TOKEN := os.Getenv("TOKEN")

	if TOKEN == "" {
		log.Fatal("Please set TOKEN envoirment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "API token required")
		}

		if token != TOKEN {
			respondWithError(c, 401, "Invalid API token")
		}

		c.Next()
	}
}

// @title MELI Bootcamp API
// @version 1.0
// description This API Handle MELI Products
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos_y_condiciones

// @contact.name API Support
// @contact.url http://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func Main() {
	err := godotenv.Load("./dia10/cmd/server/.env")
	if err != nil {
		log.Fatalln(err.Error())
	}

	handler.TOKEN = os.Getenv("TOKEN")

	db := jsonStore.New(jsonStore.FileType, "./dia10/cmd/server/products_test.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwager.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	{
		pr.Use(TokenAuthMiddleware())
		pr.POST("/", p.Store())
		pr.GET("/", p.GetAll())
		pr.PUT("/:id", p.Update())
		pr.PATCH("/:id", p.UpdateName())
		pr.DELETE("/:id", p.Delete())
	}
	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
