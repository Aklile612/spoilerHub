<script setup lang="ts">
/**
 * Default layout — AppHeader + main content + footer.
 * White background, centered max-width container.
 */
</script>

<template>
  <div class="flex min-h-screen flex-col bg-white font-sans text-gray-900">
    <!-- Global header -->
    <AppHeader />

    <!-- Page content -->
    <main class="flex-1">
      <slot />
    </main>

    <!-- Footer -->
    <footer class="border-t border-gray-100 bg-gray-50 py-6">
      <div class="mx-auto max-w-7xl px-4 text-center sm:px-6 lg:px-8">
        <p class="text-sm text-gray-400">
          © {{ new Date().getFullYear() }} SpoilerHub — Built with Nuxt 3 &amp;
          Go
        </p>
        <p class="mt-1 text-xs text-gray-300">
          Movie data from TMDB · 
        </p>
        <p class="mt-1 text-sm text-gray-300">
          Created by <a href="https://aklileansa.vercel.app/" class="text-gray-400 hover:text-gray-500">Aklile Ansa</a>
        </p>
      </div>
    </footer>
  </div>
</template>
