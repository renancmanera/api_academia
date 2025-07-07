package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/renancmanera/api_academia/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Carrega as variáveis do .env (caso ainda não tenha sido carregado)
	_ = godotenv.Load()

	// Tenta usar DATABASE_URL primeiro
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Se não existir, monta manualmente a string de conexão
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			os.Getenv("SSL_MODE"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	DB = db
	log.Println("Conexão com o banco de dados estabelecida com sucesso!")

	// Migração automática da tabela usuarios
	err = db.AutoMigrate(&domain.Usuario{})
	if err != nil {
		log.Fatalf("Erro ao migrar tabela Usuario: %v", err)
	}
}
