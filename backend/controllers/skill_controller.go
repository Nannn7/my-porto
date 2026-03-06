package controllers

import (
	"net/http"

	"myporto-backend/services"
	"myporto-backend/utils"

	"github.com/gin-gonic/gin"
)

type SkillController struct {
	service *services.SkillService
}

func NewSkillController(service *services.SkillService) *SkillController {
	return &SkillController{service: service}
}

func (sc *SkillController) GetSkills(c *gin.Context) {
	skills, err := sc.service.GetAllSkills()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to fetch skills", err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "skills fetched", skills)
}
