# SpoilerHub - Production-Ready Full-Stack Movie Spoiler Application

## 📋 Overview

SpoilerHub is a modern full-stack web application that allows users to search for movies and receive:

- Movie poster and backdrop images
- Title, release year, and rating
- Genre information
- Movie overview
- AI-generated detailed spoiler explanations
- Ending breakdown analysis
- Thematic analysis

The application integrates with **TMDB (The Movie Database)** for movie information and **Google Gemini AI** for intelligent spoiler generation.

---

## 🏗️ Architecture

### Backend (Go + Gin)
- RESTful API with clean architecture
- Service-oriented design pattern
- In-memory caching for spoiler results (thread-safe with mutex)
- Middleware for CORS handling
- Structured error handling

### Frontend (Nuxt 3 + TypeScript + TailwindCSS)
- Server-side rendering enabled
- Composition API with custom composables
- Responsive dark-themed UI
- Client-side state management
- Loading states and error handling

### External APIs
1. **TMDB API** - Movie database
2. **Google Gemini API** - AI spoiler generation

---

## 📁 Project Structure

```
spoiler_api/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go              # Application entry point
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go             # Configuration management
│   │   ├── handlers/
│   │   │   └── movie_handler.go      # HTTP request handlers
│   │   ├── models/
│   │   │   └── movie.go              # Data structures
│   │   ├── services/
│   │   │   ├── tmdb_service.go       # TMDB API client
│   │   │   └── gemini_service.go     # Gemini API client with caching
│   │   └── routes/
│   │       └── routes.go             # Route configuration
│   ├── go.mod                        # Go dependencies
│   └── .env.example                  # Example environment file
│
├── frontend/
│   ├── app.vue                       # Root component (homepage)
│   ├── nuxt.config.ts                # Nuxt configuration
│   ├── tailwind.config.ts            # Tailwind CSS configuration
│   ├── package.json                  # Node dependencies
│   ├── assets/
│   │   └── css/
│   │       ├── tailwind.css          # Tailwind imports
│   │       └── global.css            # Global styles
│   ├── components/
│   │   ├── MovieSearch.vue           # Search component
│   │   └── MovieDetails.vue          # Movie details component
│   ├── composables/
│   │   └── useMovieAPI.ts            # API interaction composable
│   ├── layouts/
│   │   └── default.vue               # Default layout
│   ├── pages/
│   │   └── index.vue                 # Home/search page
│   └── .env.example                  # Example environment file
│
└── README.md                          # This file
```

---

## 🚀 Quick Start

### Prerequisites

