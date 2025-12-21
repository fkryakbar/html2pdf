package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"html2pdf/handlers"
	customMiddleware "html2pdf/middleware"
)

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		// Generate a default API key for development
		apiKey = "dev-api-key-change-in-production"
		log.Printf("WARNING: Using default API key. Set API_KEY environment variable in production.")
	}

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create router
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// API routes with authentication
	r.Route("/api", func(r chi.Router) {
		r.Use(customMiddleware.APIKeyAuth(apiKey))
		r.Post("/convert", handlers.ConvertHTMLToPDF)
	})

	// Health check endpoint (no auth required)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Printf("Server starting on port %s", port)
	log.Printf("API Key: %s", apiKey)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
