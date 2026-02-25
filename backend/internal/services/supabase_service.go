package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"spoiler_api/internal/models"
)

// SupabaseService handles Supabase database interactions
type SupabaseService struct {
	baseURL    string
	apiKey     string
	client     *http.Client
}

// NewSupabaseService creates a new Supabase service instance
func NewSupabaseService(supabaseURL, supabaseKey string) *SupabaseService {
	return &SupabaseService{
		baseURL: strings.TrimRight(supabaseURL, "/"),
		apiKey:  supabaseKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// supabaseMovie represents a movie row in the Supabase database
type supabaseMovie struct {
	ID          string   `json:"id,omitempty"`
	Title       string   `json:"title"`
	Year        string   `json:"year"`
	Poster      string   `json:"poster"`
	Backdrop    string   `json:"backdrop"`
	Rating      float64  `json:"rating"`
	Genres      []string `json:"genres"`
	Overview    string   `json:"overview"`
	Spoiler     string   `json:"spoiler"`
	SearchCount int      `json:"search_count"`
	CreatedAt   string   `json:"created_at,omitempty"`
	UpdatedAt   string   `json:"updated_at,omitempty"`
}

// FindMovieByTitleAndYear looks up a movie in the database by title and year
func (s *SupabaseService) FindMovieByTitleAndYear(title, year string) (*models.MovieResponse, error) {
	// Use PostgREST query: filter by title (case-insensitive) and year
	endpoint := fmt.Sprintf("%s/rest/v1/movies?title=ilike.%s&year=eq.%s&limit=1",
		s.baseURL,
		url.QueryEscape(title),
		url.QueryEscape(year),
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	s.setHeaders(req)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query Supabase: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Supabase query error: status %d, response: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Supabase response: %w", err)
	}

	var movies []supabaseMovie
	if err := json.Unmarshal(body, &movies); err != nil {
		return nil, fmt.Errorf("failed to parse Supabase response: %w", err)
	}

	if len(movies) == 0 {
		return nil, nil // Not found — not an error, just no cache hit
	}

	movie := movies[0]

	// Increment search count in the background
	go s.incrementSearchCount(movie.ID)

	return &models.MovieResponse{
		Title:    movie.Title,
		Year:     movie.Year,
		Poster:   movie.Poster,
		Backdrop: movie.Backdrop,
		Rating:   movie.Rating,
		Genres:   movie.Genres,
		Overview: movie.Overview,
		Spoiler:  movie.Spoiler,
	}, nil
}

// SaveMovie stores a movie with its spoiler in the database
func (s *SupabaseService) SaveMovie(movie *models.MovieResponse) error {
	record := supabaseMovie{
		Title:       movie.Title,
		Year:        movie.Year,
		Poster:      movie.Poster,
		Backdrop:    movie.Backdrop,
		Rating:      movie.Rating,
		Genres:      movie.Genres,
		Overview:    movie.Overview,
		Spoiler:     movie.Spoiler,
		SearchCount: 1,
	}

	jsonBody, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("failed to marshal movie for Supabase: %w", err)
	}

	endpoint := fmt.Sprintf("%s/rest/v1/movies", s.baseURL)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create Supabase request: %w", err)
	}

	s.setHeaders(req)
	req.Header.Set("Prefer", "resolution=merge-duplicates") // Upsert on conflict

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to save to Supabase: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Supabase save error: status %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}

// incrementSearchCount increments the search_count for a movie by ID
func (s *SupabaseService) incrementSearchCount(movieID string) {
	// Use Supabase RPC to increment the counter
	endpoint := fmt.Sprintf("%s/rest/v1/rpc/increment_search_count", s.baseURL)

	payload := map[string]string{"movie_id": movieID}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return
	}

	s.setHeaders(req)

	resp, err := s.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

// GetAllMovies retrieves all cached movies from the database
func (s *SupabaseService) GetAllMovies() ([]models.MovieResponse, error) {
	endpoint := fmt.Sprintf("%s/rest/v1/movies?order=search_count.desc&limit=50", s.baseURL)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	s.setHeaders(req)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query Supabase: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Supabase query error: status %d, response: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Supabase response: %w", err)
	}

	var dbMovies []supabaseMovie
	if err := json.Unmarshal(body, &dbMovies); err != nil {
		return nil, fmt.Errorf("failed to parse Supabase response: %w", err)
	}

	var movies []models.MovieResponse
	for _, m := range dbMovies {
		movies = append(movies, models.MovieResponse{
			Title:    m.Title,
			Year:     m.Year,
			Poster:   m.Poster,
			Backdrop: m.Backdrop,
			Rating:   m.Rating,
			Genres:   m.Genres,
			Overview: m.Overview,
			Spoiler:  m.Spoiler,
		})
	}

	return movies, nil
}

// setHeaders sets the required Supabase headers on a request
func (s *SupabaseService) setHeaders(req *http.Request) {
	req.Header.Set("apikey", s.apiKey)
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")
}
