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
	err := config.DB.Find(&projects).Error
	return projects, err
}
