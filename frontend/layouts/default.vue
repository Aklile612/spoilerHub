<template>
  <div
    class="min-h-screen bg-gradient-to-br from-dark-900 via-dark-800 to-dark-900"
  >
    <!-- Header -->
    <header class="border-b border-dark-700 bg-dark-800 bg-opacity-50 backdrop-blur">
      <div class="max-w-6xl mx-auto px-4 py-6">
        <NuxtLink to="/" class="flex items-center gap-3 hover:opacity-80 transition-opacity">
          <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
            <span class="text-white font-bold text-lg">🍿</span>
          </div>
          <h1 class="text-2xl font-bold text-white">SpoilerHub</h1>
        </NuxtLink>
      </div>
    </header>

    <!-- Main Content -->
    <main class="max-w-6xl mx-auto px-4 py-12">
      <slot />
    </main>

    <!-- Footer -->
    <footer class="border-t border-dark-700 bg-dark-800 bg-opacity-50 backdrop-blur mt-16">
      <div class="max-w-6xl mx-auto px-4 py-8">
        <div class="text-center text-gray-400 text-sm space-y-2">
          <p>© 2024 SpoilerHub. All rights reserved.</p>
          <p>
            Powered by
            <a
              href="https://www.themoviedb.org/"
              target="_blank"
              rel="noopener noreferrer"
              class="text-blue-400 hover:text-blue-300"
            >
              TMDB
            </a>
            and
            <a
              href="https://ai.google.dev/"
              target="_blank"
              rel="noopener noreferrer"
              class="text-blue-400 hover:text-blue-300"
            >
              Google Gemini AI
            </a>
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<style scoped>
/* Layout transitions */
.page-enter-active,
.page-leave-active {
  transition: opacity 0.3s ease;
}

.page-enter-from,
.page-leave-to {
  opacity: 0;
}
</style>
