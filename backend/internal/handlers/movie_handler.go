package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"spoiler_api/internal/models"
	"spoiler_api/internal/services"
)

// MovieHandler handles movie-related API requests
type MovieHandler struct {
	tmdbService     *services.TMDBService
	geminiService   *services.GeminiService
	supabaseService *services.SupabaseService
}

// NewMovieHandler creates a new movie handler
func NewMovieHandler(tmdbService *services.TMDBService, geminiService *services.GeminiService, supabaseService *services.SupabaseService) *MovieHandler {
	return &MovieHandler{
		tmdbService:     tmdbService,
		geminiService:   geminiService,
		supabaseService: supabaseService,
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

	// Search for movie on TMDB first to get the canonical title and year
	tmdbMovie, err := h.tmdbService.SearchMovie(title)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Extract year from release date
	year := h.tmdbService.ExtractYear(tmdbMovie.ReleaseDate)

	// Step 1: Check Supabase database for cached result
	if h.supabaseService != nil {
		cachedMovie, err := h.supabaseService.FindMovieByTitleAndYear(tmdbMovie.Title, year)
		if err != nil {
			log.Printf("Supabase lookup warning: %v", err)
		} else if cachedMovie != nil {
			log.Printf("Cache HIT: serving '%s (%s)' from Supabase", cachedMovie.Title, cachedMovie.Year)
			c.JSON(http.StatusOK, cachedMovie)
			return
		}
	}

	log.Printf("Cache MISS: generating spoiler for '%s (%s)' via Gemini", tmdbMovie.Title, year)

	// Get genres mapping
	genreMap, err := h.tmdbService.GetGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "failed to fetch genres",
		})
		return
	}

	// Extract genre names
	genres := h.tmdbService.ExtractGenreNames(tmdbMovie.GenreIDs, genreMap)

	// Step 2: Generate spoiler explanation using Gemini (only on cache miss)
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

	// Step 3: Save to Supabase in the background
	if h.supabaseService != nil {
		go func() {
			if err := h.supabaseService.SaveMovie(&response); err != nil {
				log.Printf("Failed to save movie to Supabase: %v", err)
			} else {
				log.Printf("Saved '%s (%s)' to Supabase", response.Title, response.Year)
			}
		}()
	}

	// Return successful response
	c.JSON(http.StatusOK, response)
}

// GetTrendingMovies returns the most searched movies from the database
func (h *MovieHandler) GetTrendingMovies(c *gin.Context) {
	if h.supabaseService == nil {
		c.JSON(http.StatusServiceUnavailable, models.ErrorResponse{
			Error: "database not configured",
		})
		return
	}

	movies, err := h.supabaseService.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: fmt.Sprintf("failed to fetch trending movies: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movies": movies,
		"count":  len(movies),
	})
}

// HealthCheck handles the health check endpoint
func (h *MovieHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":     "healthy",
		"cache_size": h.geminiService.GetCacheSize(),
		"database":   h.supabaseService != nil,
	})
}
