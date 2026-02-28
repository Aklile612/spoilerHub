<script setup lang="ts">
/**
 * AppHeader — slim glass-morphism navigation bar with artistic SP monogram + search.
 */
</script>

<template>
  <header class="fixed top-0 left-0 right-0 z-[60] bg-black/25 backdrop-blur-lg">
    <div
      class="mx-auto flex max-w-7xl items-center justify-between gap-3 px-4 py-2 sm:px-6 lg:px-8"
    >
      <!-- Logo — artistic SP monogram -->
      <NuxtLink
        to="/"
        class="group flex shrink-0 items-center gap-2.5 transition-all hover:opacity-90"
        aria-label="SpoilerHub Home"
      >
        <!-- SP monogram mark -->
        <div class="relative flex h-8 w-8 items-center justify-center">
          <!-- Glowing ring -->
          <div
            class="absolute inset-0 rounded-lg bg-gradient-to-br from-brand via-amber-400 to-orange-500 opacity-80 blur-[1px] transition-all group-hover:opacity-100 group-hover:blur-[2px]"
          />
          <!-- Inner background -->
          <div class="absolute inset-[1.5px] rounded-[6px] bg-black/80" />
          <!-- SP text -->
          <span
            class="relative z-10 text-[13px] font-black italic tracking-tighter text-brand"
            style="font-family: 'Inter', system-ui, sans-serif"
          >
            SP
          </span>
        </div>

        <!-- Wordmark -->
        <div class="flex items-baseline gap-0">
          <span
            class="text-[15px] font-light uppercase tracking-[0.2em] text-white/90"
          >
            Spoiler
          </span>
          <span
            class="text-[15px] font-black uppercase tracking-[0.2em] text-brand"
          >
            Hub
          </span>
        </div>
      </NuxtLink>

      <!-- Search bar -->
      <SearchBar />
    </div>
  </header>
</template>
