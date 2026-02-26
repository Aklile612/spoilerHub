<script setup lang="ts">
/**
 * Premium Movie Detail Page — cinematic, editorial spoiler experience.
 * Route: /movie/[title]
 *
 * Sections:
 *  1. Compact cinematic banner (backdrop + poster + info)
 *  2. Spoiler warning gate
 *  3. Story breakdown (collapsible accordion with timeline)
 *  4. Character fates grid
 *  5. Key moments horizontal timeline
 *  6. Theory & interpretation
 *  7. Ratings & reactions
 *  8. Related movies
 */
import type { Movie } from "~/types/movie";

const route = useRoute();
const config = useRuntimeConfig();
const baseUrl = config.public.apiBaseUrl as string;

// ── Movie title from route ──
const movieTitle = computed(() =>
  decodeURIComponent(route.params.title as string)
);

// ── Fetch movie with spoiler ──
const {
  data: movie,
  pending: loading,
  error,
} = useFetch<Movie>(`${baseUrl}/api/movie`, {
  query: { title: movieTitle.value },
  key: `movie-${movieTitle.value}`,
});

// ── Parse spoiler text into structured sections ──
const spoilerRef = computed(() => movie.value?.spoiler);
const {
  storySections,
  characterFates,
  keyMoments,
  interpretations,
  hasStructuredContent,
} = useSpoilerParser(spoilerRef);

// ── Spoiler reveal state ──
const isSpoilerRevealed = ref(false);
function revealSpoilers() {
  isSpoilerRevealed.value = true;
  nextTick(() => {
    document
      .getElementById("story-breakdown")
      ?.scrollIntoView({ behavior: "smooth", block: "start" });
  });
}

// ── Accordion state ──
const openSections = ref<Set<string>>(new Set());
function toggleSection(id: string) {
  if (openSections.value.has(id)) {
    openSections.value.delete(id);
  } else {
    openSections.value.add(id);
  }
  openSections.value = new Set(openSections.value);
}

// ── Interpretation tab ──
const activeInterpretation = ref("ai");

// ── Related movies ──
const { data: relatedData } = useFetch<{ movies: Movie[] }>(
  `${baseUrl}/api/trending`,
  {
    key: "related-movies",
    default: () => ({ movies: [] }),
  }
);
const relatedMovies = computed(() =>
  (relatedData.value?.movies ?? [])
    .filter((m) => m.title !== movie.value?.title)
    .slice(0, 6)
);

