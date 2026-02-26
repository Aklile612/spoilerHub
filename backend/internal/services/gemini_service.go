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
	prompt := fmt.Sprintf(`You are an elite film analyst writing for a premium movie spoiler platform.

Movie Title: %s
Release Year: %s

You MUST structure your response using EXACTLY these markdown headings. Do NOT skip any section.

## Movie Overview
Write a compelling 2-3 sentence non-spoiler summary that hooks the reader.

## ⚠️ SPOILER WARNING
Write exactly: "Everything below contains major plot spoilers, twists, and ending details."

## The Beginning
Describe the setup, world-building, and introduction of main characters. 2-3 paragraphs.

## Major Turning Point
Describe the key event that changes everything. What shifts? What revelation occurs? 2-3 paragraphs.

## The Climax
Describe the peak conflict, major confrontations, and pivotal decisions. 2-3 paragraphs.

## Ending Explained
Explain the ending in detail. If ambiguous, provide multiple interpretations. 2-3 paragraphs.

## Post-Credit Scene
If there is a post-credit scene, describe it. If not, write "This film does not have a post-credit scene."

## Key Moments
List exactly 5 pivotal scenes. Format each as:
- **[Scene Title]** — One sentence description of what happens and why it matters.

## Character Fates
List the main characters (up to 6). Format each as:
- **[Character Name]** | [Actor Name] | [ALIVE/DEAD/UNKNOWN] | One sentence about their arc and final fate.

## What It Really Means
### Symbolism
Explain 2-3 key symbols or motifs in the film.
### Hidden Clues
Describe 2-3 subtle details viewers might have missed.
### Fan Theories
Present 2-3 popular or plausible fan theories.
### Unanswered Questions
List 2-3 questions the film leaves unanswered.

RULES:
- Total length: 800-1200 words.
- Use bold (**text**) for character names and important terms.
- Do NOT fabricate facts — if you're unsure, say so.
- Write in an engaging, editorial tone — like a premium film magazine.
- Every section heading must start with ## exactly as shown above.`, title, year)

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
