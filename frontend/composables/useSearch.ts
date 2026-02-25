import type { Movie, MoviesResponse } from "~/types/movie";

/**
 * Composable for searching movies via /api/search?q=term
 */
export function useSearch() {
  const config = useRuntimeConfig();
  const baseUrl = config.public.apiBaseUrl as string;

  const results = ref<Movie[]>([]);
  const loading = ref(false);
  const error = ref<string | null>(null);
  const query = ref("");

  async function search(q: string) {
    if (!q.trim()) {
      results.value = [];
      return;
    }

    query.value = q;
    loading.value = true;
    error.value = null;

    try {
      const data = await $fetch<MoviesResponse>(`${baseUrl}/api/search`, {
        query: { q },
      });
      results.value = data.movies ?? [];
    } catch (err: unknown) {
      error.value =
        err instanceof Error ? err.message : "Failed to search movies";
      results.value = [];
    } finally {
      loading.value = false;
    }
  }

  return { results, loading, error, query, search };
}
