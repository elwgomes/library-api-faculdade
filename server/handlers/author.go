package handlers

import (
	"encoding/json"
	"library-api/core/models"
	db "library-api/infra/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if result := db.DB.Create(&author); result.Error != nil {
		http.Error(w, "Could not create author", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	if result := db.DB.Find(&authors); result.Error != nil {
		http.Error(w, "Could not fetch authors", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(authors)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var author models.Author
	if result := db.DB.First(&author, id); result.Error != nil {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(author)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var author models.Author
	if result := db.DB.First(&author, id); result.Error != nil {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db.DB.Save(&author)
	json.NewEncoder(w).Encode(author)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if result := db.DB.Delete(&models.Author{}, id); result.Error != nil {
		http.Error(w, "Could not delete author", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
