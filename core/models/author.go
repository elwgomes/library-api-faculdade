package models

import "time"

type Author struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Biography string    `json:"biography"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
