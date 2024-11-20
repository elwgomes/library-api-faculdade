package handlers

import (
	"encoding/json"
	"library-api/core/models"
	db "library-api/infra/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if result := db.DB.Create(&comment); result.Error != nil {
		http.Error(w, "Could not create comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func GetCommentsByBook(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var comments []models.Comment
	if result := db.DB.Where("book_id = ?", bookID).Preload("User").Find(&comments); result.Error != nil {
		http.Error(w, "Could not fetch comments", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}
