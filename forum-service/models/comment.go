package models

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	PostID    uint   `gorm:"not null;index"` // Links to the Post this comment belongs to
	UserID    uint   `gorm:"not null"`       // ID of the user who made the comment
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
}
