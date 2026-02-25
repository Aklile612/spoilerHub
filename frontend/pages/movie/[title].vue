<script setup lang="ts">
/**
 * Movie detail page — shows full movie info + expandable spoiler.
 * Route: /movie/[title]
 * Fetches from GET /api/movie?title=...
 */
import type { Movie } from "~/types/movie";

const route = useRoute();
const config = useRuntimeConfig();
const baseUrl = config.public.apiBaseUrl as string;

// Get movie title from route param
const movieTitle = computed(() =>
  decodeURIComponent(route.params.title as string)
);

// Fetch movie with spoiler
const {
  data: movie,
  pending: loading,
  error,
} = useFetch<Movie>(`${baseUrl}/api/movie`, {
  query: { title: movieTitle.value },
  key: `movie-${movieTitle.value}`,
});

// Dynamic page title
useHead({
  title: computed(
    () =>
      movie.value
        ? `${movie.value.title} (${movie.value.year}) — SpoilerHub`
        : "Loading… — SpoilerHub"
  ),
});

/** Backdrop fallback */
const backdropFallback =
  "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='1280' height='720' fill='%23111827'%3E%3Crect width='1280' height='720'/%3E%3C/svg%3E";

const posterFallback =
  "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='300' height='450' fill='%23e5e7eb'%3E%3Crect width='300' height='450'/%3E%3Ctext x='50%25' y='50%25' dominant-baseline='middle' text-anchor='middle' font-family='sans-serif' font-size='18' fill='%239ca3af'%3ENo Poster%3C/text%3E%3C/svg%3E";
</script>

<template>
  <div>
    <!-- Loading state -->
    <div v-if="loading" class="animate-pulse">
      <!-- Skeleton backdrop -->
      <div class="skeleton h-64 w-full sm:h-80 lg:h-96" />
      <div class="mx-auto max-w-5xl px-4 py-8 sm:px-6 lg:px-8">
        <div class="flex flex-col gap-6 sm:flex-row">
          <div class="skeleton h-72 w-48 shrink-0 rounded-xl" />
          <div class="flex-1 space-y-4">
            <div class="skeleton h-8 w-2/3 rounded" />
            <div class="skeleton h-4 w-1/4 rounded" />
            <div class="flex gap-2">
              <div class="skeleton h-6 w-16 rounded-full" />
              <div class="skeleton h-6 w-20 rounded-full" />
            </div>
            <div class="skeleton h-24 w-full rounded" />
          </div>
        </div>
      </div>
    </div>

    <!-- Error state -->
    <div
      v-else-if="error"
      class="flex min-h-[60vh] flex-col items-center justify-center px-4 text-center"
    >
      <span class="text-5xl">😕</span>
      <h2 class="mt-4 text-xl font-bold text-gray-900">Movie Not Found</h2>
      <p class="mt-2 text-sm text-gray-500">
        We couldn't find details for "{{ movieTitle }}".
      </p>
      <NuxtLink
        to="/"
        class="mt-6 inline-flex items-center gap-2 rounded-lg bg-brand px-5 py-2.5 text-sm font-semibold text-gray-900 transition-colors hover:bg-brand-dark"
      >
        ← Back to Home
      </NuxtLink>
    </div>

    <!-- Movie content -->
    <div v-else-if="movie" class="animate-fade-in">
      <!-- Backdrop hero — fixed background, fades to white -->
      <div class="relative overflow-hidden">
        <!-- Backdrop image (absolute, behind everything) -->
        <div class="absolute inset-0 h-[420px] sm:h-[480px]">
          <img
            :src="movie.backdrop || backdropFallback"
            :alt="`${movie.title} backdrop`"
            class="h-full w-full object-cover"
          />
          <!-- Heavy gradient: fades to white so content below is readable -->
          <div
            class="absolute inset-0 bg-gradient-to-b from-black/50 via-black/30 to-white"
          />
        </div>

        <!-- Back button (inside the backdrop area) -->
        <div class="relative z-10 px-4 pt-4 sm:px-6">
          <NuxtLink
            to="/"
            class="inline-flex items-center gap-1 rounded-full bg-white/20 px-3 py-1.5 text-xs font-medium text-white backdrop-blur-sm transition-colors hover:bg-white/30"
          >
            ← Back
          </NuxtLink>
        </div>

        <!-- Movie details — sits on top of the faded backdrop -->
        <div class="relative z-10 mx-auto max-w-5xl px-4 pt-16 sm:px-6 sm:pt-24 lg:px-8">
          <!-- Poster + info row -->
          <div class="flex flex-col gap-6 sm:flex-row sm:items-end">
            <!-- Poster -->
            <div class="shrink-0">
              <img
                :src="movie.poster || posterFallback"
                :alt="`${movie.title} poster`"
                class="h-auto w-40 rounded-xl border-4 border-white shadow-xl sm:w-48"
              />
            </div>

            <!-- Info -->
            <div class="pb-2">
              <h1
                class="text-2xl font-extrabold text-gray-900 drop-shadow-sm sm:text-3xl lg:text-4xl"
              >
                {{ movie.title }}
              </h1>

              <div class="mt-2 flex flex-wrap items-center gap-3">
                <!-- Year -->
                <span class="text-sm text-gray-500">{{ movie.year }}</span>

                <!-- Rating -->
                <RatingBadge :rating="movie.rating" />
              </div>

              <!-- Genres -->
              <div
                v-if="movie.genres?.length"
                class="mt-3 flex flex-wrap gap-2"
              >
                <span
                  v-for="genre in movie.genres"
                  :key="genre"
                  class="rounded-full border border-gray-200 bg-white/80 px-3 py-1 text-xs font-medium text-gray-600 backdrop-blur-sm"
                >
                  {{ genre }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Content below hero (on white background) -->
      <div class="mx-auto max-w-5xl px-4 sm:px-6 lg:px-8">

        <!-- Overview -->
        <section class="mt-8">
          <h2 class="text-lg font-bold text-gray-900">Overview</h2>
          <p class="mt-2 leading-relaxed text-gray-600">
            {{ movie.overview || "No overview available." }}
          </p>
        </section>

        <!-- Spoiler toggle section -->
        <SpoilerToggle
          v-if="movie.spoiler"
          :spoiler="movie.spoiler"
        />

        <!-- No spoiler available -->
        <div
          v-else
          class="mt-8 rounded-lg border border-gray-200 bg-gray-50 p-6 text-center"
        >
          <p class="text-sm text-gray-500">
            No AI-generated spoiler available for this movie yet.
          </p>
        </div>

        <!-- Bottom spacing -->
        <div class="h-12" />
      </div>
    </div>
  </div>
</template>
