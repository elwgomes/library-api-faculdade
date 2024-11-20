package main

import (
	db "library-api/infra/database"
	"library-api/server"
	"log"
	"net/http"
)

func main() {
	db.ConnectDB()
	router := server.SetupRouter()

	log.Println("Aplicação iniciada na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
