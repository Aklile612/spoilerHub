# SpoilerHub Backend API

## Quick Start for Backend

### Prerequisites
- Go 1.21+

### Setup

1. **Copy environment file:**
   ```bash
   cp .env.example .env
   ```

2. **Configure API keys in `.env`:**
   ```
   TMDB_API_KEY=your_tmdb_api_key
   GEMINI_API_KEY=your_gemini_api_key
   ```

3. **Install dependencies:**
   ```bash
   go mod download
   ```

4. **Run server:**
   ```bash
   go run cmd/server/main.go
   ```

The API will start at `http://localhost:8080`

## API Documentation

### GET /health
Health check endpoint returning server status and cache size.

### GET /api/movie?title=MovieTitle
Search for a movie and get detailed information with AI spoilers.

**Response:**
```json
{
  "title": "Movie Title",
  "year": "2024",
  "poster": "https://...",
  "backdrop": "https://...",
  "rating": 8.5,
  "genres": ["Action", "Sci-Fi"],
  "overview": "...",
  "spoiler": "⚠️ SPOILER WARNING\n..."
}
```

## Architecture

- **handlers/** - HTTP request handlers
- **services/** - Business logic and API clients
  - `tmdb_service.go` - TMDB API integration
  - `gemini_service.go` - Gemini API with caching
- **models/** - Data structures
- **routes/** - Route definitions
- **config/** - Configuration management

## Features

- Thread-safe in-memory caching
- CORS middleware
- Comprehensive error handling
- RESTful API design
