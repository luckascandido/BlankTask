package common

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	DbPort           string `env:"DB_PORT" envDefault:"5432"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
}

func Mypg() (*gorm.DB, error) {
	// carrega as variaveis do banco de dados
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("Erro ao carregar variáveis de ambiente:", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB, cfg.DbPort)
	// Faz a conexão com o banco
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o banco: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter banco de dados subjacente: %w", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Println("Error closing database connection:", err)
		}
	}()
	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao pingar o banco de dados: %w", err)
	}

	fmt.Println("Conectado ao banco de dados!")
	return db, nil
}
