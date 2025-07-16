package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"unique;not null" json:"email"`
	Username     string    `gorm:"unique;not null" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"default:'patient'" json:"role"`
	Verified     bool      `gorm:"default:false" json:"verified"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}