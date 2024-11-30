package main

import (
	"log"
	"mini-hibp/internal/database"
	"mini-hibp/internal/handler"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// db init
	db, err := database.InitDatabase(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	router := http.NewServeMux()

	router.HandleFunc("/api/v1/hibp", handler.CheckHandler(db))

	// CORS setup
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Ganti dengan URL frontend kamu
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap router with CORS middleware
	handlerWithCORS := corsHandler.Handler(router)

	log.Printf("Server running on port %s", port)

	err = http.ListenAndServe(":"+port, handlerWithCORS)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
