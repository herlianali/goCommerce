package main

import (
	"log"
	"net/http"
	"os"

	appHttp "github.com/herlianali/goCommerce/internal/http"
	"github.com/herlianali/goCommerce/internal/infrastructure/database"
	"github.com/herlianali/goCommerce/internal/infrastructure/repository"
	"github.com/herlianali/goCommerce/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db := database.ConnectPostgres()
	defer db.Close()

	userRepo := repository.NewPostgresUserRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")

	authUC := usecase.NewAuthUsecase(userRepo, jwtSecret)

	// jangan langsung panggil register/login di main nanti, buat endpoint API saja

	router := appHttp.NewRouter()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
