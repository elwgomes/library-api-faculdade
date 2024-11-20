package handlers

import (
	"encoding/json"
	"library-api/core/models"
	db "library-api/infra/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if result := db.DB.Create(&book); result.Error != nil {
		http.Error(w, "Could not create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	if result := db.DB.Preload("Author").Preload("Category").Find(&books); result.Error != nil {
		http.Error(w, "Could not fetch books", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var book models.Book
	if result := db.DB.Preload("Author").Preload("Category").First(&book, id); result.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var book models.Book
	if result := db.DB.First(&book, id); result.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db.DB.Save(&book)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if result := db.DB.Delete(&models.Book{}, id); result.Error != nil {
		http.Error(w, "Could not delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
