import type { Movie } from "~/types/movie";

/**
 * Composable to fetch a single movie with spoiler by title.
 * Used on the movie detail page.
 */
export function useMovie(title: Ref<string> | string) {
  const config = useRuntimeConfig();
  const baseUrl = config.public.apiBaseUrl as string;
  const titleValue = toValue(title);

  const {
    data: movie,
    pending: loading,
    error,
    refresh,
  } = useFetch<Movie>(`${baseUrl}/api/movie`, {
    query: { title: titleValue },
    key: `movie-${titleValue}`,
    // Don't fetch if title is empty
    immediate: !!titleValue,
  });

  return { movie, loading, error, refresh };
}
