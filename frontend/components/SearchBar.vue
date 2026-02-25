<script setup lang="ts">
/**
 * SearchBar — debounced search input in the header.
 * Emits navigation to /search?q=term on submit.
 */
const router = useRouter();
const route = useRoute();

const searchQuery = ref((route.query.q as string) || "");
let debounceTimer: ReturnType<typeof setTimeout> | null = null;

function handleSubmit() {
  const q = searchQuery.value.trim();
  if (q) {
    router.push({ path: "/search", query: { q } });
  }
}

function handleInput() {
  // Clear previous timer
  if (debounceTimer) clearTimeout(debounceTimer);

  debounceTimer = setTimeout(() => {
    const q = searchQuery.value.trim();
    if (q.length >= 2) {
      router.push({ path: "/search", query: { q } });
    }
  }, 500);
}
</script>

<template>
  <form
    @submit.prevent="handleSubmit"
    class="relative flex w-full max-w-md items-center"
    role="search"
  >
    <!-- Search icon -->
    <svg
      class="pointer-events-none absolute left-3 h-4 w-4 text-gray-400"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      viewBox="0 0 24 24"
      aria-hidden="true"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z"
      />
    </svg>

    <input
      v-model="searchQuery"
      @input="handleInput"
      type="search"
      placeholder="Search movies…"
      aria-label="Search movies"
      class="w-full rounded-lg border border-gray-200 bg-gray-50 py-2 pl-10 pr-4 text-sm text-gray-900 placeholder-gray-400 transition-colors focus:border-brand focus:bg-white focus:outline-none focus:ring-1 focus:ring-brand"
    />
  </form>
</template>
