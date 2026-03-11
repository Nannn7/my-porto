package models

import "time"

type AdminSession struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	AdminUserID uint      `json:"admin_user_id" gorm:"index;not null"`
	TokenHash   string    `json:"-" gorm:"uniqueIndex;size:128;not null"`
	ExpiresAt   time.Time `json:"expires_at" gorm:"index;not null"`
	CreatedAt   time.Time `json:"created_at"`
}
