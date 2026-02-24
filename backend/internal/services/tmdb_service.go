package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"spoiler_api/internal/models"
)

// TMDBService handles TMDB API interactions
type TMDBService struct {
	apiKey string
	client *http.Client
}

// NewTMDBService creates a new TMDB service instance
func NewTMDBService(apiKey string) *TMDBService {
	return &TMDBService{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

// SearchMovie searches for a movie by title on TMDB
func (s *TMDBService) SearchMovie(title string) (*models.TMDBMovie, error) {
	// URL encode the title
	encodedTitle := url.QueryEscape(title)
	
	// Construct TMDB search endpoint
	searchURL := fmt.Sprintf(
		"https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s",
		s.apiKey,
		encodedTitle,
	)

	// Make request to TMDB
	resp, err := s.client.Get(searchURL)
	if err != nil {
		return nil, fmt.Errorf("failed to search TMDB: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API error: status code %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read TMDB response: %w", err)
	}

	// Parse JSON response
	var searchResult models.TMDBSearchResult
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return nil, fmt.Errorf("failed to parse TMDB response: %w", err)
	}

	// Check if any results were found
	if len(searchResult.Results) == 0 {
		return nil, fmt.Errorf("no movies found for title: %s", title)
	}

	// Return the first (best match) result
	return &searchResult.Results[0], nil
}

// GetGenres retrieves all genres from TMDB
func (s *TMDBService) GetGenres() (map[int]string, error) {
	// Construct genres endpoint
	genresURL := fmt.Sprintf(
		"https://api.themoviedb.org/3/genre/movie/list?api_key=%s",
		s.apiKey,
	)

	// Make request to TMDB
	resp, err := s.client.Get(genresURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch genres: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB genres API error: status code %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read genres response: %w", err)
	}

	// Parse JSON response
	var genreResponse models.TMDBGenreResponse
	if err := json.Unmarshal(body, &genreResponse); err != nil {
		return nil, fmt.Errorf("failed to parse genres response: %w", err)
	}

	// Convert genre slice to map for easy lookup
	genreMap := make(map[int]string)
	for _, genre := range genreResponse.Genres {
		genreMap[genre.ID] = genre.Name
	}

	return genreMap, nil
}

// FormatPosterURL formats the poster URL
func (s *TMDBService) FormatPosterURL(posterPath string) string {
	if posterPath == "" {
		return ""
	}
	return fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", posterPath)
}

// FormatBackdropURL formats the backdrop URL
func (s *TMDBService) FormatBackdropURL(backdropPath string) string {
	if backdropPath == "" {
		return ""
	}
	return fmt.Sprintf("https://image.tmdb.org/t/p/w1280%s", backdropPath)
}

// ExtractYear extracts year from release date
func (s *TMDBService) ExtractYear(releaseDate string) string {
	if len(releaseDate) >= 4 {
		return releaseDate[:4]
	}
	return ""
}

// ExtractGenreNames extracts genre names from genre IDs
func (s *TMDBService) ExtractGenreNames(genreIDs []int, genreMap map[int]string) []string {
	var genres []string
	for _, id := range genreIDs {
		if name, exists := genreMap[id]; exists {
			genres = append(genres, name)
		}
	}
	return genres
}

// TruncateOverview truncates overview to first 500 characters
func (s *TMDBService) TruncateOverview(overview string, maxLength int) string {
	if len(overview) > maxLength {
		return strings.TrimSpace(overview[:maxLength]) + "..."
	}
	return overview
}
