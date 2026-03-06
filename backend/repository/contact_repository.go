package repository

import (
	"myporto-backend/config"
	"myporto-backend/models"
)

type ContactRepository struct{}

func NewContactRepository() *ContactRepository {
	return &ContactRepository{}
}

func (r *ContactRepository) Create(message *models.ContactMessage) error {
	return config.DB.Create(message).Error
}
