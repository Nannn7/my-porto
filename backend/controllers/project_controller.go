package controllers

import (
	"net/http"

	"myporto-backend/services"
	"myporto-backend/utils"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	service *services.ProjectService
}

func NewProjectController(service *services.ProjectService) *ProjectController {
	return &ProjectController{service: service}
}

func (pc *ProjectController) GetProjects(c *gin.Context) {
	projects, err := pc.service.GetAllProjects()
	if err != nil {
		utils.JSON(c, http.StatusInternalServerError, "failed to fetch projects", nil)
		return
	}

	utils.JSON(c, http.StatusOK, "projects fetched", projects)
}
