<template>
  <Transition name="fade">
    <div
      v-if="movie"
      class="w-full max-w-4xl mx-auto space-y-6 animate-fade-in"
    >
      <!-- Backdrop Image -->
      <div class="relative h-64 md:h-96 rounded-lg overflow-hidden group">
        <img
          v-if="movie.backdrop"
          :src="movie.backdrop"
          :alt="movie.title"
          class="w-full h-full object-cover brightness-50 group-hover:brightness-75 transition-all duration-300"
        />
        <div
          v-else
          class="w-full h-full bg-gradient-to-br from-dark-700 to-dark-900 flex items-center justify-center"
        >
          <svg
            class="w-16 h-16 text-dark-600"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
            />
          </svg>
        </div>

        <!-- Rating Badge (Overlay) -->
        <div class="absolute top-4 right-4 bg-dark-900 bg-opacity-90 rounded-lg p-3">
          <div class="flex items-center gap-2">
            <svg
              class="w-5 h-5 text-yellow-400"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
              />
            </svg>
            <span class="text-lg font-bold text-white">{{
              movie.rating.toFixed(1)
            }}</span>
          </div>
        </div>
      </div>

      <!-- Movie Info Card -->
      <div class="card p-6 space-y-4">
        <!-- Title and Year -->
        <div>
          <h1 class="text-3xl md:text-4xl font-bold text-white">
            {{ movie.title }}
          </h1>
          <p class="text-gray-400 text-lg mt-2">{{ movie.year }}</p>
        </div>

        <!-- Genres -->
        <div v-if="movie.genres.length > 0" class="flex flex-wrap gap-2">
          <span
            v-for="genre in movie.genres"
            :key="genre"
            class="badge badge-primary"
          >
            {{ genre }}
          </span>
        </div>

        <!-- Overview -->
        <div class="border-t border-dark-700 pt-4">
          <h2 class="text-xl font-semibold text-gray-100 mb-3">Overview</h2>
          <p class="text-gray-300 leading-relaxed">
            {{ movie.overview }}
          </p>
        </div>
      </div>

      <!-- Poster Image -->
      <div class="flex justify-center">
        <div class="w-full max-w-xs">
          <img
            v-if="movie.poster"
            :src="movie.poster"
            :alt="movie.title"
            class="w-full rounded-lg shadow-2xl border border-dark-700 hover:shadow-3xl transition-shadow duration-300"
          />
          <div
            v-else
            class="w-full aspect-[2/3] bg-gradient-to-br from-dark-700 to-dark-900 rounded-lg shadow-2xl border border-dark-700 flex items-center justify-center"
          >
            <svg
              class="w-12 h-12 text-dark-600"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
              />
            </svg>
          </div>
        </div>
      </div>

      <!-- Spoiler Section -->
      <div class="card overflow-hidden">
        <button
          @click="toggleSpoiler"
          class="w-full px-6 py-4 bg-dark-700 hover:bg-dark-600 transition-colors duration-200 flex items-center justify-between"
        >
          <span class="text-lg font-semibold text-white">
            ⚠️ Reveal Spoiler Explanation
          </span>
          <svg
            class="w-6 h-6 text-gray-400 transition-transform duration-300"
            :class="{ 'rotate-180': showSpoiler }"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 14l-7 7m0 0l-7-7m7 7V3"
            />
          </svg>
        </button>

        <!-- Spoiler Content (Hidden by Default) -->
        <Transition name="slide-fade">
          <div v-if="showSpoiler" class="p-6 border-t border-dark-700">
            <div
              class="prose prose-invert max-w-none text-gray-300 space-y-4"
              v-html="formattedSpoiler"
            />
          </div>
        </Transition>
      </div>

      <!-- Back Button -->
      <div class="text-center">
        <button
          @click="emit('back')"
          class="btn btn-secondary inline-flex items-center gap-2"
        >
          <svg
            class="w-5 h-5"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M10 19l-7-7m0 0l7-7m-7 7h18"
            />
          </svg>
          Search Another Movie
        </button>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import type { Movie } from "~/composables/useMovieAPI";

interface Props {
  movie: Movie | null;
}

const props = defineProps<Props>();
const emit = defineEmits<{
  back: [];
}>();

const showSpoiler = ref(false);

const toggleSpoiler = () => {
  showSpoiler.value = !showSpoiler.value;
};

// Format spoiler text to HTML (convert markdown-like formatting)
const formattedSpoiler = computed(() => {
  if (!props.movie?.spoiler) return "";

  let html = props.movie.spoiler;

  // Convert markdown-style headings to HTML
  html = html.replace(/^### (.*?)$/gm, "<h3 class='text-xl font-semibold text-gray-100 mt-4 mb-2'>$1</h3>");
  html = html.replace(/^## (.*?)$/gm, "<h2 class='text-2xl font-bold text-white mt-6 mb-3'>$1</h2>");
  html = html.replace(/^# (.*?)$/gm, "<h1 class='text-3xl font-bold text-white mt-8 mb-4'>$1</h1>");

  // Convert bold text
  html = html.replace(/\*\*(.*?)\*\*/g, "<strong class='font-semibold text-gray-100'>$1</strong>");

  // Convert italic text
  html = html.replace(/\*(.*?)\*/g, "<em class='italic'>$1</em>");

  // Convert line breaks to paragraphs
  html = html.replace(/\n\n/g, "</p><p>");
  html = `<p>${html}</p>`;

  return html;
});
</script>
