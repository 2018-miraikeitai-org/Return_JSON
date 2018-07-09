package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // gin.H は map[string]interface{} へのショートカット
			"Result": "Hello gin",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
