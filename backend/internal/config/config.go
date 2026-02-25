package config

import (
	"os"
)

// Config holds all application configuration
type Config struct {
	Port           string
	TMDBAPIKey     string
	GeminiAPIKey   string
	Environment    string
	SupabaseURL    string
	SupabaseKey    string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		TMDBAPIKey:   getEnv("TMDB_API_KEY", ""),
		GeminiAPIKey: getEnv("GEMINI_API_KEY", ""),
		Environment:  getEnv("ENVIRONMENT", "development"),
		SupabaseURL:  getEnv("SUPABASE_URL", ""),
		SupabaseKey:  getEnv("SUPABASE_KEY", ""),
	}
}

// getEnv retrieves environment variable or returns default
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
