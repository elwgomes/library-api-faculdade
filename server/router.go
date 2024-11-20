package server

import (
	"library-api/server/handlers"
	// "library-api/server/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Middleware de autenticação
	// router.Use(middleware.AuthMiddleware)

	// Rotas para usuários
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	// Rotas para livros
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// Rotas para autores
	router.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors", handlers.GetAllAuthors).Methods("GET")

	// Rotas para categorias
	router.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories", handlers.GetAllCategories).Methods("GET")

	// Rotas para comentários
	router.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
	router.HandleFunc("/books/{id}/comments", handlers.GetCommentsByBook).Methods("GET")

	return router
}
