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
) {
	api := router.Group("/api")
	{
		api.GET("/projects", projectController.GetProjects)
		api.GET("/skills", skillController.GetSkills)
		api.POST("/contact", contactController.CreateContact)

		admin := api.Group("/admin")
		{
			admin.POST("/login", adminController.Login)
			admin.POST("/projects", projectController.CreateProject)
			admin.PUT("/projects/:id", projectController.UpdateProject)
			admin.DELETE("/projects/:id", projectController.DeleteProject)
		}
	}
}
