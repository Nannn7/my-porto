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

func (s *ProjectService) CreateProject(project *models.Project) error {
	return s.repository.Create(project)
}

func (s *ProjectService) UpdateProject(id uint, project *models.Project) error {
	if _, err := s.repository.FindByID(id); err != nil {
		return err
	}

	return s.repository.UpdateByID(id, project)
}

func (s *ProjectService) DeleteProject(id uint) error {
	if _, err := s.repository.FindByID(id); err != nil {
		return err
	}

	return s.repository.DeleteByID(id)
}
