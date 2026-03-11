package repository

import (
	"errors"
	"time"

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

func (r *AdminRepository) FindByID(id uint) (*models.AdminUser, error) {
	var admin models.AdminUser
	if err := config.DB.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) UpdatePassword(id uint, passwordHash string) error {
	return config.DB.Model(&models.AdminUser{}).Where("id = ?", id).Update("password", passwordHash).Error
}

func (r *AdminRepository) FirstOrCreateDefaultAdmin(username, passwordHash string) error {
	var count int64
	if err := config.DB.Model(&models.AdminUser{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	defaultAdmin := models.AdminUser{Username: username, Password: passwordHash}
	return config.DB.Create(&defaultAdmin).Error
}

func (r *AdminRepository) CreateSession(session *models.AdminSession) error {
	return config.DB.Create(session).Error
}

func (r *AdminRepository) FindActiveSessionByHash(tokenHash string, now time.Time) (*models.AdminSession, error) {
	var session models.AdminSession
	if err := config.DB.
		Where("token_hash = ? AND expires_at > ?", tokenHash, now).
		First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *AdminRepository) DeleteExpiredSessions(now time.Time) error {
	return config.DB.Where("expires_at <= ?", now).Delete(&models.AdminSession{}).Error
}
