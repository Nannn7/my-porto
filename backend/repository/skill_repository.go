package repository

import (
	"myporto-backend/config"
	"myporto-backend/models"
)

type SkillRepository struct{}

func NewSkillRepository() *SkillRepository {
	return &SkillRepository{}
}

func (r *SkillRepository) FindAll() ([]models.Skill, error) {
	var skills []models.Skill
	err := config.DB.Order("id DESC").Find(&skills).Error
	return skills, err
}
