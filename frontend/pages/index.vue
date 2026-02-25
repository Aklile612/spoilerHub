<script setup lang="ts">
/**
 * Home page — discovers top-rated movies for 2025 & 2026.
 * Infinite scroll: loads more movies as user scrolls down.
 */
import type { TrendingResponse } from "~/types/movie";

// Fetch top-rated movies with pagination
const { movies, loading, loadingMore, error, hasMore, loadMore } =
  useMovies("2025,2026");

// Fetch trending (cached) movies from Supabase
const config = useRuntimeConfig();
const baseUrl = config.public.apiBaseUrl as string;

const { data: trendingData, pending: trendingLoading } =
  useFetch<TrendingResponse>(`${baseUrl}/api/trending`, {
    key: "trending",
    default: () => ({ movies: [], count: 0 }),
  });

const trendingMovies = computed(() => trendingData.value?.movies ?? []);

// Infinite scroll: observe a sentinel element at the bottom of the grid
const scrollSentinel = ref<HTMLElement | null>(null);

onMounted(() => {
  if (!scrollSentinel.value) return;

  const observer = new IntersectionObserver(
    (entries) => {
      if (entries[0].isIntersecting && hasMore.value && !loadingMore.value) {
        loadMore();
      }
    },
    { rootMargin: "200px" }
  );

  observer.observe(scrollSentinel.value);

  onUnmounted(() => observer.disconnect());
});

useHead({
  title: "SpoilerHub — Discover Movie Spoilers",
});
</script>

<template>
  <div>
    <!-- Hero section -->
    <section class="bg-gradient-to-b from-gray-50 to-white py-12 sm:py-16">
      <div class="mx-auto max-w-7xl px-4 text-center sm:px-6 lg:px-8">
        <h1
          class="animate-fade-in text-3xl font-extrabold tracking-tight text-gray-900 sm:text-4xl lg:text-5xl"
        >
          🍿 Discover Movie
          <span class="text-brand">Spoilers</span>
        </h1>
        <p
          class="animate-fade-in mx-auto mt-4 max-w-2xl text-base text-gray-500 sm:text-lg"
          style="animation-delay: 100ms"
        >
          Search any movie and get AI-generated plot breakdowns, ending
          explanations, and thematic analysis.
        </p>
      </div>
    </section>

    <!-- Trending movies section (from Supabase cache) -->
    <section
      v-if="trendingMovies.length > 0"
      class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8"
    >
      <div class="mb-6 flex items-center gap-2">
        <span class="text-xl">🔥</span>
        <h2 class="text-xl font-bold text-gray-900">Popular Spoilers</h2>
        <span
          class="rounded-full bg-brand/10 px-2 py-0.5 text-xs font-semibold text-brand-dark"
        >
          {{ trendingMovies.length }}
        </span>
      </div>
      <MovieGrid :movies="trendingMovies" :loading="trendingLoading" />
    </section>

    <!-- Divider (only show if both sections have content) -->
    <div
      v-if="trendingMovies.length > 0"
      class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8"
    >
      <hr class="border-gray-100" />
    </div>

    <!-- 2026 Movies section -->
    <section class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
      <div class="mb-6 flex items-center gap-2">
        <span class="text-xl">🎬</span>
        <h2 class="text-xl font-bold text-gray-900">Top Rated — 2025 & 2026</h2>
      </div>

      <!-- Error state -->
      <div
        v-if="error"
        class="rounded-lg border border-red-200 bg-red-50 p-6 text-center"
      >
        <p class="text-sm font-medium text-red-800">
          Failed to load movies. Please try again later.
        </p>
        <p class="mt-1 text-xs text-red-600">{{ error.message }}</p>
      </div>

      <!-- Movie grid -->
      <MovieGrid v-else :movies="movies" :loading="loading" />

      <!-- Load more indicator & scroll sentinel -->
      <div v-if="!loading && !error" class="mt-8 flex flex-col items-center">
        <!-- Loading spinner when fetching next page -->
        <div
          v-if="loadingMore"
          class="flex items-center gap-2 text-sm text-gray-400"
        >
          <svg
            class="h-5 w-5 animate-spin text-brand"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              class="opacity-25"
              cx="12" cy="12" r="10"
              stroke="currentColor" stroke-width="4"
            />
            <path
              class="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
            />
          </svg>
          Loading more movies…
        </div>

        <!-- End of results -->
        <p
          v-else-if="!hasMore && movies.length > 0"
          class="text-sm text-gray-300"
        >
          You've seen all the top-rated movies 🎬
        </p>

        <!-- Invisible sentinel that triggers loading the next page -->
        <div ref="scrollSentinel" class="h-1 w-full" />
      </div>
    </section>
  </div>
</template>
