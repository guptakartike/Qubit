package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/guptakartike/qubit/internal/auth"
	"github.com/guptakartike/qubit/internal/auth/handler"
	"github.com/guptakartike/qubit/internal/auth/repository"
	"github.com/guptakartike/qubit/internal/auth/service"
	"github.com/guptakartike/qubit/internal/database"
	"github.com/guptakartike/qubit/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic("invalid port")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	pgxPool, err := database.NewPool(databaseURL)
	if err != nil {
		panic("pool connection failed: " + err.Error())
	}
	defer pgxPool.Close()

	fmt.Println("Database connection successful")

	repo := repository.NewPostgresRepository(pgxPool)

	registrationService := service.NewRegistrationService(
		repo,
		auth.Hasher{},
	)

	registrationHandler := handler.NewRegistrationHandler(
		registrationService,
	)

	srv := server.New(
		portInt,
		registrationHandler,
	)

	fmt.Println("Qubit API running on http://localhost:" + port)

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
