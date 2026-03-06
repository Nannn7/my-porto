package repository

import (
	"errors"
	"myporto-backend/config"
	"myporto-backend/models"

	"gorm.io/gorm"
)

type AdminRepository struct{}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

func (r *AdminRepository) FindByUsername(username string) (*models.AdminUser, error) {
	var admin models.AdminUser
	if err := config.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) FirstOrCreateDefaultAdmin() error {
	var count int64
	if err := config.DB.Model(&models.AdminUser{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	defaultAdmin := models.AdminUser{Username: "admin", Password: "admin123"}
	return config.DB.Create(&defaultAdmin).Error
}
