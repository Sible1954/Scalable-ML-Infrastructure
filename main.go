
package main

import (
	"fmt"
	"log"
	"net/http"
	"scalable-ml-infrastructure/api"
	"scalable-ml-infrastructure/config"
	"scalable-ml-infrastructure/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	fmt.Printf("Starting Scalable ML Infrastructure Server on port %s...
", cfg.ServerPort)

	router := http.NewServeMux()

	api.RegisterHandlers(router)

	loggingHandler := middleware.LoggingMiddleware(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: loggingHandler,
	}

	log.Printf("Server listening on http://localhost:%s
", cfg.ServerPort)
	log.Fatal(server.ListenAndServe())
}
