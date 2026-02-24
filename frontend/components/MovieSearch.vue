<template>
  <div class="w-full max-w-md mx-auto">
    <!-- Search Form -->
    <form @submit.prevent="handleSearch" class="space-y-4">
      <div class="relative">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search for a movie..."
          class="input w-full pl-12"
          :disabled="isLoading"
          @keydown.enter="handleSearch"
        />
        <svg
          class="absolute left-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-500"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          />
        </svg>
      </div>

      <!-- Search Button -->
      <button
        type="submit"
        class="btn btn-primary w-full flex items-center justify-center gap-2"
        :disabled="isLoading"
      >
        <span v-if="!isLoading">Search</span>
        <span v-else class="flex items-center gap-2">
          <svg class="spinner w-4 h-4" fill="none" viewBox="0 0 24 24">
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
          Searching...
        </span>
      </button>
    </form>

    <!-- Error Message -->
    <Transition name="fade">
      <div
        v-if="error"
        class="mt-4 p-4 bg-red-900 bg-opacity-20 border border-red-600 rounded-lg text-red-300"
      >
        <p class="text-sm">{{ error }}</p>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const emit = defineEmits<{
  search: [title: string];
}>();

const searchQuery = ref("");
const isLoading = ref(false);
const error = ref<string | null>(null);

const handleSearch = async () => {
  // Clear previous error
  error.value = null;

  // Validate input
  if (!searchQuery.value.trim()) {
    error.value = "Please enter a movie title";
    return;
  }

  isLoading.value = true;
  emit("search", searchQuery.value);

  // Reset loading state after a delay (parent will handle actual loading)
  setTimeout(() => {
    isLoading.value = false;
  }, 500);
};

// Allow parent to set loading state
defineExpose({
  setLoading: (value: boolean) => {
    isLoading.value = value;
  },
  setError: (message: string | null) => {
    error.value = message;
  },
});
</script>
