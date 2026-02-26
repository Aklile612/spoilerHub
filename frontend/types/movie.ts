/**
 * TypeScript interfaces matching the backend API responses.
 * Keep in sync with backend/internal/models/movie.go
 */

/** Single movie — returned by /api/movie and as items in lists */
export interface Movie {
  title: string;
  year: string;
  poster: string;
  backdrop: string;
  rating: number;
  genres: string[];
  overview: string;
  spoiler?: string; // Only present when fetched with spoiler
  runtime?: number; // Runtime in minutes (from TMDB detail endpoint)
}

/** Parsed character fate from Gemini spoiler text */
export interface CharacterFate {
  name: string;
  actor: string;
  status: "ALIVE" | "DEAD" | "UNKNOWN";
  summary: string;
}

/** Parsed key moment from Gemini spoiler text */
export interface KeyMoment {
  title: string;
  description: string;
}

/** Parsed story section from Gemini spoiler text */
export interface StorySection {
  id: string;
  title: string;
  icon: string;
  content: string;
}

/** Parsed interpretation section */
export interface InterpretationSection {
  title: string;
  content: string;
}

/** Response shape from /api/movies and /api/search */
export interface MoviesResponse {
  movies: Movie[];
  count: number;
  year?: string;
  query?: string;
  page?: number;
  total_pages?: number;
}

/** Response shape from /api/trending */
export interface TrendingResponse {
  movies: Movie[];
  count: number;
}

/** Backend error response */
export interface ApiError {
  error: string;
}
