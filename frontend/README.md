# SpoilerHub Frontend

## Quick Start for Frontend

### Prerequisites
- Node.js 18+
- npm or yarn

### Setup

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Configure backend URL (optional):**
   ```bash
   cp .env.example .env
   # Edit .env if your backend is not on http://localhost:8080
   ```

3. **Start development server:**
   ```bash
   npm run dev
   ```

Frontend will be available at `http://localhost:3000`


## Features

- Server-side rendering (SSR)
- Dark theme UI with Tailwind CSS
- Responsive design (mobile, tablet, desktop)
- Client-side state management
- Error handling
- Loading states
- Smooth animations

## Composables

### useMovieAPI
Manages API interactions and state:
- `fetchMovie(title)` - Search for movie
- `clearMovie()` - Clear current movie
- Reactive: `movie`, `loading`, `error`

## Components

### MovieSearch
Search input component with:
- Form validation
- Submit handling
- Loading state
- Error display

### MovieDetails
Movie information display with:
- Poster and backdrop images
- Rating badge
- Genre tags
- Movie overview
- Hidden spoiler section (click to reveal)
- Back button

### Default Layout
Main layout with:
- Header with branding
- Content slot
- Footer with attribution
