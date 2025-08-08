package main

import (
	"log"
	"net/http"

	"sakib.com/reelingit/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("Failed to initialize a logger $v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("public")))
	logInstance.Info("Serving the files!")

	// Start server
	const addr = ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed: %v", err)
	}
}
