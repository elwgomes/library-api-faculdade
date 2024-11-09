package config

import (
	"fmt"
	"library-api/core/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db eh a variável global q mantem a conexão com o banco.
var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"mysql", "mysql", "localhost", "3306", "library_db")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	// aqui nois chama as migration
	err = migrateModels()
	if err != nil {
		log.Fatalf("Erro ao realizar a migração dos modelos: %v", err)
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso.")
}

// aqui nois faz as migration
func migrateModels() error {
	// as migration sao criadas com base nos models, pra facilitar nossa vida viados
	err := DB.AutoMigrate(
		&models.User{},
		&models.Author{},
		&models.Book{},
		&models.Category{},
		&models.Comment{},
	)
	return err
}
