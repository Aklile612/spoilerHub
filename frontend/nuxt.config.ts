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
  // In production, set NUXT_PUBLIC_API_BASE_URL env var in Vercel
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
            "Discover detailed spoiler explanations for your favorite movies",
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
          href: "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'%3E%3Crect width='100' height='100' rx='20' fill='%23111'/%3E%3Crect x='4' y='4' width='92' height='92' rx='16' fill='none' stroke='%23F5C518' stroke-width='6'/%3E%3Ctext x='50' y='66' font-family='Inter,system-ui,sans-serif' font-size='48' font-weight='900' font-style='italic' fill='%23F5C518' text-anchor='middle'%3ESP%3C/text%3E%3C/svg%3E",
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
