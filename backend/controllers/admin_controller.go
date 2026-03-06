package controllers

import (
	"net/http"

	"myporto-backend/services"
	"myporto-backend/utils"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	service *services.AdminService
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAdminController(service *services.AdminService) *AdminController {
	return &AdminController{service: service}
}

func (ac *AdminController) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	if err := ac.service.Login(request.Username, request.Password); err != nil {
		utils.Error(c, http.StatusUnauthorized, "login failed", err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "login success", gin.H{"username": request.Username})
}
