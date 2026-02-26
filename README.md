# SpoilerHub - Production-Ready Full-Stack Movie Spoiler Application
# Landing Page
<img width="1345" height="597" alt="image" src="https://github.com/user-attachments/assets/ace22517-82d3-45a4-adc0-6aec9dceb0fe" />
<img width="1345" height="597" alt="image" src="https://github.com/user-attachments/assets/2f13ab36-88ab-4863-8209-68019833145c" />

## Detail Page
<img width="1339" height="563" alt="image" src="https://github.com/user-attachments/assets/c767d562-a605-49be-86a9-9ccea2ba86aa" />

## 📋 Overview

SpoilerHub is a modern full-stack web application that allows users to search for movies and receive:

- Movie poster and backdrop images
- Title, release year, and rating
- Genre information
- Movie overview
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
---


