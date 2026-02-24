<template>
  <div>
    <!-- Hero Section -->
    <section class="text-center space-y-6 py-12">
      <div class="space-y-4">
        <h1 class="text-5xl md:text-6xl font-bold bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
          SpoilerHub
        </h1>
        <p class="text-xl md:text-2xl text-gray-300 max-w-2xl mx-auto">
          Discover detailed spoiler explanations for your favorite movies powered by AI
        </p>
      </div>

      <!-- Search Section -->
      <div class="py-8">
        <MovieSearch
          @search="handleSearch"
          ref="searchComponent"
        />
      </div>
    </section>

    <!-- Loading State -->
    <Transition name="fade">
      <div
        v-if="loading"
        class="flex flex-col items-center justify-center py-20 space-y-4"
      >
        <svg class="spinner w-16 h-16 text-blue-500" fill="none" viewBox="0 0 24 24">
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
    </Transition>

    <!-- Movie Details -->
    <MovieDetails v-if="movie && !loading" :movie="movie" @back="handleBack" />

    <!-- Features Section (when no movie is shown) -->
    <section v-if="!movie && !loading" class="py-16 space-y-8">
      <h2 class="text-3xl font-bold text-center text-white mb-12">
        Features
      </h2>
      <div class="grid md:grid-cols-3 gap-6">
        <div class="card p-6 space-y-3">
          <div class="text-4xl">🎬</div>
          <h3 class="text-lg font-semibold text-white">Movie Details</h3>
          <p class="text-gray-400">
            Get comprehensive information including poster, rating, genres, and overview
          </p>
        </div>

        <div class="card p-6 space-y-3">
          <div class="text-4xl">⚠️</div>
          <h3 class="text-lg font-semibold text-white">Spoiler Explanations</h3>
          <p class="text-gray-400">
            AI-generated detailed explanations of plot twists, endings, and themes
          </p>
        </div>

        <div class="card p-6 space-y-3">
          <div class="text-4xl">⚡</div>
          <h3 class="text-lg font-semibold text-white">Instant Results</h3>
          <p class="text-gray-400">
            Search any movie and get instant results with comprehensive analysis
          </p>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
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
