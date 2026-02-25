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
