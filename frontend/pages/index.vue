<template>
  <div class="w-full max-w-4xl mx-auto p-4">
    <div
      v-if="loading"
      class="flex flex-col items-center justify-center py-20 space-y-4"
    >
      <svg class="spinner w-12 h-12 text-blue-500" fill="none" viewBox="0 0 24 24">
        <circle
          class="opacity-25"
          cx="12"
          cy="12"
          r="10"
          stroke="currentColor"
          stroke-width="4"
        />
        <path
          class="opacity-75"
          fill="currentColor"
          d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
        />
      </svg>
      <p class="text-gray-400 text-lg">
        Generating spoiler explanation with AI...
      </p>
    </div>

    <MovieDetails v-else-if="movie" :movie="movie" @back="handleBack" />

    <MovieSearch v-else @search="handleSearch" ref="searchComponent" />
  </div>
</template>

<script setup lang="ts">
import { useMovieAPI } from "~/composables/useMovieAPI";

const { movie, loading, error, fetchMovie, clearMovie } = useMovieAPI();
const searchComponent = ref(null);

const handleSearch = async (title: string) => {
  // Set loading state on search component
  if (searchComponent.value) {
    searchComponent.value.setLoading(true);
  }

  // Fetch movie data
  const result = await fetchMovie(title);

  // Set error state on search component
  if (searchComponent.value) {
    searchComponent.value.setLoading(false);
    searchComponent.value.setError(error.value);
  }
};

const handleBack = () => {
  clearMovie();
};
</script>
