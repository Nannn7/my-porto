package services

import (
	"myporto-backend/models"
	"myporto-backend/repository"
)

type SkillService struct {
	repository *repository.SkillRepository
}

func NewSkillService(repository *repository.SkillRepository) *SkillService {
	return &SkillService{repository: repository}
}

func (s *SkillService) GetAllSkills() ([]models.Skill, error) {
	return s.repository.FindAll()
}
