package main

import (
	db "library-api/infra/database"
	"log"
)

func main() {
	db.ConnectDB()
	log.Println("Aplicação iniciada...")
}
