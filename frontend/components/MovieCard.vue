<script setup lang="ts">
import type { Movie } from "~/types/movie";

/**
 * MovieCard — cinematic poster card with overlay info.
 * Full poster background, dark gradient, always-visible CTA.
 */
const props = defineProps<{
  movie: Movie;
}>();

/** Color-coded rating: green 7+, yellow 5–7, red <5 */
const ratingColor = computed(() => {
  if (props.movie.rating >= 7) return 'text-green-400';
  if (props.movie.rating >= 5) return 'text-yellow-400';
  return 'text-red-400';
});

/** Placeholder for missing posters */
const posterFallback =
  "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='300' height='450' fill='%231a1a2e'%3E%3Crect width='300' height='450'/%3E%3Ctext x='50%25' y='50%25' dominant-baseline='middle' text-anchor='middle' font-family='sans-serif' font-size='18' fill='%236b7280'%3ENo Poster%3C/text%3E%3C/svg%3E";
</script>

<template>
  <NuxtLink
    :to="{
      path: '/movie/' + encodeURIComponent(movie.title),
    }"
    class="group relative block aspect-[2/3] overflow-hidden rounded-2xl bg-gray-900 shadow-md transition-all duration-300 hover:shadow-xl hover:shadow-black/30 hover:-translate-y-1.5 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2"
  >
    <!-- Full-bleed poster -->
    <img
      :src="movie.poster || posterFallback"
      :alt="`${movie.title} poster`"
      class="absolute inset-0 h-full w-full object-cover transition-transform duration-500 group-hover:scale-110"
      loading="lazy"
    />

    <!-- Dark gradient overlay — stronger at bottom -->
    <div class="absolute inset-0 bg-gradient-to-t from-black via-black/30 to-transparent" />

    <!-- Hover glow overlay -->
    <div class="absolute inset-0 bg-brand/0 transition-colors duration-300 group-hover:bg-brand/5" />

    <!-- Rating badge — top left -->
    <div class="absolute left-2.5 top-2.5 z-10">
      <span :class="ratingColor" class="inline-flex items-center gap-1 rounded-lg bg-black/60 px-2 py-1 text-[11px] font-bold backdrop-blur-sm">
        <svg class="h-3 w-3" fill="currentColor" viewBox="0 0 20 20">
          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
        </svg>
        {{ movie.rating.toFixed(1) }}
      </span>
    </div>

    <!-- Bottom content overlay -->
    <div class="absolute inset-x-0 bottom-0 z-10 flex flex-col justify-end p-3.5 sm:p-4">
      <!-- Genre pills -->
      <div v-if="movie.genres?.length" class="mb-2 flex flex-wrap gap-1">
        <span
          v-for="genre in movie.genres.slice(0, 2)"
          :key="genre"
          class="rounded-full bg-white/10 px-2 py-0.5 text-[10px] font-medium text-white/70 backdrop-blur-sm"
        >
          {{ genre }}
        </span>
        <span class="ml-0.5 self-center text-[10px] text-white/40">{{ movie.year }}</span>
      </div>

      <!-- Title -->
      <h3 class="line-clamp-2 text-sm font-bold leading-snug text-white sm:text-base">
        {{ movie.title }}
      </h3>

      <!-- See Spoiler CTA — hidden by default, visible on hover -->
      <div class="mt-2.5 flex items-center gap-1.5 opacity-0 translate-y-2 transition-all duration-300 group-hover:opacity-100 group-hover:translate-y-0">
        <span class="inline-flex items-center gap-1.5 rounded-full bg-brand px-3 py-1.5 text-[11px] font-bold text-gray-900 shadow-sm shadow-brand/30 transition-all duration-300 group-hover:shadow-md group-hover:shadow-brand/40 group-hover:scale-105">
          <svg class="h-3 w-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
            <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          See Spoiler
        </span>
      </div>
    </div>
  </NuxtLink>
</template>
