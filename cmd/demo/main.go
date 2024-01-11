package main

import (
	"construccion_demo/internal/handlers"
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// load secrets read from os
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
	}

	// initiating environment variables
	port := os.Getenv("PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	// app configuration
	mux := http.NewServeMux()
	appConfig := handlers.New(port, host, user, password, database, dbPort)

	// routes / endpoints
	mux.HandleFunc("/", appConfig.Home)

	fmt.Print("Demo server running!\nYour server is running at: http://127.0.0.1:" + appConfig.Port)
	err = http.ListenAndServe(":"+appConfig.Port, mux)
	if err != nil {
		slog.Error(err.Error())
		panic("The server can't start...")
		return
	}
}
