<script setup lang="ts">
/**
 * HeroCarousel — Full-screen cinematic carousel featuring 4 spotlight movies.
 * Auto-slides every 6 seconds. Users can click dots or arrows to navigate.
 * Each slide shows backdrop image, title, overview, rating, and a "See Spoiler" CTA.
 */

interface HeroMovie {
  title: string;
  overview: string;
  rating: number;
  backdrop: string;
  year: string;
  genres: string[];
}

const heroMovies: HeroMovie[] = [
  {
    title: "Marty Supreme",
    overview:
      "Marty Mauser, a young man with a dream no one respects, goes to hell and back in pursuit of greatness. Set in the gritty underworld of competitive table tennis, Timothée Chalamet delivers a career-defining performance as a hustler who risks everything to prove himself on the biggest stage.",
    rating: 7.6,
    backdrop: "https://image.tmdb.org/t/p/original/qKWDHofjMHPSEOTLaixkC9ZmjTT.jpg",
    year: "2025",
    genres: ["Drama"],
  },
  {
    title: "Whiplash",
    overview:
      "Under the direction of a ruthless instructor, a talented young drummer begins to pursue perfection at any cost, even his humanity. Andrew Neiman is an ambitious jazz student at the finest music conservatory in the country, pushed to his absolute limit by Terence Fletcher, an instructor who will stop at nothing to realize his students' potential.",
    rating: 8.4,
    backdrop: "https://image.tmdb.org/t/p/original/1kuYEvLkX2nTkbfzN6X0w0xQFQU.jpg",
    year: "2014",
    genres: ["Drama", "Music"],
  },
  {
    title: "Sinners",
    overview:
      "Trying to leave their troubled lives behind, twin brothers return to their hometown to start again, only to discover that an even greater evil is waiting to welcome them back. Ryan Coogler reunites with Michael B. Jordan in this visceral supernatural thriller set in the Deep South, blending horror with searing social commentary.",
    rating: 7.5,
    backdrop: "https://image.tmdb.org/t/p/original/nAxGnGHOsfzufThz20zgmRwKur3.jpg",
    year: "2025",
    genres: ["Horror", "Action", "Thriller"],
  },
  {
    title: "The Brutalist",
    overview:
      "When an innovative modern architect flees post-war Europe, he is given the opportunity to rebuild his legacy in America. Adrien Brody stars in Brady Corbet's sweeping epic about ambition, art, and the immigrant experience, set during the dawn of modern America. A wealthy patron forever changes his life.",
    rating: 7.0,
    backdrop: "https://image.tmdb.org/t/p/original/hmZnqijPaaACjenDkrbWcCmcADI.jpg",
    year: "2024",
    genres: ["Drama"],
  },
];

const currentIndex = ref(0);
const isTransitioning = ref(false);
let autoplayTimer: ReturnType<typeof setInterval> | null = null;

const currentMovie = computed(() => heroMovies[currentIndex.value]);

function goTo(index: number) {
  if (index === currentIndex.value || isTransitioning.value) return;
  isTransitioning.value = true;
  currentIndex.value = index;
  resetAutoplay();
  setTimeout(() => {
    isTransitioning.value = false;
  }, 700);
}

function next() {
  goTo((currentIndex.value + 1) % heroMovies.length);
}

function prev() {
  goTo((currentIndex.value - 1 + heroMovies.length) % heroMovies.length);
}

function resetAutoplay() {
  if (autoplayTimer) clearInterval(autoplayTimer);
  autoplayTimer = setInterval(next, 10000);
}

onMounted(() => {
  resetAutoplay();
});

onUnmounted(() => {
  if (autoplayTimer) clearInterval(autoplayTimer);
});

/** Format rating to one decimal */
function formatRating(rating: number): string {
  return rating.toFixed(1);
}

function ratingColor(rating: number): string {
  if (rating >= 7) return "bg-green-500";
  if (rating >= 5) return "bg-yellow-500";
  return "bg-red-500";
}
</script>

