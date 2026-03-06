package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/api/projects", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"projects": []string{
				"Portfolio Website",
				"Todo App",
				"Ecommerce Demo",
			},
		})
	})

	r.Run(":8080")
}