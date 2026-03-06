package models

import "time"

type Project struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TechStack   string    `json:"tech_stack"`
	ImageURL    string    `json:"image_url"`
	ProjectURL  string    `json:"project_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
