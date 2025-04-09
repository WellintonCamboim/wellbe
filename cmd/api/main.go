package main

import (
	"fmt"
	"log"
	"net/http"
	"wellbe/internal/config"
)

func main() {
	// Carregar configurações
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Inicializar servidor
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: setupRoutes(),
	}

	log.Printf("Starting server on port %d", cfg.ServerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}