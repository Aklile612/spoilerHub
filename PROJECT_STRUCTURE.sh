#!/usr/bin/env bash

# SpoilerHub Application Tree
# Shows the complete project structure

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║         SpoilerHub - Complete Project Structure               ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

tree_structure='
spoiler_api/
│
├── 📄 README.md                    # Main project documentation
├── 📄 SETUP.md                     # Setup and deployment guide
├── 📄 Makefile                     # Build automation
├── 🔧 setup.sh                     # Automated setup script
│
├── backend/                        # Go Backend API
│   ├── 📄 README.md               # Backend documentation
│   ├── 📄 .env.example            # Example environment
│   ├── 📄 .gitignore              # Git ignore rules
│   ├── go.mod                     # Go module dependencies
│   │
│   ├── cmd/
│   │   └── server/
│   │       └── main.go            # Application entry point
│   │           └── Features:
│   │               • Server initialization
│   │               • Route configuration
│   │               • CORS middleware
│   │               • Dependency injection
│   │
│   └── internal/
│       ├── config/
│       │   └── config.go          # Configuration loader
│       │       └── Loads from environment variables
│       │
│       ├── models/
│       │   └── movie.go           # Data structures
│       │       ├── MovieResponse
│       │       ├── TMDBMovie
│       │       ├── GeminiRequest/Response
│       │       └── ErrorResponse
│       │
│       ├── services/
│       │   ├── tmdb_service.go    # TMDB API client
│       │   │   └── Methods:
│       │   │       • SearchMovie()
│       │   │       • GetGenres()
│       │   │       • FormatPosterURL()
│       │   │       • ExtractYear()
│       │   │       • ExtractGenreNames()
│       │   │
│       │   └── gemini_service.go  # Gemini API client with caching
│       │       └── Features:
│       │           • GenerateSpoiler()
│       │           • In-memory cache (thread-safe)
│       │           • Cache management methods
│       │
│       ├── handlers/
│       │   └── movie_handler.go   # HTTP request handlers
│       │       └── Endpoints:
│       │           • GET /api/movie
│       │           • GET /health
│       │
│       └── routes/
│           └── routes.go          # Route definitions
│               └── API routes setup
│
├── frontend/                       # Nuxt 3 Frontend
│   ├── 📄 README.md               # Frontend documentation
│   ├── 📄 .env.example            # Example environment
│   ├── 📄 .gitignore              # Git ignore rules
│   ├── package.json               # Node dependencies
│   ├── nuxt.config.ts             # Nuxt configuration
│   ├── tailwind.config.ts         # Tailwind CSS config
│   ├── tsconfig.json              # TypeScript config
│   │
│   ├── app.vue                    # Root component (Homepage)
│   │   └── Features:
│   │       • Hero section
│   │       • Search integration
│   │       • Movie details display
│   │       • Features showcase
│   │
│   ├── assets/
│   │   └── css/
│   │       ├── tailwind.css       # Tailwind imports
│   │       └── global.css         # Global styles & animations
│   │
│   ├── components/
│   │   ├── MovieSearch.vue        # Search component
│   │   │   └── Features:
│   │   │       • Input validation
│   │   │       • Loading state
│   │   │       • Error display
│   │   │       • Form submission
│   │   │
│   │   └── MovieDetails.vue       # Movie details display
│   │       └── Features:
│   │           • Poster & backdrop images
│   │           • Rating badge
│   │           • Genre tags
│   │           • Overview section
│   │           • Spoiler toggle (hidden by default)
│   │           • Markdown formatting support
│   │           • Back navigation
│   │
│   ├── composables/
│   │   └── useMovieAPI.ts         # API integration composable
│   │       └── Exports:
│   │           • fetchMovie()
│   │           • clearMovie()
│   │           • Reactive: movie, loading, error
│   │
│   ├── layouts/
│   │   └── default.vue            # Main layout
│   │       └── Features:
│   │           • Header with branding
│   │           • Content slot
│   │           • Footer with attribution
│   │           • Responsive design
│   │
│   └── pages/
│       └── index.vue              # Home/search page
│           └── Auto-routed by Nuxt
│
└── Configuration Files
    ├── .env files (create from .env.example)
    │   Backend needs: TMDB_API_KEY, GEMINI_API_KEY, PORT
    │   Frontend needs: NUXT_PUBLIC_API_BASE_URL
    │
    └── git files
        └── .gitignore in both backend/ and frontend/

════════════════════════════════════════════════════════════════

Backend Stack:
  • Go 1.21+
  • Gin Web Framework
  • HTTP Client for external APIs
  • Mutex for thread-safe caching
  • JSON marshaling/unmarshaling

Frontend Stack:
  • Nuxt 3
  • Vue 3 Composition API
  • TypeScript
  • TailwindCSS
  • Server-Side Rendering

External Integrations:
  • TMDB API v3 (movie database)
  • Google Gemini API (AI spoiler generation)

════════════════════════════════════════════════════════════════
'

echo "$tree_structure"

echo ""
echo "🎬 Total Files Created:"
echo "  Backend: 8 Go source files + 2 config files"
echo "  Frontend: 9 Vue/TS files + 4 config files"
echo "  Documentation: 4 markdown files + 2 scripts"
echo ""
echo "✅ Project is ready for development!"
echo ""
