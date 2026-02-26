import type { Config } from "tailwindcss";

export default {
  content: [
    "./components/**/*.{vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./composables/**/*.ts",
    "./app.vue",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Inter", "system-ui", "sans-serif"],
      },
      colors: {
        brand: {
          DEFAULT: "#F5C518",
          dark: "#E0B400",
        },
        spoiler: {
          DEFAULT: "#C62828",
          light: "#EF5350",
          dark: "#8E0000",
          bg: "#FFF5F5",
          border: "#FFCDD2",
        },
        surface: {
          DEFAULT: "#F4F4F4",
          raised: "#FFFFFF",
          sunken: "#EBEBEB",
        },
      },
      maxWidth: {
        "8xl": "88rem",
      },
      borderRadius: {
        "2xl": "16px",
      },
      animation: {
        "fade-in": "fadeIn 0.5s ease-out forwards",
        "slide-up": "slideUp 0.4s ease-out forwards",
        "slide-in-right": "slideInRight 0.5s ease-out forwards",
      },
      keyframes: {
        fadeIn: {
          "0%": { opacity: "0" },
          "100%": { opacity: "1" },
        },
        slideUp: {
          "0%": { opacity: "0", transform: "translateY(16px)" },
          "100%": { opacity: "1", transform: "translateY(0)" },
        },
        slideInRight: {
          "0%": { opacity: "0", transform: "translateX(24px)" },
          "100%": { opacity: "1", transform: "translateX(0)" },
        },
      },
    },
  },
  plugins: [],
} satisfies Config;
