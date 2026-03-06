package controllers

import (
	"net/http"

	"myporto-backend/models"
	"myporto-backend/services"
	"myporto-backend/utils"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
	service *services.ContactService
}

type ContactRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Message string `json:"message" binding:"required"`
}

func NewContactController(service *services.ContactService) *ContactController {
	return &ContactController{service: service}
}

func (cc *ContactController) CreateContact(c *gin.Context) {
	var request ContactRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	message := models.ContactMessage{
		Name:    request.Name,
		Email:   request.Email,
		Message: request.Message,
	}

	if err := cc.service.CreateContactMessage(&message); err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to submit contact message", err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "contact message submitted", message)
}
