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
	// Initialize account repository
	accountRepo, err := data.NewAccountRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize account repository: %v", err)
	}

	// Http multiplexer
	mux := http.NewServeMux()

	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)
	accountHandler := handlers.NewAccountHandler(accountRepo, logInstance)

	mux.HandleFunc("GET /api/movies/top", movieHandler.GetTopMovies)
	mux.HandleFunc("GET /api/movies/random", movieHandler.GetRandomMovies)
	mux.HandleFunc("GET /api/movies/search", movieHandler.SearchMovies)
	mux.HandleFunc("GET /api/movies/{id}", movieHandler.GetMovie)
	mux.HandleFunc("GET /api/genres", movieHandler.GetGenres)
	mux.HandleFunc("POST /api/account/register/", accountHandler.Register)
	mux.HandleFunc("POST /api/account/authenticate/", accountHandler.Authenticate)
	mux.HandleFunc("GET /api/account/favorites/", accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.GetFavorites)))
	mux.HandleFunc("GET /api/account/watchlist/", accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.GetWatchlist)))
	mux.HandleFunc("POST /api/account/save-to-collection/", accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.SaveToCollection)))

	catchAllHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}
	/*
		// SSR for the movie details
		http.HandleFunc("/movies/", func(w http.ResponseWriter, r *http.Request) {
			if strings.Count(r.URL.Path, "/") == 2 && strings.HasPrefix(r.URL.Path, "/movies/") {
				handlers.SSRMovieDetailsHandler(movieRepo, logInstance)(w, r)
			} else {
				catchAllClientRoutesHandler(w, r)
			}
		})*/

	// Catch All
	mux.HandleFunc("GET /movies", catchAllHandler)
	mux.HandleFunc("GET /movies/", catchAllHandler)
	mux.HandleFunc("GET /account/", catchAllHandler)

	// Serve static files
	mux.Handle("GET /", http.FileServer(http.Dir("public")))
	logInstance.Info("Serving the files!")

	// Start server
	const addr = ":8080"
	if err := http.ListenAndServe(addr, mux); err != nil {
		logInstance.Error("Server failed: %v", err)
	}
}