<template>
  <section class="relative h-[75vh] min-h-[500px] w-full overflow-hidden bg-black">
    <!-- Backdrop images — all stacked, only active one is visible -->
    <div
      v-for="(movie, index) in heroMovies"
      :key="movie.title"
      class="absolute inset-0 transition-opacity duration-700 ease-in-out"
      :class="index === currentIndex ? 'opacity-100 z-10' : 'opacity-0 z-0'"
    >
      <img
        :src="movie.backdrop"
        :alt="movie.title"
        class="h-full w-full object-cover"
        loading="eager"
      />
    </div>

    <!-- Gradient overlays for readability -->
    <div class="absolute inset-0 z-20 bg-gradient-to-t from-black/90 via-black/40 to-black/20" />
    <div class="absolute inset-0 z-20 bg-gradient-to-r from-black/70 via-transparent to-transparent" />

    <!-- Movie info — bottom left -->
    <div class="absolute bottom-0 left-0 z-30 w-full px-6 pb-16 pt-8 sm:px-10 lg:px-16 lg:pb-20">
      <div class="max-w-2xl">
        <!-- Genres -->
        <div class="mb-3 flex flex-wrap items-center gap-2">
          <span
            v-for="genre in currentMovie.genres"
            :key="genre"
            class="rounded-full border border-white/30 px-3 py-0.5 text-[11px] font-medium uppercase tracking-wider text-white/80"
          >
            {{ genre }}
          </span>
          <span class="text-xs text-white/50">{{ currentMovie.year }}</span>
        </div>

        <!-- Title -->
        <h2
          class="text-3xl font-extrabold leading-tight text-white sm:text-4xl lg:text-5xl"
          style="text-shadow: 0 2px 12px rgba(0,0,0,0.5)"
        >
          {{ currentMovie.title }}
        </h2>

        <!-- Rating -->
        <div class="mt-3 flex items-center gap-3">
          <span
            :class="ratingColor(currentMovie.rating)"
            class="inline-flex items-center gap-1 rounded-md px-2 py-0.5 text-xs font-bold text-white shadow-sm"
          >
            <svg class="h-3 w-3" fill="currentColor" viewBox="0 0 20 20" aria-hidden="true">
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
            </svg>
            {{ formatRating(currentMovie.rating) }}
          </span>
          <span class="text-sm text-white/60">TMDB</span>
        </div>

        <!-- Overview (truncated) -->
        <p class="mt-4 line-clamp-3 max-w-xl text-sm leading-relaxed text-white/80 sm:text-base">
          {{ currentMovie.overview }}
        </p>

        <!-- See Spoiler button -->
        <NuxtLink
          :to="'/movie/' + encodeURIComponent(currentMovie.title)"
          class="mt-6 inline-flex items-center gap-2.5 rounded-lg bg-brand px-6 py-3 text-sm font-bold text-gray-900 shadow-lg transition-all duration-200 hover:bg-brand-dark hover:shadow-xl hover:scale-105 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-black"
        >
          <!-- Play icon -->
          <svg class="h-5 w-5" viewBox="0 0 24 24" fill="currentColor">
            <path d="M8 5v14l11-7z" />
          </svg>
          See Spoiler
        </NuxtLink>
      </div>
    </div>

    <!-- Navigation arrows -->
    <button
      @click="prev"
      class="absolute left-3 top-1/2 z-30 -translate-y-1/2 rounded-full bg-black/40 p-2 text-white/80 backdrop-blur-sm transition-all hover:bg-black/60 hover:text-white sm:left-5 sm:p-3"
      aria-label="Previous movie"
    >
      <svg class="h-5 w-5 sm:h-6 sm:w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
        <path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
      </svg>
    </button>
    <button
      @click="next"
      class="absolute right-3 top-1/2 z-30 -translate-y-1/2 rounded-full bg-black/40 p-2 text-white/80 backdrop-blur-sm transition-all hover:bg-black/60 hover:text-white sm:right-5 sm:p-3"
      aria-label="Next movie"
    >
      <svg class="h-5 w-5 sm:h-6 sm:w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
        <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
      </svg>
    </button>

    <!-- Dots / Slide tracker -->
    <div class="absolute bottom-6 left-1/2 z-30 flex -translate-x-1/2 items-center gap-2">
      <button
        v-for="(movie, index) in heroMovies"
        :key="'dot-' + movie.title"
        @click="goTo(index)"
        class="group relative h-2.5 rounded-full transition-all duration-300 focus-visible:outline-none"
        :class="
          index === currentIndex
            ? 'w-8 bg-brand'
            : 'w-2.5 bg-white/40 hover:bg-white/70'
        "
        :aria-label="`Go to ${movie.title}`"
      >
        <!-- Tooltip on hover -->
        <span
          class="absolute -top-8 left-1/2 -translate-x-1/2 whitespace-nowrap rounded bg-black/80 px-2 py-1 text-[10px] font-medium text-white opacity-0 transition-opacity group-hover:opacity-100"
        >
          {{ movie.title }}
        </span>
      </button>
    </div>

  </section>
</template>
