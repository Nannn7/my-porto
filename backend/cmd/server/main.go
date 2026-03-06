package main

import (
	"log"

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
	config.ConnectDB()
	if err := config.DB.AutoMigrate(&models.Project{}); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	projectRepository := repository.NewProjectRepository()
	projectService := services.NewProjectService(projectRepository)
	projectController := controllers.NewProjectController(projectService)

	routes.RegisterRoutes(router, projectController)

	if err := router.RunTLS(":8080", "cert/cert.pem", "cert/key.pem"); err != nil {
		log.Fatal(err)
	}
}