- **Go 1.21+** (Backend)
- **Node.js 18+** and **npm** or **yarn** (Frontend)
- **TMDB API Key** (Get from https://www.themoviedb.org/settings/api)
- **Google Gemini API Key** (Get from https://ai.google.dev/)

### Backend Setup

1. **Navigate to backend directory:**
   ```bash
   cd backend
   ```

2. **Create environment file:**
   ```bash
   cp .env.example .env
   ```

3. **Edit `.env` and add your API keys:**
   ```env
   PORT=8080
   ENVIRONMENT=development
   TMDB_API_KEY=your_tmdb_api_key_here
   GEMINI_API_KEY=your_gemini_api_key_here
   ```

4. **Download dependencies:**
   ```bash
   go mod download
   ```

5. **Run the backend server:**
   ```bash
   go run cmd/server/main.go
   ```

   Server will start at `http://localhost:8080`

### Frontend Setup

1. **Navigate to frontend directory:**
   ```bash
   cd frontend
   ```

2. **Create environment file:**
   ```bash
   cp .env.example .env
   ```

3. **Edit `.env` (optional - default is http://localhost:8080):**
   ```env
   NUXT_PUBLIC_API_BASE_URL=http://localhost:8080
   ```

4. **Install dependencies:**
   ```bash
   npm install
   # or
   yarn install
   ```

5. **Run the development server:**
   ```bash
   npm run dev
   # or
   yarn dev
   ```

   Frontend will be available at `http://localhost:3000`

---

## 📡 API Endpoints

### GET `/api/movie`

Searches for a movie and returns detailed information with AI-generated spoilers.

**Query Parameters:**
- `title` (required) - Movie title to search for

**Example Request:**
```bash
curl "http://localhost:8080/api/movie?title=Inception"
```

**Example Response:**
```json
{
  "title": "Inception",
  "year": "2010",
  "poster": "https://image.tmdb.org/t/p/w500/...",
  "backdrop": "https://image.tmdb.org/t/p/w1280/...",
  "rating": 8.8,
  "genres": ["Action", "Science Fiction", "Thriller"],
  "overview": "Cobb, a skilled thief who steals corporate secrets...",
  "spoiler": "⚠️ SPOILER WARNING\n\nMovie Overview\n...\n\nFull Spoiler Breakdown\n...\n\nEnding Explained\n...\n\nThemes & Meaning\n..."
}
```

### GET `/health`

Health check endpoint.

**Example Request:**
```bash
curl "http://localhost:8080/health"
```

**Example Response:**
```json
{
  "status": "healthy",
  "cache_size": 5
}
```

---

## 🎨 Frontend Features

### Search Page
- Clean search interface with autocomplete
- Real-time input validation
- Loading spinner during API calls
- Error message display

### Movie Details Page
- Backdrop image display
- Movie poster
- Rating badge with star icon
- Genre tags
- Overview section
- Hidden spoiler section (click to reveal)
- Smooth animations and transitions
- Dark theme with modern design
- Responsive layout (mobile, tablet, desktop)

### UI Components
- **MovieSearch.vue** - Search input and submission
- **MovieDetails.vue** - Movie information and spoiler display
- **Default Layout** - Header and footer with branding

### Composables
- **useMovieAPI.ts** - API integration, state management, error handling

---

## 🔒 Security Features

1. **API Keys Protection**
   - All API keys stored in environment variables
   - Never exposed in frontend code
   - All external API calls routed through backend

2. **CORS Middleware**
   - Proper CORS headers configured
   - Prevents unauthorized cross-origin requests

3. **Input Validation**
   - Query parameter validation on backend
   - XSS protection through Vue's built-in escaping

4. **Error Handling**
   - Graceful error handling throughout application
   - User-friendly error messages
   - No sensitive information in error responses

---

## 📊 Caching Strategy

The backend implements **in-memory caching** for spoiler generations:

- **Cache Key**: `{movie_title}_{release_year}`
- **Thread-Safe**: Protected with `sync.RWMutex`
- **Benefits**:
  - Reduced API calls to Gemini
  - Faster response times for repeated searches
  - Lower operational costs

**Cache Methods:**
- `GenerateSpoiler()` - Checks cache before API call
- `ClearCache()` - Clears all cached items
- `GetCacheSize()` - Returns number of cached entries

---

## 🛠️ Development

### Backend Development

**Project Structure:**
- `cmd/` - Application entry points
- `internal/` - Private application code
- `models/` - Data structures
- `services/` - Business logic and API clients
- `handlers/` - HTTP request handlers
- `routes/` - Route definitions

**Adding a New Endpoint:**

1. Create handler in `handlers/movie_handler.go`
2. Create service in `services/` if needed
3. Add route in `routes/routes.go`
4. Test the endpoint

### Frontend Development

**Project Structure:**
- `components/` - Reusable Vue components
- `composables/` - Reusable composition API logic
- `pages/` - Page components (auto-routed by Nuxt)
- `layouts/` - Layout templates
- `assets/` - Static assets and stylesheets

**Adding a New Page:**

1. Create `.vue` file in `pages/` directory
2. Nuxt automatically handles routing
3. Use composables for data fetching

---

## 📦 Dependencies

### Backend
- `github.com/gin-gonic/gin v1.9.1` - HTTP web framework

### Frontend
- `nuxt` - Vue 3 framework with SSR
- `@nuxtjs/tailwindcss` - Tailwind CSS integration
- `vue` - JavaScript framework
- `typescript` - Type safety

---

## 🧪 Testing the Application

### Test the API Manually

```bash
# Test health check
curl http://localhost:8080/health

# Test movie search
curl "http://localhost:8080/api/movie?title=The%20Shawshank%20Redemption"

# Test with special characters
curl "http://localhost:8080/api/movie?title=Dune%3A%20Part%20Two"
```

### Test the Frontend

1. Open browser to `http://localhost:3000`
2. Try searching for popular movies:
   - "Inception"
   - "The Matrix"
   - "Interstellar"
   - "Pulp Fiction"

---

## 🚢 Production Deployment

### Backend Deployment

1. **Build binary:**
   ```bash
   go build -o spoilerhub-api cmd/server/main.go
   ```

2. **Set environment variables:**
   ```bash
   export TMDB_API_KEY=your_key
   export GEMINI_API_KEY=your_key
   export ENVIRONMENT=production
   export PORT=8080
   ```

3. **Run:**
   ```bash
   ./spoilerhub-api
   ```

### Frontend Deployment

1. **Build for production:**
   ```bash
   npm run build
   ```

2. **Generate static files:**
   ```bash
   npm run generate
   ```

3. **Preview production build:**
   ```bash
   npm run preview
   ```

4. **Deploy to CDN/Server:**
   - Upload `.output/public` to your hosting provider
   - Set `NUXT_PUBLIC_API_BASE_URL` to your API server URL

---

## 📈 Performance Optimization

### Backend
- In-memory caching reduces API calls
- Efficient HTTP client pooling
- Minimal memory footprint
- Fast JSON marshaling/unmarshaling

### Frontend
- Server-side rendering (SSR) for fast initial load
- Component-level code splitting
- Image optimization (backdrop and poster)
- Lazy loading of components
- CSS minification with Tailwind

---

## 🐛 Troubleshooting

### Backend Won't Start
- Check if port 8080 is available
- Verify TMDB_API_KEY and GEMINI_API_KEY are set
- Check Go version is 1.21+

### Frontend Can't Connect to API
- Ensure backend is running on port 8080
- Check NUXT_PUBLIC_API_BASE_URL in `.env`
- Verify CORS middleware is enabled
- Check browser console for errors

### No Movie Results
- Verify TMDB API key is valid
- Check internet connection
- Try searching for popular movies first

### Spoiler Generation Fails
- Verify Gemini API key is valid
- Check API quota hasn't been exceeded
- Ensure movie title is correct

---

## 📄 Environment Variables Reference

### Backend (.env)
```env
PORT=8080                              # Server port (default: 8080)
ENVIRONMENT=development                # Environment (development/production)
TMDB_API_KEY=your_api_key_here        # TMDB API key (required)
GEMINI_API_KEY=your_api_key_here      # Gemini API key (required)
```

### Frontend (.env)
```env
NUXT_PUBLIC_API_BASE_URL=http://localhost:8080  # Backend API URL
```

---

## 🤝 Contributing

Contributions are welcome! Please follow the existing code style and structure.

---

## 📝 License

This project is open source and available under the MIT License.

---

## 🙋 Support

For issues or questions:
1. Check existing documentation
2. Review error messages
3. Check browser/server console logs
4. Verify API keys are valid

---

## 🎯 Future Enhancements

- [ ] User authentication and saved favorites
- [ ] Advanced search filters
- [ ] Movie ratings and reviews
- [ ] Social sharing features
- [ ] Dark/Light mode toggle (UI only, currently dark)
- [ ] Movie recommendations
- [ ] Spoiler rating system
- [ ] Content moderation

---

**Enjoy exploring movie spoilers with SpoilerHub! 🍿**
# spoilerHub
