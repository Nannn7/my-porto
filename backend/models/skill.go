package models

type Skill struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Level    string `json:"level"`
}
