<script setup lang="ts">
/**
 * HeroCarousel — Cinematic full-screen carousel with poster + info layout.
 * 4 spotlight movies. Auto-slides every 8 seconds.
 * Design: backdrop bg, poster on the left, movie info beside it.
 */

interface HeroMovie {
  title: string;
  overview: string;
  rating: number;
  backdrop: string;
  poster: string;
  year: string;
  runtime: string;
  genres: string[];
}

const heroMovies: HeroMovie[] = [
  {
    title: "Dune: Part Two",
    overview:
      "Follow the mythic journey of Paul Atreides as he unites with Chani and the Fremen while on a path of revenge against the conspirators who destroyed his family. Facing a choice between the love of his life and the fate of the known universe, Paul endeavors to prevent a terrible future only he can foresee.",
    rating: 8.1,
    backdrop: "https://image.tmdb.org/t/p/original/ylkdrn23p3gQcHx7ukIfuy2CkTE.jpg",
    poster: "https://image.tmdb.org/t/p/w500/1pdfLvkbY9ohJlCjQH2CZjjYVvJ.jpg",
    year: "2024",
    runtime: "2h 46m",
    genres: ["Sci-Fi", "Adventure"],
  },
  {
    title: "Blade Runner 2049",
    overview:
      "Thirty years after the events of the first film, a new blade runner, LAPD Officer K, unearths a long-buried secret that has the potential to plunge what's left of society into chaos. K's discovery leads him on a quest to find Rick Deckard, a former LAPD blade runner who has been missing for thirty years.",
    rating: 8.0,
    backdrop: "https://image.tmdb.org/t/p/original/sAtoMqDVhNDQBc3QJL3RF6hlhGq.jpg",
    poster: "https://image.tmdb.org/t/p/w500/gajva2L0rPYkEWjzgFlBXCAVBE5.jpg",
    year: "2017",
    runtime: "2h 44m",
    genres: ["Sci-Fi", "Drama"],
  },
  {
    title: "Whiplash",
    overview:
      "Under the direction of a ruthless instructor, a talented young drummer begins to pursue perfection at any cost, even his humanity. Andrew Neiman is an ambitious jazz student pushed to his absolute limit by Terence Fletcher, an instructor who will stop at nothing to realize his students' potential.",
    rating: 8.4,
    backdrop: "https://image.tmdb.org/t/p/original/1kuYEvLkX2nTkbfzN6X0w0xQFQU.jpg",
    poster: "https://image.tmdb.org/t/p/w500/7fn624j5lj3xTme2SgiLCeuedmO.jpg",
    year: "2014",
    runtime: "1h 47m",
    genres: ["Drama", "Music"],
  },
  {
    title: "Interstellar",
    overview:
      "When Earth becomes uninhabitable in the future, a farmer and ex-NASA pilot is tasked with piloting a spacecraft along with a team of researchers to find a new planet for humans. Traveling through a wormhole near Saturn, they venture into the unknown while confronting the limits of human endurance and the boundless nature of love.",
    rating: 8.7,
    backdrop: "https://image.tmdb.org/t/p/original/pbrkL804c8yAv3zBZR4QPEafpAR.jpg",
    poster: "https://image.tmdb.org/t/p/w500/gEU2QniE6E77NI6lCU6MxlNBvIx.jpg",
    year: "2014",
    runtime: "2h 49m",
    genres: ["Sci-Fi", "Adventure", "Drama"],
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
  autoplayTimer = setInterval(next, 8000);
}

onMounted(() => {
  resetAutoplay();
});

onUnmounted(() => {
  if (autoplayTimer) clearInterval(autoplayTimer);
});

function formatRating(rating: number): string {
  return rating.toFixed(1);
}
</script>

<template>
  <section class="relative h-[85vh] min-h-[550px] w-full overflow-hidden bg-black">
    
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

    <!-- Gradient overlays -->
    <div class="absolute inset-0 z-20 bg-gradient-to-t from-black via-black/10 to-black/30" />
    <div class="absolute inset-0 z-20 bg-gradient-to-r from-black/80 via-black/30 to-transparent" />

    <!-- Content: poster + info -->
    <div class="absolute inset-0 z-30 flex items-end pb-20 sm:items-center sm:pb-0">
      <div class="mx-auto flex w-full max-w-7xl items-end gap-6 px-6 sm:items-center sm:gap-10 lg:px-16">

        <!-- Poster card -->
        <div class="hidden shrink-0 sm:block">
          <div class="relative w-40 overflow-hidden rounded-xl shadow-2xl shadow-black/50 ring-1 ring-white/10 transition-transform duration-700 lg:w-48">
            <img
              :src="currentMovie.poster"
              :alt="currentMovie.title + ' poster'"
              class="h-full w-full object-cover"
              loading="eager"
            />
          </div>
        </div>

        <!-- Movie info -->
        <div class="flex-1 max-w-2xl">
          <!-- Tags row -->
          <div class="mb-3 flex flex-wrap items-center gap-2">
            <span class="rounded bg-brand px-2.5 py-0.5 text-[11px] font-bold uppercase tracking-wide text-black">
              Now Streaming
            </span>
            <span
              v-for="genre in currentMovie.genres"
              :key="genre"
              class="rounded border border-white/20 bg-white/10 px-2.5 py-0.5 text-[11px] font-medium text-white/90 backdrop-blur-sm"
            >
              {{ genre }}
            </span>
            <!-- Inline rating -->
            <span class="inline-flex items-center gap-1 text-[11px] font-bold text-brand">
              <svg class="h-3 w-3" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
              </svg>
              {{ formatRating(currentMovie.rating) }}
            </span>
          </div>

          <!-- Title -->
          <h2
            class="text-3xl font-black uppercase leading-none tracking-tight text-white sm:text-4xl lg:text-5xl"
            style="text-shadow: 0 4px 20px rgba(0,0,0,0.6)"
          >
            {{ currentMovie.title }}
          </h2>

          <!-- Meta: year · runtime -->
          <div class="mt-2.5 flex items-center gap-2 text-[13px] text-white/50">
            <span>{{ currentMovie.year }}</span>
            <span>·</span>
            <span>{{ currentMovie.runtime }}</span>
          </div>

          <!-- Overview -->
          <p class="mt-4 line-clamp-3 max-w-lg text-sm leading-relaxed text-white/70">
            {{ currentMovie.overview }}
          </p>

          <!-- Action buttons -->
          <div class="mt-6 flex items-center gap-3">
            <!-- Reveal Spoiler -->
            <NuxtLink
              :to="'/movie/' + encodeURIComponent(currentMovie.title)"
              class="inline-flex items-center gap-2 rounded-full border border-white/20 bg-white/10 px-5 py-2.5 text-sm font-semibold text-white backdrop-blur-sm transition-all duration-200 hover:bg-white/20 hover:border-white/30"
            >
              <!-- Eye icon -->
              <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              Reveal Spoiler
              <!-- <span class="text-base">⚠️</span> -->
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <!-- Navigation arrows -->
    <button
      @click="prev"
      class="absolute left-3 top-1/2 z-30 -translate-y-1/2 rounded-full bg-black/30 p-2.5 text-white/70 backdrop-blur-sm transition-all hover:bg-black/50 hover:text-white sm:left-5"
      aria-label="Previous movie"
    >
      <svg class="h-4 w-4 sm:h-5 sm:w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
        <path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
      </svg>
    </button>
    <button
      @click="next"
      class="absolute right-3 top-1/2 z-30 -translate-y-1/2 rounded-full bg-black/30 p-2.5 text-white/70 backdrop-blur-sm transition-all hover:bg-black/50 hover:text-white sm:right-5"
      aria-label="Next movie"
    >
      <svg class="h-4 w-4 sm:h-5 sm:w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2.5">
        <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
      </svg>
    </button>

    <!-- Dots -->
    <div class="absolute bottom-8 left-1/2 z-30 flex -translate-x-1/2 items-center gap-2">
      <button
        v-for="(movie, index) in heroMovies"
        :key="'dot-' + movie.title"
        @click="goTo(index)"
        class="group relative h-2 rounded-full transition-all duration-300 focus-visible:outline-none"
        :class="
          index === currentIndex
            ? 'w-7 bg-brand'
            : 'w-2 bg-white/30 hover:bg-white/60'
        "
        :aria-label="`Go to ${movie.title}`"
      >
        <span
          class="absolute -top-8 left-1/2 -translate-x-1/2 whitespace-nowrap rounded bg-black/80 px-2 py-1 text-[10px] font-medium text-white opacity-0 transition-opacity group-hover:opacity-100"
        >
          {{ movie.title }}
        </span>
      </button>
    </div>
  </section>
</template>
