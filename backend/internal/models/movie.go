package models

// MovieResponse represents the API response for a movie with spoiler details
type MovieResponse struct {
	Title    string   `json:"title"`
	Year     string   `json:"year"`
	Poster   string   `json:"poster"`
	Backdrop string   `json:"backdrop"`
	Rating   float64  `json:"rating"`
	Genres   []string `json:"genres"`
	Overview string   `json:"overview"`
	Spoiler  string   `json:"spoiler"`
}

// TMDBSearchResult represents the TMDB API search response
type TMDBSearchResult struct {
	Results []TMDBMovie `json:"results"`
}

// TMDBMovie represents a single movie from TMDB API
type TMDBMovie struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	ReleaseDate   string `json:"release_date"`
	PosterPath    string `json:"poster_path"`
	BackdropPath  string `json:"backdrop_path"`
	VoteAverage   float64 `json:"vote_average"`
	Overview      string `json:"overview"`
	GenreIDs      []int  `json:"genre_ids"`
}

// TMDBGenreResponse represents the genres from TMDB API
type TMDBGenreResponse struct {
	Genres []TMDBGenre `json:"genres"`
}

// TMDBGenre represents a genre from TMDB API
type TMDBGenre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GeminiRequest represents the request structure for Gemini API
type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

// GeminiContent represents content for Gemini API
type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
}

// GeminiPart represents a text part for Gemini API
type GeminiPart struct {
	Text string `json:"text"`
}

// GeminiResponse represents the response from Gemini API
type GeminiResponse struct {
	Candidates []GeminiCandidate `json:"candidates"`
}

// GeminiCandidate represents a candidate response from Gemini API
type GeminiCandidate struct {
	Content GeminiContent `json:"content"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}
