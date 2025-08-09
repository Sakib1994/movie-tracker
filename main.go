package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"sakib.com/reelingit/data"
	"sakib.com/reelingit/handlers"
	"sakib.com/reelingit/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("Failed to initialize a logger %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}
	// Initialize logger
	logInstance := initializeLogger()

	// Database connection
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatalf("DATABASE_URL not set in environment")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	// delare movie repo
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to connect to MovieRepo: %v", err)
	}

	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("public")))
	logInstance.Info("Serving the files!")

	// Start server
	const addr = ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed: %v", err)
	}
}
