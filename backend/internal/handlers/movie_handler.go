package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"spoiler_api/internal/models"
	"spoiler_api/internal/services"
)

// MovieHandler handles movie-related API requests
type MovieHandler struct {
	tmdbService   *services.TMDBService
	geminiService *services.GeminiService
}

// NewMovieHandler creates a new movie handler
func NewMovieHandler(tmdbService *services.TMDBService, geminiService *services.GeminiService) *MovieHandler {
	return &MovieHandler{
		tmdbService:   tmdbService,
		geminiService: geminiService,
	}
}

// GetMovie handles the GET /api/movie endpoint
// @Summary Get movie details with spoiler explanation
// @Description Searches for a movie and generates AI spoiler explanation
// @Param title query string true "Movie title to search for"
// @Produce json
// @Success 200 {object} models.MovieResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *MovieHandler) GetMovie(c *gin.Context) {
	// Get title from query parameters
	title := c.Query("title")

	// Validate query parameter
	if title == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "title query parameter is required",
		})
		return
	}

	// Search for movie on TMDB
	tmdbMovie, err := h.tmdbService.SearchMovie(title)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Get genres mapping
	genreMap, err := h.tmdbService.GetGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "failed to fetch genres",
		})
		return
	}

	// Extract year from release date
	year := h.tmdbService.ExtractYear(tmdbMovie.ReleaseDate)

	// Extract genre names
	genres := h.tmdbService.ExtractGenreNames(tmdbMovie.GenreIDs, genreMap)

	// Generate spoiler explanation using Gemini
	spoiler, err := h.geminiService.GenerateSpoiler(tmdbMovie.Title, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: fmt.Sprintf("failed to generate spoiler explanation: %v", err),
		})
		return
	}

	// Build response
	response := models.MovieResponse{
		Title:    tmdbMovie.Title,
		Year:     year,
		Poster:   h.tmdbService.FormatPosterURL(tmdbMovie.PosterPath),
		Backdrop: h.tmdbService.FormatBackdropURL(tmdbMovie.BackdropPath),
		Rating:   tmdbMovie.VoteAverage,
		Genres:   genres,
		Overview: h.tmdbService.TruncateOverview(tmdbMovie.Overview, 500),
		Spoiler:  spoiler,
	}

	// Return successful response
	c.JSON(http.StatusOK, response)
}

// HealthCheck handles the health check endpoint
func (h *MovieHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"cache_size": h.geminiService.GetCacheSize(),
	})
}
