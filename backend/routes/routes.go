package routes

import (
	"myporto-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	projectController *controllers.ProjectController,
	skillController *controllers.SkillController,
	contactController *controllers.ContactController,
	adminController *controllers.AdminController,
	adminAuthMiddleware gin.HandlerFunc,
) {
	api := router.Group("/api")
	{
		api.GET("/projects", projectController.GetProjects)
		api.GET("/skills", skillController.GetSkills)
		api.POST("/contact", contactController.CreateContact)

		admin := api.Group("/admin")
		{
			admin.POST("/login", adminController.Login)

			protected := admin.Group("")
			protected.Use(adminAuthMiddleware)
			{
				protected.GET("/projects", projectController.GetProjects)
				protected.POST("/projects", projectController.CreateProject)
				protected.PUT("/projects/:id", projectController.UpdateProject)
				protected.DELETE("/projects/:id", projectController.DeleteProject)
			}
		}
	}
}
