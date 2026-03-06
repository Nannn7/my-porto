package main

import (
	"log"
	"os"
	"strings"

	"myporto-backend/config"
	"myporto-backend/controllers"
	"myporto-backend/middleware"
	"myporto-backend/models"
	"myporto-backend/repository"
	"myporto-backend/routes"
	"myporto-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	if err := config.DB.AutoMigrate(&models.Project{}); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	projectRepository := repository.NewProjectRepository()
	projectService := services.NewProjectService(projectRepository)
	projectController := controllers.NewProjectController(projectService)

	routes.RegisterRoutes(router, projectController)

	if err := runServer(router); err != nil {
		log.Fatal(err)
	}
}

func runServer(router *gin.Engine) error {
	if strings.EqualFold(os.Getenv("APP_ENV"), "development") {
		log.Println("Starting server in non-TLS mode for development")
		return router.Run(":8080")
	}

	certPath := getEnv("TLS_CERT_PATH", "backend/cert/cert.pem")
	keyPath := getEnv("TLS_KEY_PATH", "backend/cert/key.pem")

	log.Printf("Starting server with TLS cert=%s key=%s", certPath, keyPath)
	return router.RunTLS(":8080", certPath, keyPath)
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}
