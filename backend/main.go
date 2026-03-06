package main

import (
	"myporto-backend/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	r := gin.Default()

	r.GET("/api/projects", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API running",
		})
	})

	r.RunTLS(":8080", "cert.pem", "key.pem")
}
