package handlers

import (
	storage "construccion_demo/internal/storage/postgres"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type service struct {
	Port           string
	DB             *gorm.DB
	ProductService *storage.ProductService
}

func New(port, host, user, password, database, dbport string) *service {
	db := createDatabase(host, user, password, database, dbport)
	return &service{
		Port:           port,
		DB:             db,
		ProductService: createProductService(db),
	}
}

func createDatabase(host, user, password, database, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, database, port)
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic("Can't connect to storage...")
	}
	slog.Info("Database connection started...")
	return db
}

func createProductService(db *gorm.DB) *storage.ProductService {
	return &storage.ProductService{
		DB: db,
	}
}
