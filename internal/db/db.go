package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/xadrijo/grpc-vidflex/internal/product"
)

type Storage struct {
	db *sqlx.DB
}

func New() (Storage, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbName,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Storage{}, nil
	}

	return Storage{
		db: db,
	}, nil
}

func (s Storage) Get(id int32) (product.Product, error) {
	return product.Product{}, nil
}

func (s Storage) Save(p product.Product) (product.Product, error) {
	return product.Product{}, nil
}
