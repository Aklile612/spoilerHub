<script setup lang="ts">
/**
 * AppHeader — top navigation bar with logo + search.
 * Sticky, clean white background, subtle bottom border.
 */
</script>

<template>
  <header
    class="sticky top-0 z-50 border-b border-gray-100 bg-white/95 backdrop-blur-sm"
  >
    <div
      class="mx-auto flex max-w-7xl items-center justify-between gap-4 px-4 py-3 sm:px-6 lg:px-8"
    >
      <!-- Logo -->
      <NuxtLink
        to="/"
        class="flex shrink-0 items-center gap-2 transition-opacity hover:opacity-80"
        aria-label="SpoilerHub Home"
      >
        <span class="text-2xl" aria-hidden="true">🍿</span>
        <span class="text-xl font-extrabold tracking-tight text-gray-900">
          Spoiler<span class="text-brand">Hub</span>
        </span>
      </NuxtLink>

      <!-- Search bar -->
      <SearchBar />
    </div>
  </header>
</template>
