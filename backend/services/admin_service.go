package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"myporto-backend/config"
	"myporto-backend/models"
	"myporto-backend/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminService struct {
	repository *repository.AdminRepository
}

func NewAdminService(repository *repository.AdminRepository) *AdminService {
	return &AdminService{repository: repository}
}

func (s *AdminService) Login(username, password string) (string, time.Time, string, error) {
	trimmedUsername := strings.TrimSpace(username)
	trimmedPassword := strings.TrimSpace(password)

	if trimmedUsername == "" || trimmedPassword == "" {
		return "", time.Time{}, "", errors.New("username and password are required")
	}

	admin, err := s.repository.FindByUsername(trimmedUsername)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", time.Time{}, "", errors.New("invalid username or password")
		}
		return "", time.Time{}, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(trimmedPassword)); err != nil {
		// Backward-compatible fallback for legacy plain-text admin password rows.
		if admin.Password != trimmedPassword {
			return "", time.Time{}, "", errors.New("invalid username or password")
		}

		passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(trimmedPassword), bcrypt.DefaultCost)
		if hashErr == nil {
			_ = s.repository.UpdatePassword(admin.ID, string(passwordHash))
		}
	}

	token, err := generateSessionToken()
	if err != nil {
		return "", time.Time{}, "", err
	}

	now := time.Now()
	expiresAt := now.Add(config.GetEnvDurationHours("ADMIN_SESSION_TTL_HOURS", 24))

	if err := s.repository.DeleteExpiredSessions(now); err != nil {
		return "", time.Time{}, "", err
	}

	session := models.AdminSession{
		AdminUserID: admin.ID,
		TokenHash:   hashToken(token),
		ExpiresAt:   expiresAt,
	}

	if err := s.repository.CreateSession(&session); err != nil {
		return "", time.Time{}, "", err
	}

	return token, expiresAt, admin.Username, nil
}

func (s *AdminService) ValidateSession(token string) (*models.AdminUser, error) {
	trimmedToken := strings.TrimSpace(token)
	if trimmedToken == "" {
		return nil, errors.New("missing access token")
	}

	session, err := s.repository.FindActiveSessionByHash(hashToken(trimmedToken), time.Now())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid or expired session")
		}
		return nil, err
	}

	admin, err := s.repository.FindByID(session.AdminUserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin user not found")
		}
		return nil, err
	}

	return admin, nil
}

func (s *AdminService) EnsureDefaultAdmin() error {
	username := config.GetEnv("DEFAULT_ADMIN_USERNAME", "admin")
	password := config.GetEnv("DEFAULT_ADMIN_PASSWORD", "admin123")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repository.FirstOrCreateDefaultAdmin(username, string(passwordHash))
}

func generateSessionToken() (string, error) {
	buffer := make([]byte, 32)
	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(buffer), nil
}

func hashToken(token string) string {
	digest := sha256.Sum256([]byte(token))
	return hex.EncodeToString(digest[:])
}
