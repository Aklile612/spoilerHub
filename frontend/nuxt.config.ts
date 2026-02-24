// https://nuxt.com/docs/guide/concepts/rendering#hybrid-rendering
export default defineNuxtConfig({
  // Enable SSR
  ssr: true,

  // Modules
  modules: ["@nuxtjs/tailwindcss"],

  // Tailwind configuration
  tailwindcss: {
    cssPath: "~/assets/css/tailwind.css",
    configPath: "tailwind.config.ts",
  },

  // Runtime configuration
  runtimeConfig: {
    public: {
      apiBaseUrl: "http://localhost:8080",
    },
  },

  // App configuration
  app: {
    head: {
      title: "SpoilerHub - Movie Spoiler Explanations",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          name: "description",
          content: "Discover detailed spoiler explanations for your favorite movies powered by AI",
        },
        {
          name: "theme-color",
          content: "#111827",
        },
        {
          name: "og:title",
          content: "SpoilerHub - Movie Spoiler Explanations",
        },
        {
          name: "og:description",
          content: "Discover detailed spoiler explanations for your favorite movies powered by AI",
        },
        {
          name: "og:type",
          content: "website",
        },
      ],
      link: [
        {
          rel: "icon",
          type: "image/svg+xml",
          href: "data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='75' font-size='75'>🍿</text></svg>",
        },
      ],
    },
  },

  // Global CSS
  css: ["~/assets/css/global.css"],

  // Nitro (server-side) configuration
  nitro: {
    prerender: {
      crawlLinks: false,
    },
  },

  // Build configuration
  build: {
    transpile: [],
  },

  // Development server configuration
  devServer: {
    port: 3000,
  },

  // Component auto-import configuration
  components: true,

  // Composable auto-import configuration
  imports: {
    autoImport: true,
  },

  // TypeScript configuration
  typescript: {
    strict: true,
  },
});
