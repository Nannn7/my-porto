package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"myporto-backend/models"
	"myporto-backend/services"
	"myporto-backend/utils"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	service *services.ProjectService
}

type ProjectRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=120"`
	Description string `json:"description" binding:"required,min=10,max=2000"`
	TechStack   string `json:"tech_stack" binding:"omitempty,max=300"`
	ImageURL    string `json:"image_url" binding:"omitempty,url,max=600"`
	ProjectURL  string `json:"project_url" binding:"omitempty,url,max=600"`
}

func NewProjectController(service *services.ProjectService) *ProjectController {
	return &ProjectController{service: service}
}

func (pc *ProjectController) GetProjects(c *gin.Context) {
	projects, err := pc.service.GetAllProjects()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to fetch projects", err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "projects fetched", projects)
}

func (pc *ProjectController) CreateProject(c *gin.Context) {
	var request ProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if err := normalizeAndValidateProjectRequest(&request); err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	project := models.Project{
		Title:       request.Title,
		Description: request.Description,
		TechStack:   request.TechStack,
		ImageURL:    request.ImageURL,
		ProjectURL:  request.ProjectURL,
	}

	if err := pc.service.CreateProject(&project); err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to create project", err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "project created", project)
}

func (pc *ProjectController) UpdateProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid project id", err.Error())
		return
	}

	var request ProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if err := normalizeAndValidateProjectRequest(&request); err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	project := models.Project{
		ID:          uint(id),
		Title:       request.Title,
		Description: request.Description,
		TechStack:   request.TechStack,
		ImageURL:    request.ImageURL,
		ProjectURL:  request.ProjectURL,
	}

	if err := pc.service.UpdateProject(uint(id), &project); err != nil {
		utils.Error(c, http.StatusNotFound, "project not found", err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "project updated", project)
}

func (pc *ProjectController) DeleteProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid project id", err.Error())
		return
	}

	if err := pc.service.DeleteProject(uint(id)); err != nil {
		utils.Error(c, http.StatusNotFound, "project not found", err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "project deleted", nil)
}

func normalizeAndValidateProjectRequest(request *ProjectRequest) error {
	request.Title = strings.TrimSpace(request.Title)
	request.Description = strings.TrimSpace(request.Description)
	request.TechStack = strings.TrimSpace(request.TechStack)
	request.ImageURL = strings.TrimSpace(request.ImageURL)
	request.ProjectURL = strings.TrimSpace(request.ProjectURL)

	if request.Title == "" {
		return errors.New("title is required")
	}

	if request.Description == "" {
		return errors.New("description is required")
	}

	return nil
}
