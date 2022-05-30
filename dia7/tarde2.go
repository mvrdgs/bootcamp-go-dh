package dia7

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tarde2() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Maurício",
		})
	})

	router.Run()
}
