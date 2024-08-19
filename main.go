package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "OK")
	})
	router.Run(":7777")
}
