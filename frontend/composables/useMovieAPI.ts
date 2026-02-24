// Composable for API interactions
import { ref } from "vue";

export interface Movie {
  title: string;
  year: string;
  poster: string;
  backdrop: string;
  rating: number;
  genres: string[];
  overview: string;
  spoiler: string;
}

export const useMovieAPI = () => {
  const config = useRuntimeConfig();
  const apiBaseUrl = config.public.apiBaseUrl;

  // State
  const movie = ref<Movie | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  /**
   * Fetch movie data from backend API
   * @param title - Movie title to search for
   */
  const fetchMovie = async (title: string): Promise<Movie | null> => {
    loading.value = true;
    error.value = null;

    try {
      // Validate input
      if (!title || title.trim() === "") {
        throw new Error("Please enter a movie title");
      }

      // Make API request
      const response = await fetch(
        `${apiBaseUrl}/api/movie?title=${encodeURIComponent(title)}`
      );

      // Handle errors
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to fetch movie");
      }

      // Parse response
      const data = await response.json();
      movie.value = data;

      return data;
    } catch (err) {
      const message =
        err instanceof Error ? err.message : "An unexpected error occurred";
      error.value = message;
      movie.value = null;
      return null;
    } finally {
      loading.value = false;
    }
  };

  /**
   * Clear current movie and error states
   */
  const clearMovie = () => {
    movie.value = null;
    error.value = null;
  };

  return {
    movie,
    loading,
    error,
    fetchMovie,
    clearMovie,
  };
};
