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
	if err := config.DB.AutoMigrate(
		&models.Project{},
		&models.Skill{},
		&models.ContactMessage{},
		&models.AdminUser{},
	); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	projectRepository := repository.NewProjectRepository()
	projectService := services.NewProjectService(projectRepository)
	projectController := controllers.NewProjectController(projectService)

	skillRepository := repository.NewSkillRepository()
	skillService := services.NewSkillService(skillRepository)
	skillController := controllers.NewSkillController(skillService)

	contactRepository := repository.NewContactRepository()
	contactService := services.NewContactService(contactRepository)
	contactController := controllers.NewContactController(contactService)

	adminRepository := repository.NewAdminRepository()
	adminService := services.NewAdminService(adminRepository)
	if err := adminService.EnsureDefaultAdmin(); err != nil {
		log.Fatal(err)
	}
	adminController := controllers.NewAdminController(adminService)

	routes.RegisterRoutes(
		router,
		projectController,
		skillController,
		contactController,
		adminController,
	)

	if err := router.RunTLS(":8080", "cert/cert.pem", "cert/key.pem"); err != nil {
		log.Fatal(err)
	}
}
