package models

import "time"

type Post struct {
	ID      uint   `gorm:"primaryKey"`
	Title   string `gorm:"not null"`
	Content string `gorm:"type:text"`
	UserID  uint   `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	Votes     int
}
