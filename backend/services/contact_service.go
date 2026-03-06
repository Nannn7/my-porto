package services

import (
	"myporto-backend/models"
	"myporto-backend/repository"
)

type ContactService struct {
	repository *repository.ContactRepository
}

func NewContactService(repository *repository.ContactRepository) *ContactService {
	return &ContactService{repository: repository}
}

func (s *ContactService) CreateContactMessage(message *models.ContactMessage) error {
	return s.repository.Create(message)
}
