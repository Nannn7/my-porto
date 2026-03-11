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
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	if config.GetEnvBool("APP_AUTO_MIGRATE", false) {
		if err := config.DB.AutoMigrate(
			&models.Project{},
			&models.Skill{},
			&models.ContactMessage{},
			&models.AdminUser{},
			&models.AdminSession{},
		); err != nil {
			log.Fatal(err)
		}

		log.Println("AutoMigrate completed")
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
	adminAuthMiddleware := middleware.AdminAuthMiddleware(adminService)

	routes.RegisterRoutes(
		router,
		projectController,
		skillController,
		contactController,
		adminController,
		adminAuthMiddleware,
	)

	if err := runServer(router); err != nil {
		log.Fatal(err)
	}
}

func runServer(router *gin.Engine) error {
	port := config.GetEnv("APP_PORT", "8080")
	listenAddr := ":" + port

	if !config.GetEnvBool("APP_TLS_ENABLED", false) {
		log.Printf("Starting server in HTTP mode on %s", listenAddr)
		return router.Run(listenAddr)
	}

	certPath := config.GetEnv("TLS_CERT_PATH", "cert/cert.pem")
	keyPath := config.GetEnv("TLS_KEY_PATH", "cert/key.pem")

	log.Printf("Starting server with TLS on %s cert=%s key=%s", listenAddr, certPath, keyPath)
	return router.RunTLS(listenAddr, certPath, keyPath)
}
