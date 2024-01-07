package postgres

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
	"testing"
)

var service ProductService

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		slog.Error(err.Error())
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", host, user, password, database, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't connect to storage to make tests...")
	}
	service = ProductService{DB: db}
}

func TestGetProducts(t *testing.T) {
	products, err := service.GetProducts()
	if err != nil {
		t.Error("should not have any error, err:", err.Error())
	}

	if len(*products) == 0 {
		t.Error("The database should not be empty.")
	}
}

func TestSearchProductsByName(t *testing.T) {
	products, err := service.GetProducts()
	if err != nil {
		t.Error("should not have any error, err:", err.Error())
	}

	product, err := service.SearchProductsByName((*products)[0].Name)
	if err != nil {
		t.Error("should not have any error, err:", err.Error())
	}

	if (*product)[0].Name != (*products)[0].Name {
		t.Errorf("names should be the same expected: %s got: %s", (*product)[0].Name, (*products)[0].Name)
	}
}

func TestValidateFormInput(t *testing.T) {
	v := govalidator.IsAlphanumeric("Cheese")
	if v != true {
		t.Errorf("incorrect result: expected %v got %v", false, v)
	}

	v = govalidator.IsAlphanumeric(";drop database northwind;")
	if v != false {
		t.Errorf("incorrect result: expected %v got %v", true, v)
		t.Error("Should be false.")
	}

}
