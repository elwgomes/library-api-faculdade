package models

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BookID    uint      `json:"book_id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
