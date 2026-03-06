package routes

import (
	"myporto-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, projectController *controllers.ProjectController) {
	api := router.Group("/api")
	{
		api.GET("/projects", projectController.GetProjects)
	}
}
