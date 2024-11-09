package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AuthorID    uint      `json:"author_id"`
	Author      Author    `json:"author" gorm:"foreignKey:AuthorID"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	ISBN        string    `json:"isbn" gorm:"unique"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
