<script setup lang="ts">
/**
 * MovieGrid — displays a responsive grid of MovieCard components.
 * Reused on home page, search page, and trending section.
 */
import type { Movie } from "~/types/movie";

defineProps<{
  movies: Movie[];
  loading?: boolean;
}>();
</script>

<template>
  <!-- Loading skeleton grid -->
  <div
    v-if="loading"
    class="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6"
  >
    <div
      v-for="i in 12"
      :key="i"
      class="overflow-hidden rounded-xl border border-gray-100"
    >
      <div class="skeleton aspect-[2/3]" />
      <div class="space-y-2 p-3">
        <div class="skeleton h-4 w-3/4 rounded" />
        <div class="skeleton h-3 w-1/3 rounded" />
        <div class="flex gap-1">
          <div class="skeleton h-4 w-12 rounded-full" />
          <div class="skeleton h-4 w-10 rounded-full" />
        </div>
      </div>
    </div>
  </div>

  <!-- Empty state -->
  <div
    v-else-if="movies.length === 0"
    class="flex flex-col items-center justify-center py-20 text-center"
  >
    <span class="text-5xl" aria-hidden="true">🎬</span>
    <p class="mt-4 text-lg font-medium text-gray-500">No movies found</p>
    <p class="mt-1 text-sm text-gray-400">
      Try searching for a different movie
    </p>
  </div>

  <!-- Movie grid -->
  <div
    v-else
    class="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6"
  >
    <MovieCard
      v-for="(movie, index) in movies"
      :key="movie.title + movie.year"
      :movie="movie"
      class="animate-fade-in"
      :style="{ animationDelay: `${index * 50}ms` }"
    />
  </div>
</template>
