package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"spoiler_api/internal/models"
)

// GeminiService handles Gemini API interactions with caching
type GeminiService struct {
	apiKey string
	client *http.Client
	cache  map[string]string
	mu     sync.RWMutex
}

// NewGeminiService creates a new Gemini service instance
func NewGeminiService(apiKey string) *GeminiService {
	return &GeminiService{
		apiKey: apiKey,
		client: &http.Client{},
		cache:  make(map[string]string),
	}
}

// GenerateSpoiler generates a detailed spoiler explanation for a movie
func (s *GeminiService) GenerateSpoiler(title, year string) (string, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("%s_%s", title, year)
	
	s.mu.RLock()
	if cachedSpoiler, exists := s.cache[cacheKey]; exists {
		s.mu.RUnlock()
		return cachedSpoiler, nil
	}
	s.mu.RUnlock()

	// Construct the prompt
	prompt := s.constructPrompt(title, year)

	// Create Gemini API request
	request := models.GeminiRequest{
		Contents: []models.GeminiContent{
			{
				Parts: []models.GeminiPart{
					{
						Text: prompt,
					},
				},
			},
		},
	}

	// Marshal request to JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal Gemini request: %w", err)
	}

	// Make request to Gemini API
	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1/models/gemini-2.5-flash:generateContent?key=%s",
		s.apiKey,
	)

	resp, err := s.client.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to call Gemini API: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Gemini API error: status code %d, response: %s", resp.StatusCode, string(body))
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read Gemini response: %w", err)
	}

	// Parse JSON response
	var geminiResp models.GeminiResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		return "", fmt.Errorf("failed to parse Gemini response: %w", err)
	}

	// Extract text from response
	if len(geminiResp.Candidates) == 0 {
		return "", fmt.Errorf("no candidates in Gemini response")
	}

	spoilerText := ""
	if len(geminiResp.Candidates[0].Content.Parts) > 0 {
		spoilerText = geminiResp.Candidates[0].Content.Parts[0].Text
	}

	// Cache the result
	s.mu.Lock()
	s.cache[cacheKey] = spoilerText
	s.mu.Unlock()

	return spoilerText, nil
}

// constructPrompt creates a detailed prompt for Gemini
func (s *GeminiService) constructPrompt(title, year string) string {
	prompt := fmt.Sprintf(`You are a professional film analyst and movie explainer.

Movie Title: %s
Release Year: %s

Instructions:
1. Provide a short non-spoiler summary.
2. Clearly display:
   "⚠️ SPOILER WARNING"
3. Provide:
   - Full plot breakdown
   - Major twists
   - Character arcs
   - Ending explanation
   - Hidden clues
   - If ambiguous ending, include interpretations
4. Keep explanation between 600-1000 words.
5. Structure output with markdown headings:
   - Movie Overview
   - Full Spoiler Breakdown
   - Ending Explained
   - Themes & Meaning

Do not fabricate unknown facts.`, title, year)

	return prompt
}

// ClearCache clears the spoiler cache (useful for testing or admin operations)
func (s *GeminiService) ClearCache() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache = make(map[string]string)
}

// GetCacheSize returns the number of cached items
func (s *GeminiService) GetCacheSize() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.cache)
}
