package config

import (
	"library-api/core/models"
	"log"
	"time"
)

func SeedDatabase() {
	categories := []models.Category{
		{Name: "Fantasia", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Ficção Científica", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Terror", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, category := range categories {
		if err := DB.FirstOrCreate(&category, models.Category{Name: category.Name}).Error; err != nil {
			log.Printf("Erro ao semear categoria %s: %v\n", category.Name, err)
		}
	}

	authors := []models.Author{
		{FirstName: "J.K.", LastName: "Rowling", Biography: "Autora de Harry Potter", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{FirstName: "Isaac", LastName: "Asimov", Biography: "Pai da ficção científica moderna", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, author := range authors {
		if err := DB.FirstOrCreate(&author, models.Author{FirstName: author.FirstName, LastName: author.LastName}).Error; err != nil {
			log.Printf("Erro ao semear autor %s %s: %v\n", author.FirstName, author.LastName, err)
		}
	}

	books := []models.Book{
		{
			Title:       "Harry Potter e a Pedra Filosofal",
			Description: "Primeiro livro da saga Harry Potter.",
			AuthorID:    1,
			CategoryID:  1,
			ISBN:        "9788532530783",
			PublishedAt: time.Date(1997, 6, 26, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Eu, Robô",
			Description: "Clássico da ficção científica de Isaac Asimov.",
			AuthorID:    2,
			CategoryID:  2,
			ISBN:        "9788576572007",
			PublishedAt: time.Date(1950, 12, 2, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, book := range books {
		if err := DB.FirstOrCreate(&book, models.Book{Title: book.Title}).Error; err != nil {
			log.Printf("Erro ao semear livro %s: %v\n", book.Title, err)
		}
	}

	user := models.User{
		Username:  "admin",
		Password:  "admin123",
		Role:      "librarian",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := DB.FirstOrCreate(&user, models.User{Username: user.Username}).Error; err != nil {
		log.Printf("Erro ao semear usuário %s: %v\n", user.Username, err)
	}
}
