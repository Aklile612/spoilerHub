<script setup lang="ts">
/**
 * Search results page — /search?q=movieName
 * Fetches search results from backend and displays in a grid.
 */
const route = useRoute();
const { results, loading, error, query, search } = useSearch();

// Extract query from URL
const searchQuery = computed(() => (route.query.q as string) || "");

// Fetch on mount and when query changes
watch(
  searchQuery,
  (q) => {
    if (q) search(q);
  },
  { immediate: true }
);

useHead({
  title: computed(
    () =>
      searchQuery.value
        ? `Search: "${searchQuery.value}" — SpoilerHub`
        : "Search — SpoilerHub"
  ),
});
</script>

<template>
  <div class="pt-20 mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
    <!-- Search header -->
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-gray-900">
        <template v-if="searchQuery">
          Results for "<span class="text-brand-dark">{{ searchQuery }}</span>"
        </template>
        <template v-else>Search Movies</template>
      </h1>
      <p v-if="results.length > 0" class="mt-1 text-sm text-gray-500">
        {{ results.length }} movie{{ results.length !== 1 ? "s" : "" }} found
      </p>
    </div>

    <!-- Error state -->
    <div
      v-if="error"
      class="rounded-lg border border-red-200 bg-red-50 p-6 text-center"
    >
      <p class="text-sm font-medium text-red-800">
        Search failed. Please try again.
      </p>
      <p class="mt-1 text-xs text-red-600">{{ error }}</p>
    </div>

    <!-- Results grid -->
    <MovieGrid :movies="results" :loading="loading" />

    <!-- No query state -->
    <div
      v-if="!searchQuery && !loading"
      class="flex flex-col items-center justify-center py-20 text-center"
    >
      <span class="text-5xl" aria-hidden="true">🔍</span>
      <p class="mt-4 text-lg font-medium text-gray-500">
        Search for any movie
      </p>
      <p class="mt-1 text-sm text-gray-400">
        Use the search bar above to find movie spoilers
      </p>
    </div>
  </div>
</template>
