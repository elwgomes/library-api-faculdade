package models

import "time"

type Comment struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	BookID uint `json:"book_id"`
	// Book      Book      `json:"book" gorm:"foreignKey:BookID"`
	UserID uint `json:"user_id"`
	// User      User      `json:"user" gorm:"foreignKey:UserID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
