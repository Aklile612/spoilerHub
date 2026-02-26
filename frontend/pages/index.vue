<script setup lang="ts">
/**
 * Home page — discovers trending movies for 2025 & 2026.
 * Infinite scroll: loads more movies as user scrolls down.
 */
import type { TrendingResponse } from "~/types/movie";

// Fetch trending movies with pagination
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
let observer: IntersectionObserver | null = null;

// Watch the sentinel ref — it only appears in the DOM after loading=false,
// so we can't rely on onMounted (sentinel is inside a v-if).
watch(scrollSentinel, (el) => {
  // Clean up previous observer if any
  if (observer) {
    observer.disconnect();
    observer = null;
  }

  if (!el) return;

  observer = new IntersectionObserver(
    (entries) => {
      if (entries[0].isIntersecting && hasMore.value && !loadingMore.value) {
        loadMore();
      }
    },
    { rootMargin: "400px" }
  );

  observer.observe(el);
});

onUnmounted(() => {
  if (observer) {
    observer.disconnect();
    observer = null;
  }
});

useHead({
  title: "SpoilerHub — Discover Movie Spoilers",
});

// Genre filter for trending section
const genreFilters = ["All", "Action", "Drama", "Horror", "Comedy", "Thriller"];
const activeGenre = ref("All");

const filteredMovies = computed(() => {
  if (activeGenre.value === "All") return movies.value;
  return movies.value.filter((m) =>
    m.genres?.some((g) => g.toLowerCase() === activeGenre.value.toLowerCase())
  );
});
</script>

<template>
  <div>
    <!-- Hero carousel — 4 featured films -->
    <HeroCarousel />

    <!-- 🔥 Currently Dominating Screens — premium horizontal cards -->
    <section
      v-if="trendingMovies.length > 0"
      class="relative overflow-hidden bg-gradient-to-b from-gray-50 to-white py-10 sm:py-14"
    >
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <!-- Section header -->
        <div class="mb-8 flex items-end justify-between">
          <div>
            <p class="text-[11px] font-bold uppercase tracking-[0.25em] text-brand-dark">
              Curated for you
            </p>
            <h2 class="mt-1 flex items-center gap-2 text-2xl font-black tracking-tight text-gray-900">
              Currently Dominating Searches
              <span class="inline-block h-1.5 w-8 rounded-full bg-brand" />
            </h2>
          </div>
          <span class="rounded-full bg-brand/10 px-3 py-1 text-xs font-bold text-brand-dark">
            {{ trendingMovies.length }} movies
          </span>
        </div>

        <!-- Horizontal scroll — cinematic wide cards -->
        <div class="flex gap-5 overflow-x-auto pb-4 scrollbar-hide">
          <NuxtLink
            v-for="(movie, index) in trendingMovies"
            :key="movie.title + movie.year"
            :to="'/movie/' + encodeURIComponent(movie.title)"
            class="group relative w-72 shrink-0 animate-fade-in overflow-hidden rounded-2xl shadow-lg transition-all duration-300 hover:shadow-2xl hover:shadow-brand/10 hover:-translate-y-1 sm:w-80"
            :style="{ animationDelay: `${index * 80}ms` }"
          >
            <!-- Backdrop image -->
            <div class="relative aspect-[16/10] overflow-hidden bg-gray-800">
              <img
                :src="movie.backdrop || movie.poster"
                :alt="movie.title"
                class="h-full w-full object-cover transition-transform duration-500 group-hover:scale-110"
                loading="lazy"
              />
              <!-- Dark gradient overlay -->
              <div class="absolute inset-0 bg-gradient-to-t from-black via-black/40 to-transparent" />

              <!-- Rating badge — top right -->
              <div class="absolute right-3 top-3">
                <span class="inline-flex items-center gap-1 rounded-lg bg-black/50 px-2 py-0.5 text-xs font-bold text-brand backdrop-blur-sm">
                  <svg class="h-3 w-3" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                  {{ movie.rating.toFixed(1) }}
                </span>
              </div>

              <!-- Rank number -->
              <div class="absolute left-3 top-3">
                <span class="text-3xl font-black leading-none text-white/20" style="text-shadow: 0 2px 8px rgba(0,0,0,0.3)">
                  #{{ index + 1 }}
                </span>
              </div>

              <!-- Info overlay at bottom -->
              <div class="absolute inset-x-0 bottom-0 p-4">
                <!-- Genre pills -->
                <div class="mb-2 flex flex-wrap gap-1.5">
                  <span
                    v-for="genre in (movie.genres || []).slice(0, 2)"
                    :key="genre"
                    class="rounded-full bg-white/15 px-2.5 py-0.5 text-[10px] font-medium text-white/80 backdrop-blur-sm"
                  >
                    {{ genre }}
                  </span>
                  <span class="text-[10px] text-white/40">{{ movie.year }}</span>
                </div>

                <!-- Title -->
                <h3 class="text-lg font-bold leading-tight text-white">
                  {{ movie.title }}
                </h3>

                <!-- Overview preview -->
                <p class="mt-1 line-clamp-2 text-xs leading-relaxed text-white/50">
                  {{ movie.overview }}
                </p>

                <!-- CTA — appears on hover -->
                <div
                  class="mt-3 flex items-center gap-2 opacity-0 transition-all duration-300 group-hover:opacity-100 group-hover:translate-y-0 translate-y-2"
                >
                  <span class="inline-flex items-center gap-1.5 rounded-full bg-brand px-4 py-1.5 text-[11px] font-bold text-black">
                    <svg class="h-3 w-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
                      <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                    See Spoiler
                  </span>
                  <span class="text-[10px] text-white/30">⚠️ contains spoilers</span>
                </div>
              </div>
            </div>
          </NuxtLink>
        </div>
      </div>
    </section>

    <!-- Trending Movies section -->
    <section class="mx-auto max-w-7xl px-4 py-10 sm:px-6 sm:py-14 lg:px-8">
      <!-- Header: title + genre filters -->
      <div class="mb-8 flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
        <div class="flex items-start gap-3">
          <!-- Gold accent bar -->
          <div class="mt-1 h-10 w-1.5 rounded-full bg-brand" />
          <div>
            <h2 class="text-2xl font-black tracking-tight text-gray-900">
              Trending Movies
            </h2>
            <p class="mt-0.5 text-sm text-brand-dark">
              Top rated movies for 2025–2026 season
            </p>
          </div>
        </div>

        <!-- Genre filter pills -->
        <div class="flex flex-wrap gap-2">
          <button
            v-for="genre in genreFilters"
            :key="genre"
            @click="activeGenre = genre"
            class="rounded-full px-4 py-1.5 text-xs font-semibold transition-all duration-200"
            :class="
              activeGenre === genre
                ? 'bg-gray-900 text-white shadow-md'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            "
          >
            {{ genre }}
          </button>
        </div>
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
      <MovieGrid v-else :movies="filteredMovies" :loading="loading" />

      <!-- No results for this genre -->
      <div
        v-if="!loading && !error && filteredMovies.length === 0 && movies.length > 0"
        class="flex flex-col items-center py-16 text-center"
      >
        <span class="text-4xl">🎭</span>
        <p class="mt-3 text-sm font-medium text-gray-400">
          No {{ activeGenre }} movies found — try another genre
        </p>
      </div>

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
          You've reached the end of trending movies 🎬
        </p>

        <!-- Invisible sentinel that triggers loading the next page -->
        <div ref="scrollSentinel" class="h-1 w-full" />
      </div>
    </section>
  </div>
</template>
