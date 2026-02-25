package routes

import (
	"github.com/gin-gonic/gin"

	"spoiler_api/internal/handlers"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, movieHandler *handlers.MovieHandler) {
	// Health check endpoint
	router.GET("/health", movieHandler.HealthCheck)

	// API routes
	api := router.Group("/api")
	{
		// Single movie with spoiler
		api.GET("/movie", movieHandler.GetMovie)

		// Discover movies by year
		api.GET("/movies", movieHandler.DiscoverMovies)

		// Search movies (no spoiler generation)
		api.GET("/search", movieHandler.SearchMovies)

		// Trending/cached movies endpoint
		api.GET("/trending", movieHandler.GetTrendingMovies)
	}
}
