import type { Movie, MoviesResponse } from "~/types/movie";

/**
 * Composable to fetch movies with infinite scroll pagination.
 * Loads page 1 on mount, then appends more pages as user scrolls.
 */
export function useMovies(years: string = "2025,2026") {
  const config = useRuntimeConfig();
  const baseUrl = config.public.apiBaseUrl as string;

  const movies = ref<Movie[]>([]);
  const loading = ref(true);
  const loadingMore = ref(false);
  const error = ref<Error | null>(null);
  const currentPage = ref(0);
  const totalPages = ref(1);
  const hasMore = computed(() => currentPage.value < totalPages.value);

  const isMulti = years.includes(",");
  const baseQuery = isMulti ? { years } : { year: years };

  async function loadPage(page: number) {
    try {
      const data = await $fetch<MoviesResponse>(`${baseUrl}/api/movies`, {
        query: { ...baseQuery, page },
      });

      if (data.movies?.length) {
        movies.value = [...movies.value, ...data.movies];
      }
      currentPage.value = data.page ?? page;
      totalPages.value = data.total_pages ?? 1;
    } catch (err: unknown) {
      error.value = err instanceof Error ? err : new Error("Failed to load movies");
    }
  }

  /** Load the next page */
  async function loadMore() {
    if (!hasMore.value || loadingMore.value) return;
    loadingMore.value = true;
    await loadPage(currentPage.value + 1);
    loadingMore.value = false;
  }

  /** Initial load */
  async function init() {
    loading.value = true;
    await loadPage(1);
    loading.value = false;
  }

  // Auto-init on first call
  init();

  return { movies, loading, loadingMore, error, hasMore, loadMore };
}
