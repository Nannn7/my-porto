package services

import (
	"errors"
	"myporto-backend/repository"
)

type AdminService struct {
	repository *repository.AdminRepository
}

func NewAdminService(repository *repository.AdminRepository) *AdminService {
	return &AdminService{repository: repository}
}

func (s *AdminService) Login(username, password string) error {
	admin, err := s.repository.FindByUsername(username)
	if err != nil {
		return err
	}

	if admin.Password != password {
		return errors.New("invalid username or password")
	}

	return nil
}

func (s *AdminService) EnsureDefaultAdmin() error {
	return s.repository.FirstOrCreateDefaultAdmin()
}
