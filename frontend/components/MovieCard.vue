<script setup lang="ts">
import type { Movie } from "~/types/movie";

/**
 * MovieCard — a single movie card in the grid.
 * Clicking navigates to the movie detail page.
 */
defineProps<{
  movie: Movie;
}>();

/** Placeholder for missing posters */
const posterFallback =
  "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='300' height='450' fill='%23e5e7eb'%3E%3Crect width='300' height='450'/%3E%3Ctext x='50%25' y='50%25' dominant-baseline='middle' text-anchor='middle' font-family='sans-serif' font-size='18' fill='%239ca3af'%3ENo Poster%3C/text%3E%3C/svg%3E";
</script>

<template>
  <NuxtLink
    :to="{
      path: '/movie/' + encodeURIComponent(movie.title),
    }"
    class="group block rounded-xl bg-white border border-gray-100 overflow-hidden shadow-sm transition-all duration-300 hover:shadow-lg hover:-translate-y-1 hover:border-gray-200 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-brand"
  >
    <!-- Poster image with rating badge overlay -->
    <div class="relative aspect-[2/3] overflow-hidden bg-gray-100">
      <img
        :src="movie.poster || posterFallback"
        :alt="`${movie.title} poster`"
        class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
        loading="lazy"
      />
      <!-- Rating badge — top right -->
      <div class="absolute right-2 top-2">
        <RatingBadge :rating="movie.rating" />
      </div>
    </div>

    <!-- Card body -->
    <div class="p-3">
      <!-- Title -->
      <h3
        class="line-clamp-1 text-sm font-semibold text-gray-900 group-hover:text-brand-dark"
      >
        {{ movie.title }}
      </h3>

      <!-- Year -->
      <p class="mt-0.5 text-xs text-gray-500">{{ movie.year }}</p>

      <!-- Genres -->
      <div v-if="movie.genres?.length" class="mt-1.5 flex flex-wrap gap-1">
        <span
          v-for="genre in movie.genres.slice(0, 3)"
          :key="genre"
          class="rounded-full bg-gray-100 px-2 py-0.5 text-[10px] font-medium text-gray-600"
        >
          {{ genre }}
        </span>
      </div>

      <!-- See Spoiler CTA -->
      <div
        class="mt-2.5 flex items-center gap-1 text-[11px] font-semibold text-brand-dark opacity-0 transition-opacity duration-300 group-hover:opacity-100"
      >
        <span>⚠️</span>
        <span>See Spoiler →</span>
      </div>
    </div>
  </NuxtLink>
</template>
