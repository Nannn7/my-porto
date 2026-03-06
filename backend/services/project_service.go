package services

import (
	"myporto-backend/models"
	"myporto-backend/repository"
)

type ProjectService struct {
	repository *repository.ProjectRepository
}

func NewProjectService(repository *repository.ProjectRepository) *ProjectService {
	return &ProjectService{repository: repository}
}

func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
	return s.repository.FindAll()
}
