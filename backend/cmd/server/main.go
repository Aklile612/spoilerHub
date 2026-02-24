package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"spoiler_api/internal/config"
	"spoiler_api/internal/handlers"
	"spoiler_api/internal/routes"
	"spoiler_api/internal/services"
)

func main() {
	// Load configuration from environment variables
	cfg := config.LoadConfig()

	// Validate required environment variables
	if cfg.TMDBAPIKey == "" {
		log.Fatal("TMDB_API_KEY environment variable is required")
	}
	if cfg.GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is required")
	}

	// Set Gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(corsMiddleware())

	// Initialize services
	tmdbService := services.NewTMDBService(cfg.TMDBAPIKey)
	geminiService := services.NewGeminiService(cfg.GeminiAPIKey)

	// Initialize handlers
	movieHandler := handlers.NewMovieHandler(tmdbService, geminiService)

	// Setup routes
	routes.SetupRoutes(router, movieHandler)

	// Start server
	address := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting SpoilerHub API server on %s\n", address)

	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}

// corsMiddleware adds CORS headers to responses
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