// ── Markdown renderer ──
function renderMarkdown(text: string): string {
  return text
    .replace(/^### (.+)$/gm, "<h3>$1</h3>")
    .replace(/^## (.+)$/gm, "<h2>$1</h2>")
    .replace(/\*\*(.+?)\*\*/g, "<strong>$1</strong>")
    .replace(/\*(.+?)\*/g, "<em>$1</em>")
    .replace(/^- (.+)$/gm, "<li>$1</li>")
    .replace(/((?:<li>.*<\/li>\n?)+)/g, "<ul>$1</ul>")
    .replace(/\n\n/g, "</p><p>")
    .replace(/\n/g, "<br>")
    .replace(/^(.+)/, "<p>$1</p>");
}

// ── Share ──
async function shareMovie() {
  if (navigator.share) {
    await navigator.share({
      title: `${movie.value?.title} — SpoilerHub`,
      text: `Check out the spoiler breakdown for ${movie.value?.title}`,
      url: window.location.href,
    });
  } else {
    await navigator.clipboard.writeText(window.location.href);
  }
}

// ── Bookmark ──
const isBookmarked = ref(false);
function toggleBookmark() {
  isBookmarked.value = !isBookmarked.value;
}

// ── Page head ──
useHead({
  title: computed(() =>
    movie.value
      ? `${movie.value.title} (${movie.value.year}) — SpoilerHub`
      : "Loading… — SpoilerHub"
  ),
});

// ── Fallbacks ──
const backdropFallback =
  "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='1280' height='720' fill='%23111827'%3E%3Crect width='1280' height='720'/%3E%3C/svg%3E";
const posterFallback =
  "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='300' height='450' fill='%23374151'%3E%3Crect width='300' height='450'/%3E%3Ctext x='50%25' y='50%25' dominant-baseline='middle' text-anchor='middle' font-family='sans-serif' font-size='18' fill='%239ca3af'%3ENo Poster%3C/text%3E%3C/svg%3E";

// ── Helpers ──
function statusColor(status: string) {
  switch (status) {
    case "ALIVE": return "bg-emerald-500";
    case "DEAD": return "bg-red-600";
    default: return "bg-gray-400";
  }
}
function statusBg(status: string) {
  switch (status) {
    case "ALIVE": return "bg-emerald-50 border-emerald-200";
    case "DEAD": return "bg-red-50 border-red-200";
    default: return "bg-gray-50 border-gray-200";
  }
}
function interpretationIcon(title: string) {
  switch (title) {
    case "Symbolism": return "🔮";
    case "Hidden Clues": return "🔍";
    case "Fan Theories": return "💭";
    case "Unanswered Questions": return "❓";
    default: return "📝";
  }
}
</script>

<template>
  <div class="min-h-screen bg-surface">
    <!-- ═══ LOADING ═══ -->
    <div v-if="loading" class="animate-pulse">
      <div class="relative h-72 w-full overflow-hidden bg-gray-900 sm:h-80">
        <div class="skeleton-dark h-full w-full" />
      </div>
      <div class="mx-auto -mt-20 max-w-5xl px-4 sm:px-6 lg:px-8">
        <div class="flex gap-6">
          <div class="skeleton h-64 w-44 shrink-0 rounded-2xl" />
          <div class="flex-1 space-y-3 pt-24">
            <div class="skeleton h-8 w-3/4 rounded-lg" />
            <div class="skeleton h-4 w-1/3 rounded" />
            <div class="flex gap-2">
              <div class="skeleton h-7 w-16 rounded-full" />
              <div class="skeleton h-7 w-20 rounded-full" />
              <div class="skeleton h-7 w-16 rounded-full" />
            </div>
            <div class="skeleton h-20 w-full rounded-lg" />
          </div>
        </div>
      </div>
      <div class="mx-auto mt-10 max-w-5xl space-y-4 px-4 sm:px-6 lg:px-8">
        <div class="skeleton h-32 w-full rounded-2xl" />
        <div class="skeleton h-48 w-full rounded-2xl" />
      </div>
    </div>

    <!-- ═══ ERROR ═══ -->
    <div
      v-else-if="error"
      class="flex min-h-[70vh] flex-col items-center justify-center px-4 text-center"
    >
      <div class="rounded-2xl border border-gray-200 bg-white p-10 shadow-sm">
        <span class="text-6xl">😕</span>
        <h2 class="mt-5 text-2xl font-bold text-gray-900">Movie Not Found</h2>
        <p class="mt-2 text-gray-500">
          We couldn't find details for "{{ movieTitle }}".
        </p>
        <NuxtLink
          to="/"
          class="mt-6 inline-flex items-center gap-2 rounded-full bg-brand px-6 py-2.5 text-sm font-bold text-gray-900 shadow-md transition-all hover:bg-brand-dark hover:shadow-lg"
        >
          <svg class="h-4 w-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
          </svg>
          Back to Home
        </NuxtLink>
      </div>
    </div>

    <!-- ═══ MOVIE CONTENT ═══ -->
    <div v-else-if="movie">
      <!-- ━━ 1. COMPACT CINEMATIC BANNER ━━ -->
      <div class="relative overflow-hidden">
        <!-- Backdrop -->
        <div class="absolute inset-0 h-72 sm:h-80">
          <img
            :src="movie.backdrop || backdropFallback"
            :alt="`${movie.title} backdrop`"
            class="h-full w-full object-cover object-top"
          />
          <div class="absolute inset-0 bg-gradient-to-b from-black/70 via-black/50 to-surface" />
          <div class="absolute inset-0 bg-gradient-to-r from-black/40 to-transparent" />
        </div>

        <!-- Nav -->
        <div class="relative z-10 mx-auto max-w-5xl px-4 pt-20 sm:px-6 lg:px-8">
          <div class="flex items-center justify-between">
            <NuxtLink
              to="/"
              class="inline-flex items-center gap-1.5 rounded-full bg-white/15 px-4 py-1.5 text-xs font-medium text-white backdrop-blur-md transition-all hover:bg-white/25"
            >
              <svg class="h-3.5 w-3.5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
              </svg>
              Back to Home
            </NuxtLink>

            <!-- <div class="flex items-center gap-2">
              <button
                @click="toggleBookmark"
                class="flex h-8 w-8 items-center justify-center rounded-full bg-white/15 backdrop-blur-md transition-all hover:bg-white/25"
              >
                <svg
                  class="h-4 w-4 transition-colors"
                  :class="isBookmarked ? 'fill-brand text-brand' : 'text-white'"
                  :fill="isBookmarked ? 'currentColor' : 'none'"
                  stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"
                >
                  <path stroke-linecap="round" stroke-linejoin="round" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
                </svg>
              </button>
              <button
                @click="shareMovie"
                class="flex h-8 w-8 items-center justify-center rounded-full bg-white/15 backdrop-blur-md transition-all hover:bg-white/25"
              >
                <svg class="h-4 w-4 text-white" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                </svg>
              </button>
            </div> -->
          </div>
        </div>

        <!-- Poster + Info -->
        <div class="relative z-10 mx-auto max-w-5xl px-4 pb-8 pt-6 sm:px-6 sm:pt-10 lg:px-8">
          <div class="flex flex-col gap-5 sm:flex-row sm:items-end">
            <div class="shrink-0">
              <img
                :src="movie.poster || posterFallback"
                :alt="`${movie.title} poster`"
                class="h-auto w-36 rounded-2xl shadow-2xl ring-4 ring-white/10 sm:w-44"
              />
            </div>
            <div class="flex-1 pb-1">
              <h1 class="text-2xl font-extrabold leading-tight text-white drop-shadow-lg sm:text-3xl lg:text-4xl">
                {{ movie.title }}
              </h1>
              <div class="mt-2.5 flex flex-wrap items-center gap-3 text-sm">
                <span class="text-white/70">{{ movie.year }}</span>
                <span class="text-white/30">•</span>
                <span class="gold-glow inline-flex items-center gap-1 rounded-lg bg-brand px-2.5 py-1 text-xs font-bold text-gray-900">
                  <svg class="h-3.5 w-3.5" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                  {{ movie.rating.toFixed(1) }}
                </span>
              </div>
              <div v-if="movie.genres?.length" class="mt-3 flex flex-wrap gap-2">
                <span
                  v-for="genre in movie.genres"
                  :key="genre"
                  class="rounded-full bg-white/15 px-3 py-1 text-xs font-medium text-white/90 backdrop-blur-sm"
                >
                  {{ genre }}
                </span>
              </div>
              <p class="mt-4 max-w-2xl text-sm leading-relaxed text-white/60">
                {{ movie.overview || "No overview available." }}
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- ━━ MAIN CONTENT ━━ -->
      <div class="mx-auto max-w-5xl px-4 pb-16 sm:px-6 lg:px-8">

        <!-- ━━ 2. SPOILER WARNING GATE ━━ -->
        <section v-if="movie.spoiler && !isSpoilerRevealed" class="mt-8 animate-fade-in">
          <div class="red-glow-pulse overflow-hidden rounded-2xl border-2 border-spoiler/20 bg-white shadow-lg">
            <div class="border-b border-spoiler/10 bg-spoiler/5 px-6 py-3">
              <span class="text-xs font-bold uppercase tracking-widest text-spoiler">Content Warning</span>
            </div>
            <div class="flex flex-col items-center px-6 py-10 text-center">
              <div class="flex h-16 w-16 items-center justify-center rounded-full bg-spoiler/10">
                <span class="text-4xl">⚠️</span>
              </div>
              <h3 class="mt-5 text-xl font-bold text-gray-900">
                This page contains major spoilers
              </h3>
              <p class="mt-2 max-w-md text-sm text-gray-500">
                Full plot breakdown, major twists, character fates, and ending details are revealed below.
                Continue at your own risk.
              </p>
              <button
                @click="revealSpoilers"
                class="mt-6 inline-flex items-center gap-2 rounded-full bg-spoiler px-8 py-3 text-sm font-bold text-white shadow-lg transition-all hover:bg-spoiler-dark hover:shadow-xl active:scale-95"
              >
                <svg class="h-4 w-4" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                Reveal Full Spoilers
              </button>
            </div>
          </div>
        </section>

        <!-- No spoiler -->
        <section v-if="!movie.spoiler" class="mt-8">
          <div class="flex flex-col items-center rounded-2xl border border-gray-200 bg-white px-6 py-12 text-center shadow-sm">
            <span class="text-5xl">🎬</span>
            <h3 class="mt-4 text-lg font-bold text-gray-900">No Spoiler Available</h3>
            <p class="mt-2 text-sm text-gray-500">
              AI-generated spoiler analysis is not yet available for this movie.
            </p>
            <NuxtLink
              to="/"
              class="mt-6 inline-flex items-center gap-2 rounded-full bg-brand px-6 py-2.5 text-sm font-bold text-gray-900 transition-all hover:bg-brand-dark"
            >
              Explore Other Movies
            </NuxtLink>
          </div>
        </section>

        <!-- ═══ REVEALED SPOILER CONTENT ═══ -->
        <Transition name="spoiler-reveal">
          <div v-if="isSpoilerRevealed && movie.spoiler">

            <!-- ━━ 3. STORY BREAKDOWN ━━ -->
            <section id="story-breakdown" class="mt-10">
              <div class="mb-6 flex items-start gap-3">
                <div class="mt-1 h-8 w-1.5 rounded-full bg-brand" />
                <div>
                  <h2 class="text-xl font-black tracking-tight text-gray-900">Full Plot Breakdown</h2>
                  <p class="mt-0.5 text-xs text-gray-400">Expand each section to reveal the story</p>
                </div>
              </div>

              <!-- Timeline accordions -->
              <div v-if="hasStructuredContent" class="timeline-line space-y-3 pl-8">
                <div
                  v-for="(section, index) in storySections"
                  :key="section.id"
                  class="animate-fade-in"
                  :style="{ animationDelay: `${index * 100}ms` }"
                >
                  <button
                    @click="toggleSection(section.id)"
                    class="group flex w-full items-center gap-3 rounded-2xl border bg-white px-5 py-4 text-left shadow-sm transition-all hover:shadow-md"
                    :class="openSections.has(section.id) ? 'border-brand/30 shadow-brand/5' : 'border-gray-100 hover:border-gray-200'"
                  >
                    <div
                      class="absolute -left-[1px] flex h-8 w-8 items-center justify-center rounded-full border-2 bg-white text-sm"
                      :class="openSections.has(section.id) ? 'border-brand' : 'border-gray-200'"
                    >
                      {{ section.icon }}
                    </div>
                    <div class="flex-1">
                      <h3 class="text-sm font-bold text-gray-900">{{ section.title }}</h3>
                    </div>
                    <svg
                      class="h-4 w-4 shrink-0 text-gray-400 transition-transform duration-300"
                      :class="{ 'rotate-180 text-brand': openSections.has(section.id) }"
                      fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"
                    >
                      <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
                    </svg>
                  </button>
                  <Transition name="accordion">
                    <div v-if="openSections.has(section.id)" class="ml-3 mt-1 border-l-2 border-brand/20 pl-5">
                      <div class="rounded-xl bg-white px-5 py-4 text-sm leading-relaxed text-gray-600 shadow-sm">
                        <div class="detail-spoiler-content" v-html="renderMarkdown(section.content)" />
                      </div>
                    </div>
                  </Transition>
                </div>
              </div>

              <!-- Fallback: raw spoiler for old cached data -->
              <div v-else class="rounded-2xl border border-gray-100 bg-white px-6 py-5 shadow-sm">
                <div class="detail-spoiler-content prose prose-sm max-w-none" v-html="renderMarkdown(movie.spoiler)" />
              </div>
            </section>

            <!-- ━━ 4. CHARACTER FATES ━━ -->
            <section v-if="characterFates.length > 0" id="character-fates" class="mt-12">
              <div class="mb-6 flex items-start gap-3">
                <div class="mt-1 h-8 w-1.5 rounded-full bg-brand" />
                <div>
                  <h2 class="text-xl font-black tracking-tight text-gray-900">Character Fates</h2>
                  <p class="mt-0.5 text-xs text-gray-400">How each character's story ends</p>
                </div>
              </div>
              <div class="grid gap-3 sm:grid-cols-2 lg:grid-cols-3">
                <div
                  v-for="(char, idx) in characterFates"
                  :key="char.name"
                  class="group animate-fade-in cursor-default overflow-hidden rounded-2xl border bg-white shadow-sm transition-all duration-300 hover:-translate-y-0.5 hover:shadow-md"
                  :class="statusBg(char.status)"
                  :style="{ animationDelay: `${idx * 80}ms` }"
                >
                  <div class="p-4">
                    <div class="flex items-start justify-between">
                      <div class="min-w-0 flex-1">
                        <h4 class="truncate text-sm font-bold text-gray-900">{{ char.name }}</h4>
                        <p v-if="char.actor" class="truncate text-xs text-gray-400">{{ char.actor }}</p>
                      </div>
                      <span
                        class="status-pulse ml-2 shrink-0 rounded-full px-2.5 py-0.5 text-[10px] font-bold uppercase tracking-wider text-white"
                        :class="statusColor(char.status)"
                      >
                        {{ char.status }}
                      </span>
                    </div>
                    <p class="mt-2.5 text-xs leading-relaxed text-gray-500">{{ char.summary }}</p>
                  </div>
                </div>
              </div>
            </section>

            <!-- ━━ 5. KEY MOMENTS ━━ -->
            <section v-if="keyMoments.length > 0" id="key-moments" class="mt-12">
              <div class="mb-6 flex items-start gap-3">
                <div class="mt-1 h-8 w-1.5 rounded-full bg-brand" />
                <div>
                  <h2 class="text-xl font-black tracking-tight text-gray-900">Key Moments</h2>
                  <p class="mt-0.5 text-xs text-gray-400">Pivotal scenes that define the film</p>
                </div>
              </div>
              <div class="flex gap-4 overflow-x-auto pb-4 scrollbar-hide">
                <div
                  v-for="(moment, idx) in keyMoments"
                  :key="idx"
                  class="group w-60 shrink-0 animate-slide-in-right cursor-default overflow-hidden rounded-2xl border border-gray-100 bg-white shadow-sm transition-all duration-300 hover:-translate-y-0.5 hover:border-brand/30 hover:shadow-lg"
                  :style="{ animationDelay: `${idx * 100}ms` }"
                >
                  <div class="h-1 w-full bg-gradient-to-r from-brand via-amber-400 to-orange-500" />
                  <div class="p-4">
                    <div class="mb-3 flex items-center gap-2">
                      <span class="flex h-6 w-6 items-center justify-center rounded-full bg-brand/10 text-xs font-bold text-brand-dark">
                        {{ idx + 1 }}
                      </span>
                      <span class="text-[10px] font-bold uppercase tracking-widest text-gray-300">Scene</span>
                    </div>
                    <h4 class="text-sm font-bold leading-snug text-gray-900">{{ moment.title }}</h4>
                    <p class="mt-2 text-xs leading-relaxed text-gray-500">{{ moment.description }}</p>
                  </div>
                </div>
              </div>
            </section>

            <!-- ━━ 6. THEORY & INTERPRETATION ━━ -->
            <section v-if="interpretations.length > 0" id="interpretation" class="mt-12">
              <!-- <div class="mb-6 flex items-start gap-3">
                <div class="mt-1 h-8 w-1.5 rounded-full bg-brand" />
                <div>
                  <h2 class="text-xl font-black tracking-tight text-gray-900">What It Really Means</h2>
                  <p class="mt-0.5 text-xs text-gray-400">Deeper analysis and interpretations</p>
                </div>
              </div> -->

              <!-- Tabs -->
              <!-- <div class="mb-5 flex gap-2">
                <button
                  @click="activeInterpretation = 'ai'"
                  class="rounded-full px-4 py-1.5 text-xs font-semibold transition-all"
                  :class="activeInterpretation === 'ai' ? 'bg-gray-900 text-white shadow-md' : 'bg-gray-100 text-gray-500 hover:bg-gray-200'"
                >
                  🤖 AI Analysis
                </button>
                <button
                  @click="activeInterpretation = 'community'"
                  class="rounded-full px-4 py-1.5 text-xs font-semibold transition-all"
                  :class="activeInterpretation === 'community' ? 'bg-gray-900 text-white shadow-md' : 'bg-gray-100 text-gray-500 hover:bg-gray-200'"
                >
                  💬 Community Theories
                </button>
              </div> -->

              <!-- <div v-if="activeInterpretation === 'ai'" class="space-y-4">
                <div
                  v-for="(interp, idx) in interpretations"
                  :key="interp.title"
                  class="animate-fade-in overflow-hidden rounded-2xl border border-gray-100 bg-white shadow-sm"
                  :style="{ animationDelay: `${idx * 80}ms` }"
                >
                  <div class="flex items-start gap-4 p-5">
                    <div class="mt-1 h-full w-1 shrink-0 rounded-full bg-brand/30" />
                    <div class="flex-1">
                      <div class="flex items-center gap-2">
                        <span class="text-lg">{{ interpretationIcon(interp.title) }}</span>
                        <h3 class="text-sm font-bold text-gray-900">{{ interp.title }}</h3>
                      </div>
                      <div class="detail-spoiler-content mt-2 text-sm leading-relaxed text-gray-600" v-html="renderMarkdown(interp.content)" />
                    </div>
                  </div>
                </div>
              </div>

              <div v-else class="flex flex-col items-center rounded-2xl border border-gray-100 bg-white px-6 py-12 text-center shadow-sm">
                <span class="text-4xl">💬</span>
                <p class="mt-3 text-sm font-medium text-gray-400">Community theories coming soon</p>
                <p class="mt-1 text-xs text-gray-300">User-submitted theories and discussions will appear here</p>
              </div> -->
            </section>

            <!-- ━━ 7. RATINGS & REACTIONS ━━ -->
            <section id="ratings" class="mt-12">
              <div class="mb-6 flex items-start gap-3">
                <div class="mt-1 h-8 w-1.5 rounded-full bg-brand" />
                <div>
                  <h2 class="text-xl font-black tracking-tight text-gray-900">Ratings & Reactions</h2>
                  <p class="mt-0.5 text-xs text-gray-400">How audiences responded</p>
                </div>
              </div>
              <div class="grid gap-4 sm:grid-cols-2">
                <!-- TMDB Rating -->
                <div class="overflow-hidden rounded-2xl border border-gray-100 bg-white p-6 shadow-sm">
                  <div class="flex items-center gap-3">
                    <div class="flex h-14 w-14 items-center justify-center rounded-2xl bg-brand/10">
                      <svg class="h-7 w-7 text-brand" fill="currentColor" viewBox="0 0 20 20">
                        <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                      </svg>
                    </div>
                    <div>
                      <p class="text-xs font-semibold text-gray-400">TMDB Rating</p>
                      <p class="text-3xl font-black text-gray-900">
                        {{ movie.rating.toFixed(1) }}
                        <span class="text-sm font-medium text-gray-400">/ 10</span>
                      </p>
                    </div>
                  </div>
                  <div class="mt-4">
                    <div class="h-2 w-full overflow-hidden rounded-full bg-gray-100">
                      <div
                        class="progress-bar-animated h-full rounded-full"
                        :class="movie.rating >= 7 ? 'bg-emerald-500' : movie.rating >= 5 ? 'bg-amber-500' : 'bg-red-500'"
                        :style="{ '--progress-width': `${(movie.rating / 10) * 100}%` }"
                      />
                    </div>
                    <div class="mt-1.5 flex justify-between text-[10px] text-gray-300">
                      <span>0</span>
                      <span>5</span>
                      <span>10</span>
                    </div>
                  </div>
                </div>

                <!-- AI Analysis card -->
                <!-- <div class="overflow-hidden rounded-2xl border border-gray-100 bg-white p-6 shadow-sm">
                  <div class="flex items-center gap-3">
                    <div class="flex h-14 w-14 items-center justify-center rounded-2xl bg-spoiler/10">
                      <span class="text-2xl">🧠</span>
                    </div>
                    <div>
                      <p class="text-xs font-semibold text-gray-400">AI Spoiler Analysis</p>
                      <p class="text-sm font-bold text-gray-900">Powered by Gemini</p>
                    </div>
                  </div>
                  <div class="mt-4 space-y-2">
                    <div class="flex items-center justify-between">
                      <span class="text-xs text-gray-500">Story Sections</span>
                      <span class="text-xs font-bold text-gray-900">{{ storySections.length }}</span>
                    </div>
                    <div class="flex items-center justify-between">
                      <span class="text-xs text-gray-500">Characters Tracked</span>
                      <span class="text-xs font-bold text-gray-900">{{ characterFates.length }}</span>
                    </div>
                    <div class="flex items-center justify-between">
                      <span class="text-xs text-gray-500">Key Moments</span>
                      <span class="text-xs font-bold text-gray-900">{{ keyMoments.length }}</span>
                    </div>
                    <div class="flex items-center justify-between">
                      <span class="text-xs text-gray-500">Interpretations</span>
                      <span class="text-xs font-bold text-gray-900">{{ interpretations.length }}</span>
                    </div>
                  </div>
                </div> -->
              </div>
            </section>
          </div>
        </Transition>

        <!-- ━━ 8. RELATED MOVIES ━━ -->
        <section v-if="relatedMovies.length > 0" id="related" class="mt-14">
          <div class="mb-6 flex items-start gap-3">
            <div class="mt-1 h-8 w-1.5 rounded-full bg-brand" />
            <div>
              <h2 class="text-xl font-black tracking-tight text-gray-900">Related Movies</h2>
              <p class="mt-0.5 text-xs text-gray-400">More movies you might want to spoil</p>
            </div>
          </div>
          <div class="grid grid-cols-2 gap-3 sm:grid-cols-3 lg:grid-cols-6">
            <NuxtLink
              v-for="(related, idx) in relatedMovies"
              :key="related.title"
              :to="'/movie/' + encodeURIComponent(related.title)"
              class="group animate-fade-in overflow-hidden rounded-2xl border border-gray-100 bg-white shadow-sm transition-all duration-300 hover:-translate-y-1 hover:shadow-lg"
              :style="{ animationDelay: `${idx * 60}ms` }"
            >
              <div class="relative aspect-[2/3] overflow-hidden bg-gray-100">
                <img
                  :src="related.poster || posterFallback"
                  :alt="related.title"
                  class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
                  loading="lazy"
                />
                <div class="absolute right-2 top-2">
                  <span class="inline-flex items-center gap-0.5 rounded-md bg-black/60 px-1.5 py-0.5 text-[10px] font-bold text-brand backdrop-blur-sm">
                    <svg class="h-2.5 w-2.5" fill="currentColor" viewBox="0 0 20 20">
                      <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                    </svg>
                    {{ related.rating.toFixed(1) }}
                  </span>
                </div>
              </div>
              <div class="p-2.5">
                <h4 class="line-clamp-1 text-xs font-bold text-gray-900 group-hover:text-brand-dark">{{ related.title }}</h4>
                <p class="mt-0.5 text-[10px] text-gray-400">{{ related.year }}</p>
              </div>
            </NuxtLink>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>
