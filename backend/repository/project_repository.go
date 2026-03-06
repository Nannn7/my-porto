package repository

import (
	"myporto-backend/config"
	"myporto-backend/models"
)

type ProjectRepository struct{}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{}
}

func (r *ProjectRepository) FindAll() ([]models.Project, error) {
	var projects []models.Project
	err := config.DB.Order("id DESC").Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return config.DB.Create(project).Error
}

func (r *ProjectRepository) UpdateByID(id uint, updates *models.Project) error {
	return config.DB.Model(&models.Project{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ProjectRepository) DeleteByID(id uint) error {
	return config.DB.Delete(&models.Project{}, id).Error
}

func (r *ProjectRepository) FindByID(id uint) (*models.Project, error) {
	var project models.Project
	if err := config.DB.First(&project, id).Error; err != nil {
		return nil, err
	}

	return &project, nil
}
