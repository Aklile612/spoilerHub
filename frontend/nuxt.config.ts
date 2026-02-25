// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  // Enable SSR for fast initial load
  ssr: true,

  // Modules
  modules: ["@nuxtjs/tailwindcss"],

  // Tailwind configuration
  tailwindcss: {
    cssPath: "~/assets/css/tailwind.css",
    configPath: "tailwind.config.ts",
  },

  // Runtime configuration — backend API base URL
  runtimeConfig: {
    public: {
      apiBaseUrl: "http://localhost:8080",
    },
  },

  // App head configuration
  app: {
    head: {
      title: "SpoilerHub — Movie Spoilers & Explanations",
      htmlAttrs: { lang: "en" },
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          name: "description",
          content:
            "Discover detailed AI-generated spoiler explanations for your favorite movies",
        },
        { name: "theme-color", content: "#ffffff" },
      ],
      link: [
        // Google Fonts — Inter for clean professional typography
        {
          rel: "preconnect",
          href: "https://fonts.googleapis.com",
        },
        {
          rel: "preconnect",
          href: "https://fonts.gstatic.com",
          crossorigin: "",
        },
        {
          rel: "stylesheet",
          href: "https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap",
        },
        {
          rel: "icon",
          type: "image/svg+xml",
          href: "data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='75' font-size='75'>🍿</text></svg>",
        },
      ],
    },
    // Page transition animation
    pageTransition: { name: "page", mode: "out-in" },
  },

  // Global CSS
  css: ["~/assets/css/main.css"],

  // Nitro config
  nitro: {
    prerender: {
      crawlLinks: false,
    },
  },

  // Dev server
  devServer: {
    port: 3000,
  },

  // Auto-import components
  components: true,

  // TypeScript strict mode
  typescript: {
    strict: true,
  },

  // Compatibility date
  compatibilityDate: "2024-11-01",
});
