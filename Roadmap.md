# 🎵 Groupie Tracker v2 – Complete Feature Roadmap

## 🎯 Project Vision

A modern music discovery platform combining concert information, artist discovery, and interactive visualizations. Built progressively from 01 Founders requirements to a full-featured Spotify-integrated application.

**Timeline**: 12-16 weeks total (MVP in 2-3 weeks)  
**Status**: Planning → Development

### Core User Flow

```
Landing Page
  → Multi-Provider Login (Spotify/Google/Apple)
  → Artist Discovery (main page - wildcard grid with playable tracks)
    → Search Bar (top-right, HTMX type-ahead)
    → Filters (sidebar panel)
    → Artist Card (click to expand)
      → Artist Detail Page (profile + stats + tracks/albums)
        → "Find Concerts" button
          → Concert Globe Page (3D map + concert list)
            → Concert Detail (click marker or list item)
  → User Profile/Settings
  → Error Pages (404, 401, 500, 403)
```

### Key Architectural Decisions

✅ **Fully responsive web app** (mobile-first, industry standard)  
✅ **Spotify API as primary data source** (richer data than 01 API)  
✅ **Go 1.22+ routing patterns** (method-specific handlers, no manual checks)  
✅ **Comprehensive middleware** (auth, logging, recovery, CORS, rate limiting)  
✅ **Go routines for performance** (concurrent API calls, worker pools)  
✅ **Testing throughout development** (unit, integration, HTTP, E2E)  
✅ **ASCII Art Web lessons applied** (proper error handling, middleware patterns)  
✅ **Professional error pages** (404, 401, 500 with redirects)

---

## 📄 Complete Page Inventory

### **Public Pages (No Auth Required)**

1. **Landing Page** (`/`)
   - Hero section with gradient background
   - Features showcase (3 cards)
   - Login/Sign Up CTAs
   - Footer with social links

2. **Login Page** (`/login`)
   - Multi-provider OAuth buttons (Spotify, Google, Apple)
   - Error state display for invalid credentials
   - Link to signup page
   - Redirect parameter support (`/login?redirect=/discover`)

3. **Sign Up Page** (`/signup`)
   - Same OAuth providers as login
   - Optional: Traditional email signup
   - Link to login page

---

### **Protected Pages (Auth Required - Middleware)**

4. **Artist Discovery** (`/discover`) - **MAIN PAGE AFTER LOGIN**
   - **Primary landing page for authenticated users**
   - Grid of artist wildcard with embedded track players
   - Search bar (top-right corner)
   - Filter panel (sidebar, collapsible on mobile)
   - Infinite scroll or pagination

   **Artist Card Components:**
   - Square artist image (hover: scale effect)
   - Artist name and genre tags
   - Mini player: 3 top tracks with play buttons
   - "View Profile" button
   - Follow/unfollow button

   **Interactions:**
   - Click card → Artist Detail page
   - Play button → 30-second Spotify preview
   - Search → HTMX type-ahead suggestions
   - Filters → Dynamic card updates (HTMX)

5. **Artist Detail** (`/artist/{id}`)
   - **Left side**: Large circular profile picture
   - **Right side**: Stats card
     - Follower count
     - Popularity score (0-100)
     - Genre tags
     - "Follow" button (changes to "Following" when active)
   - **Tabs**: Tracks, Albums, Concerts

   **Tracks Tab:**
   - Top 10 tracks list
   - Spotify embed players (30-second previews)
   - Like/save icons

   **Albums Tab:**
   - Grid of album covers
   - Release dates
   - Click to expand details

   **Concerts Tab:**
   - Prominent "Find Concerts" button
   - Preview of next 3 upcoming concerts
   - Button links to `/concerts?artist={id}`

6. **Concert Globe** (`/concerts?artist={id}` or `/concerts`)
   - **3D Globe** (70% width, right side)
     - Concert markers (color-coded by status)
     - Interactive rotation and zoom
     - Click marker → highlight in sidebar
     - Hover → tooltip with basic info
   - **Sidebar** (30% width, left side)
     - Artist filter dropdown (if viewing all concerts)
     - Date range picker
     - "Upcoming" / "Past" toggle
     - Concert list (scrollable):
       - Venue name
       - City, country
       - Date and time
       - "Get Tickets" button

   **Mobile**: Globe becomes full-width, sidebar moves to bottom drawer

7. **Concert Detail** (`/concert/{id}`)
   - Can be modal overlay OR dedicated page
   - Concert poster/hero image
   - Artist name and lineup
   - Venue details (name, address, capacity)
   - Date and time
   - Ticket prices and availability
   - "Buy Tickets" button (external link to Ticketmaster)
   - "Add to Calendar" button
   - "Save Concert" button
   - Venue map (embedded)

8. **User Profile** (`/profile` or `/profile/{username}`)
   - Profile header (avatar, name, joined date)
   - Stats: Followed artists count, saved concerts count
   - **Sections:**
     - Followed Artists (grid)
     - Saved Concerts (list)
     - Recently Viewed Artists
     - Listening History (Phase 4+)
   - Edit profile button (if own profile)

9. **User Settings** (`/settings`)
   - **Account Section:**
     - Display name
     - Email
     - Connected accounts (Spotify, Google, Apple)
     - Change password (if email auth)
   - **Preferences:**
     - Email notifications toggle
     - Concert alerts toggle
     - Newsletter subscription
   - **Appearance:**
     - Theme: Light / Dark / Auto
   - **Danger Zone:**
     - Delete account

---

### **Error Pages (Custom Handlers)**

10. **404 Not Found** (`/404`)
    - Friendly message: "Oops! This page doesn't exist"
    - Illustration or animation
    - "Go to Discovery" button
    - Search bar (maybe you're looking for...)

11. **401 Unauthorized** (`/401`)
    - Lock icon
    - Message: "Please log in to access this page"
    - "Log In" button
    - Redirect back after login

12. **403 Forbidden** (`/403`)
    - Shield icon
    - Message: "You don't have permission to access this"
    - "Go Back" button

13. **500 Server Error** (`/500`)
    - Error icon
    - Message: "Something went wrong on our end"
    - "Refresh Page" button
    - "Contact Support" link
    - Error ID for debugging

---

### **API Endpoints (for HTMX)**

These return HTML partials or JSON:

- `GET /api/search?q={query}` - Search suggestions (HTML snippet)
- `GET /api/artists?genre={genre}&popularity={min}-{max}` - Filtered artists
- `GET /api/artist/{id}/tracks` - Load tracks dynamically
- `POST /api/artist/{id}/follow` - Follow artist (returns updated button HTML)
- `DELETE /api/artist/{id}/follow` - Unfollow artist
- `GET /api/concerts/coordinates?artist={id}` - Globe data (JSON)
- `GET /api/concerts?location={city}&date={range}` - Concert list
- `POST /api/concert/{id}/save` - Save concert

---

## 🎨 Design System

### Color Palette (Apple Music Inspired + Orange Accents)

```css
/* Primary Backgrounds */
--bg-primary: #ffffff; /* Clean white (light mode) */
--bg-secondary: #f5f5f7; /* Light gray */
--bg-dark-primary: #000000; /* Pure black (dark mode) */
--bg-dark-secondary: #1c1c1e; /* Dark gray */
--bg-card: #2c2c2e; /* Card backgrounds (dark) */

/* Accent Colors */
--accent-red: #fc3c44; /* Apple Music red (primary CTA) */
--accent-orange: #ff9500; /* Orange (highlights, energy) */
--accent-purple: #bf5af2; /* Purple (premium features) */
--accent-pink: #ff2d55; /* Pink (favorites, likes) */

/* Gradients */
--gradient-hero: linear-gradient(135deg, #fc3c44, #ff9500);
--gradient-cta: linear-gradient(135deg, #ff9500, #fc3c44);
--gradient-premium: linear-gradient(135deg, #bf5af2, #fc3c44);

/* Text */
--text-primary: #1d1d1f; /* Dark text (light mode) */
--text-secondary: #86868b; /* Gray text */
--text-dark-primary: #ffffff; /* White text (dark mode) */
--text-dark-secondary: #98989d; /* Gray text (dark mode) */

/* Semantic Colors */
--success: #34c759; /* Green for success states */
--warning: #ff9500; /* Orange for warnings */
--error: #ff3b30; /* Red for errors */
--info: #007aff; /* Blue for info */
```

### Typography

```css
--font-primary:
  -apple-system, BlinkMacSystemFont, "SF Pro Display", "Segoe UI", system-ui,
  sans-serif;
--font-mono: "SF Mono", "Monaco", "Courier New", monospace;

--heading-1: 48px / 600; /* Hero headings */
--heading-2: 36px / 600; /* Section headings */
--heading-3: 24px / 600; /* Card titles */
--body-large: 18px / 400; /* Large body text */
--body: 16px / 400; /* Regular body text */
--body-small: 14px / 400; /* Small text, captions */
```

### Spacing System

```css
--space-xs: 4px;
--space-sm: 8px;
--space-md: 16px;
--space-lg: 24px;
--space-xl: 32px;
--space-2xl: 48px;
--space-3xl: 64px;
```

---

## 🔧 Go 1.22+ Routing Patterns (New!)

### Method-Specific Handlers

**Old Way (Pre-Go 1.22) - ❌ DON'T DO THIS:**

```go
func artistHandler(w http.ResponseWriter, r *http.Request) {
    // Manual method checking - tedious!
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Manual path parsing
    id := strings.TrimPrefix(r.URL.Path, "/artist/")

    // Handle GET logic
}

http.HandleFunc("/artist/", artistHandler)
```

**New Way (Go 1.22+) - ✅ DO THIS:**

```go
// Method is specified directly in the route!
http.HandleFunc("GET /artist/{id}", getArtistHandler)
http.HandleFunc("POST /artist/{id}/follow", followArtistHandler)
http.HandleFunc("DELETE /artist/{id}/follow", unfollowArtistHandler)

func getArtistHandler(w http.ResponseWriter, r *http.Request) {
    // No method check needed - Go handles it automatically!

    // Direct path variable access - no parsing needed!
    id := r.PathValue("id")

    // Handle GET logic
    artist, err := getArtistByID(id)
    if err != nil {
        http.Error(w, "Artist not found", http.StatusNotFound)
        return
    }

    renderTemplate(w, "artist-detail.html", artist)
}
```

### Wildcard Support

```go
// Wildcard for dynamic paths
http.HandleFunc("GET /artist/{id}/concerts/{location...}", concertsByLocationHandler)

func concertsByLocationHandler(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    location := r.PathValue("location") // Can be multiple segments
}
```

### Complete Routing Example

```go
package main

import (
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    // Public routes
    mux.HandleFunc("GET /", landingHandler)
    mux.HandleFunc("GET /login", loginHandler)
    mux.HandleFunc("POST /login", processLoginHandler)
    mux.HandleFunc("GET /signup", signupHandler)

    // OAuth callbacks
    mux.HandleFunc("GET /auth/spotify/callback", spotifyCallbackHandler)
    mux.HandleFunc("GET /auth/google/callback", googleCallbackHandler)
    mux.HandleFunc("GET /auth/apple/callback", appleCallbackHandler)

    // Protected routes (will add auth middleware)
    mux.HandleFunc("GET /discover", discoverHandler)
    mux.HandleFunc("GET /artist/{id}", getArtistHandler)
    mux.HandleFunc("POST /artist/{id}/follow", followArtistHandler)
    mux.HandleFunc("DELETE /artist/{id}/follow", unfollowArtistHandler)
    mux.HandleFunc("GET /concerts", concertsHandler)
    mux.HandleFunc("GET /concert/{id}", concertDetailHandler)
    mux.HandleFunc("GET /profile", profileHandler)
    mux.HandleFunc("GET /settings", settingsHandler)

    // API endpoints for HTMX
    mux.HandleFunc("GET /api/search", searchAPIHandler)
    mux.HandleFunc("GET /api/artists", artistsAPIHandler)
    mux.HandleFunc("GET /api/concerts/coordinates", concertCoordinatesHandler)

    // Error pages
    mux.HandleFunc("GET /404", notFoundHandler)
    mux.HandleFunc("GET /401", unauthorizedHandler)
    mux.HandleFunc("GET /500", serverErrorHandler)

    // Apply middleware chain
    handler := chainMiddleware(
        mux,
        recoveryMiddleware,
        loggingMiddleware,
        corsMiddleware,
    )

    http.ListenAndServe(":8080", handler)
}
```

**Benefits of Go 1.22+ Routing:**

- ✅ No manual method checking
- ✅ Direct path variable access with `r.PathValue()`
- ✅ Cleaner, more declarative code
- ✅ Better error handling (automatic 405 Method Not Allowed)
- ✅ More readable and maintainable

---

## 🔐 Middleware Strategy (ASCII Art Web Lessons Applied)

### Complete Middleware Stack

```go
// internal/middleware/middleware.go
package middleware

import (
    "context"
    "log"
    "net/http"
    "time"
)

// Middleware type
type Middleware func(http.Handler) http.Handler

// 1. Recovery Middleware (catch panics)
func Recovery(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("PANIC: %v", err)
                log.Printf("Request: %s %s", r.Method, r.URL.Path)

                // Redirect to 500 error page
                http.Redirect(w, r, "/500", http.StatusSeeOther)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

// 2. Logging Middleware
func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Log request
        log.Printf("→ %s %s %s", r.Method, r.URL.Path, r.RemoteAddr)

        // Wrap ResponseWriter to capture status code
        wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

        next.ServeHTTP(wrapped, r)

        // Log response
        duration := time.Since(start)
        log.Printf("← %s %s %d %v", r.Method, r.URL.Path, wrapped.statusCode, duration)
    })
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

// 3. Authentication Middleware
func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get session from cookie
        cookie, err := r.Cookie("session_token")
        if err != nil {
            // No session cookie - redirect to login
            redirectURL := "/login?redirect=" + r.URL.Path
            http.Redirect(w, r, redirectURL, http.StatusSeeOther)
            return
        }

        // Validate session
        user, err := validateSession(cookie.Value)
        if err != nil {
            // Invalid session - redirect to login
            http.Redirect(w, r, "/login?error=session_expired", http.StatusSeeOther)
            return
        }

        // Add user to request context
        ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// 4. CORS Middleware
func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

// 5. Rate Limiting Middleware (simple in-memory implementation)
func RateLimit(requestsPerMinute int) Middleware {
    // Use a simple map for demo - in production use Redis
    clients := make(map[string][]time.Time)

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ip := r.RemoteAddr
            now := time.Now()

            // Clean old requests
            if requests, exists := clients[ip]; exists {
                filtered := []time.Time{}
                for _, t := range requests {
                    if now.Sub(t) < time.Minute {
                        filtered = append(filtered, t)
                    }
                }
                clients[ip] = filtered
            }

            // Check rate limit
            if len(clients[ip]) >= requestsPerMinute {
                http.Error(w, "Too many requests", http.StatusTooManyRequests)
                return
            }

            // Add current request
            clients[ip] = append(clients[ip], now)

            next.ServeHTTP(w, r)
        })
    }
}

// 6. Custom 404 Handler Middleware
func Custom404(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        next.ServeHTTP(wrapped, r.WithContext(r.Context()))

        // If handler returned 404, show custom page
        if wrapped.statusCode == http.StatusNotFound {
            http.Redirect(w, r, "/404", http.StatusSeeOther)
        }
    })
}

// Chain multiple middleware
func Chain(h http.Handler, middleware ...Middleware) http.Handler {
    // Apply in reverse order so they execute in the order listed
    for i := len(middleware) - 1; i >= 0; i-- {
        h = middleware[i](h)
    }
    return h
}
```

### Usage Example

```go
package main

import (
    "net/http"
    "yourapp/internal/middleware"
)

func main() {
    // Public routes (no auth)
    publicMux := http.NewServeMux()
    publicMux.HandleFunc("GET /", landingHandler)
    publicMux.HandleFunc("GET /login", loginHandler)
    publicMux.HandleFunc("POST /login", processLoginHandler)

    // Protected routes (require auth)
    protectedMux := http.NewServeMux()
    protectedMux.HandleFunc("GET /discover", discoverHandler)
    protectedMux.HandleFunc("GET /artist/{id}", artistDetailHandler)
    protectedMux.HandleFunc("POST /artist/{id}/follow", followArtistHandler)
    protectedMux.HandleFunc("GET /profile", profileHandler)

    // Main mux that routes to public or protected
    mainMux := http.NewServeMux()
    mainMux.Handle("/", publicMux)
    mainMux.Handle("/discover", protectedMux)
    mainMux.Handle("/artist/", protectedMux)
    mainMux.Handle("/profile", protectedMux)

    // Apply middleware chain
    // Order matters! Recovery should be first (outermost)
    handler := middleware.Chain(
        mainMux,
        middleware.Recovery,           // 1. Catch panics (outermost)
        middleware.Logging,            // 2. Log requests/responses
        middleware.CORS,               // 3. Handle CORS
        middleware.RateLimit(100),     // 4. Rate limit (100 req/min)
        middleware.Custom404,          // 5. Custom 404 pages
    )

    // Protected routes get additional auth middleware
    protectedHandler := middleware.Chain(
        protectedMux,
        middleware.Auth,               // Check authentication
    )

    log.Println("Server starting on :8080")
    http.ListenAndServe(":8080", handler)
}
```

**Middleware Lessons from ASCII Art Web:**

- ✅ Always recover from panics (don't crash server)
- ✅ Log all requests for debugging
- ✅ Consistent error handling
- ✅ Clear separation of public vs protected routes
- ✅ Reusable middleware functions
- ✅ Proper order of middleware execution

---

## ⚡ Go Routines for Performance (Concurrent API Calls)

### Strategy: Parallel API Calls

**Problem**: Fetching artist data requires multiple API calls:

1. Artist details (Spotify)
2. Top tracks (Spotify)
3. Albums (Spotify)
4. Concerts (Ticketmaster)

Doing these sequentially = 4 × 200ms = 800ms total  
Doing these concurrently = max(200ms) = 200ms total ✨

### Implementation Pattern

```go
// internal/handlers/artist.go
package handlers

import (
    "net/http"
    "sync"
)

func getArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
    artistID := r.PathValue("id")

    // Result struct to collect all data
    type ArtistPageData struct {
        Artist   *SpotifyArtist
        Tracks   []Track
        Albums   []Album
        Concerts []Concert
        Error    error
    }

    data := &ArtistPageData{}
    var wg sync.WaitGroup

    // 1. Fetch artist details
    wg.Add(1)
    go func() {
        defer wg.Done()
        artist, err := spotifyClient.GetArtist(artistID)
        if err != nil {
            data.Error = err
            return
        }
        data.Artist = artist
    }()

    // 2. Fetch top tracks
    wg.Add(1)
    go func() {
        defer wg.Done()
        tracks, err := spotifyClient.GetTopTracks(artistID)
        if err != nil {
            log.Printf("Error fetching tracks: %v", err)
            return // Don't fail entire page for this
        }
        data.Tracks = tracks
    }()

    // 3. Fetch albums
    wg.Add(1)
    go func() {
        defer wg.Done()
        albums, err := spotifyClient.GetArtistAlbums(artistID)
        if err != nil {
            log.Printf("Error fetching albums: %v", err)
            return
        }
        data.Albums = albums
    }()

    // 4. Fetch concerts (once we have artist name)
    wg.Add(1)
    go func() {
        defer wg.Done()
        // Wait a bit for artist data first
        time.Sleep(50 * time.Millisecond)
        if data.Artist == nil {
            return
        }

        concerts, err := ticketmasterClient.GetArtistConcerts(data.Artist.Name)
        if err != nil {
            log.Printf("Error fetching concerts: %v", err)
            return
        }
        data.Concerts = concerts
    }()

    // Wait for all goroutines to complete
    wg.Wait()

    // Check for critical errors
    if data.Error != nil {
        http.Error(w, "Failed to load artist", http.StatusInternalServerError)
        return
    }

    // Render template with all data
    renderTemplate(w, "artist-detail.html", data)
}
```

### Worker Pool Pattern (Discovery Page)

```go
// Fetch multiple artists concurrently
func getDiscoveryHandler(w http.ResponseWriter, r *http.Request) {
    // Get artist IDs to fetch (from database or Spotify recommendations)
    artistIDs := getRecommendedArtistIDs(r.Context())

    // Worker pool to prevent too many concurrent requests
    numWorkers := 5 // Limit concurrent Spotify API calls
    jobs := make(chan string, len(artistIDs))
    results := make(chan *SpotifyArtist, len(artistIDs))
    errors := make(chan error, len(artistIDs))

    // Start worker goroutines
    var wg sync.WaitGroup
    for w := 0; w < numWorkers; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for id := range jobs {
                artist, err := spotifyClient.GetArtist(id)
                if err != nil {
                    errors <- err
                    continue
                }
                results <- artist
            }
        }()
    }

    // Send jobs to workers
    for _, id := range artistIDs {
        jobs <- id
    }
    close(jobs)

    // Wait for all workers to finish
    go func() {
        wg.Wait()
        close(results)
        close(errors)
    }()

    // Collect results
    artists := []*SpotifyArtist{}
    for artist := range results {
        artists = append(artists, artist)
    }

    // Log any errors (but don't fail the whole page)
    for err := range errors {
        log.Printf("Error fetching artist: %v", err)
    }

    // Render page
    renderTemplate(w, "discovery.html", map[string]interface{}{
        "Artists": artists,
    })
}
```

### Context with Timeout

```go
// Prevent goroutines from hanging indefinitely
func fetchWithTimeout(artistID string) (*SpotifyArtist, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    // Channel to receive result
    resultChan := make(chan *SpotifyArtist)
    errChan := make(chan error)

    // Fetch in goroutine
    go func() {
        artist, err := spotifyClient.GetArtist(artistID)
        if err != nil {
            errChan <- err
            return
        }
        resultChan <- artist
    }()

    // Wait for result or timeout
    select {
    case artist := <-resultChan:
        return artist, nil
    case err := <-errChan:
        return nil, err
    case <-ctx.Done():
        return nil, fmt.Errorf("request timed out")
    }
}
```

**Learning Checkpoint: Go Routines**

- Duration: 1.5 hours
- Topics to cover:
  - Goroutines basics (`go` keyword)
  - Channels (buffered vs unbuffered)
  - WaitGroups for synchronization
  - Worker pool pattern
  - Race conditions and mutexes
  - Context for cancellation and timeouts
  - Common pitfalls (goroutine leaks, deadlocks)
- When: Phase 2 (before Spotify API integration)
- Resources:
  - [Go by Example: Goroutines](https://gobyexample.com/goroutines)
  - [Go by Example: Channels](https://gobyexample.com/channels)
  - [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)

---

## 🧪 Testing Strategy Throughout Project

### 1. Unit Tests (Standard Library)

```go
// internal/services/spotify/spotify_test.go
package spotify

import (
    "testing"
)

func TestGetArtist(t *testing.T) {
    client := NewClient("test-token")

    artist, err := client.GetArtist("fake-id")

    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    if artist.Name == "" {
        t.Error("Expected artist name to be populated")
    }
}

// Table-driven tests
func TestSearchArtists(t *testing.T) {
    tests := []struct {
        name          string
        query         string
        expectedCount int
        expectError   bool
    }{
        {"Valid query", "Coldplay", 1, false},
        {"Empty query", "", 0, true},
        {"Special chars", "A$AP Rocky", 1, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            results, err := SearchArtists(tt.query)

            if tt.expectError && err == nil {
                t.Error("Expected error, got none")
            }

            if !tt.expectError && err != nil {
                t.Errorf("Unexpected error: %v", err)
            }

            if len(results) != tt.expectedCount {
                t.Errorf("Expected %d results, got %d",
                    tt.expectedCount, len(results))
            }
        })
    }
}

// Test with mocks
func TestGetArtistWithMock(t *testing.T) {
    // Create mock HTTP client
    mockClient := &MockHTTPClient{
        ResponseBody: `{"id": "123", "name": "Test Artist"}`,
        StatusCode:   200,
    }

    client := &SpotifyClient{httpClient: mockClient}
    artist, err := client.GetArtist("123")

    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if artist.Name != "Test Artist" {
        t.Errorf("Expected 'Test Artist', got '%s'", artist.Name)
    }
}
```

### 2. Integration Tests

```go
// tests/integration/database_test.go
package integration

import (
    "database/sql"
    "testing"
)

func TestDatabaseIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }

    // Setup test database
    db, err := sql.Open("postgres", "postgres://test:test@localhost/test_db")
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()

    // Run migrations
    if err := runMigrations(db); err != nil {
        t.Fatalf("Failed to run migrations: %v", err)
    }

    // Test insert
    result, err := db.Exec(
        "INSERT INTO artists (name, spotify_id) VALUES ($1, $2)",
        "Test Artist", "test123",
    )
    if err != nil {
        t.Errorf("Failed to insert: %v", err)
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected != 1 {
        t.Errorf("Expected 1 row affected, got %d", rowsAffected)
    }

    // Test query
    var name string
    err = db.QueryRow(
        "SELECT name FROM artists WHERE spotify_id = $1",
        "test123",
    ).Scan(&name)

    if err != nil {
        t.Errorf("Failed to query: %v", err)
    }

    if name != "Test Artist" {
        t.Errorf("Expected 'Test Artist', got '%s'", name)
    }

    // Cleanup
    db.Exec("DELETE FROM artists WHERE spotify_id = $1", "test123")
}
```

### 3. HTTP Handler Tests

```go
// internal/handlers/artist_test.go
package handlers

import (
    "context"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetArtistHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/artist/123", nil)

    // Simulate path value from Go 1.22 router
    req.SetPathValue("id", "123")

    w := httptest.NewRecorder()

    getArtistHandler(w, req)

    resp := w.Result()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, got %d", resp.StatusCode)
    }

    // Check Content-Type
    contentType := resp.Header.Get("Content-Type")
    if contentType != "text/html; charset=utf-8" {
        t.Errorf("Expected HTML content type, got %s", contentType)
    }
}

func TestFollowArtistHandler_Unauthorized(t *testing.T) {
    req := httptest.NewRequest("POST", "/artist/123/follow", nil)
    req.SetPathValue("id", "123")

    // No user in context (not authenticated)
    w := httptest.NewRecorder()

    followArtistHandler(w, req)

    resp := w.Result()

    // Should redirect to login
    if resp.StatusCode != http.StatusSeeOther {
        t.Errorf("Expected redirect (303), got %d", resp.StatusCode)
    }

    location := resp.Header.Get("Location")
    if location != "/login" {
        t.Errorf("Expected redirect to /login, got %s", location)
    }
}

func TestFollowArtistHandler_Success(t *testing.T) {
    req := httptest.NewRequest("POST", "/artist/123/follow", nil)
    req.SetPathValue("id", "123")

    // Add user to context (authenticated)
    user := &User{ID: 1, Name: "Test User"}
    ctx := context.WithValue(req.Context(), "user", user)
    req = req.WithContext(ctx)

    w := httptest.NewRecorder()

    followArtistHandler(w, req)

    resp := w.Result()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, got %d", resp.StatusCode)
    }
}
```

### 4. Middleware Tests

```go
// internal/middleware/middleware_test.go
package middleware

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestLoggingMiddleware(t *testing.T) {
    // Create a test handler
    testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    // Wrap with logging middleware
    handler := Logging(testHandler)

    req := httptest.NewRequest("GET", "/test", nil)
    w := httptest.NewRecorder()

    handler.ServeHTTP(w, req)

    resp := w.Result()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, got %d", resp.StatusCode)
    }
}

func TestAuthMiddleware_NoSession(t *testing.T) {
    testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    handler := Auth(testHandler)

    req := httptest.NewRequest("GET", "/protected", nil)
    // No session cookie

    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    resp := w.Result()

    // Should redirect to login
    if resp.StatusCode != http.StatusSeeOther {
        t.Errorf("Expected redirect, got %d", resp.StatusCode)
    }
}

func TestRecoveryMiddleware_Panic(t *testing.T) {
    // Handler that panics
    panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        panic("test panic")
    })

    handler := Recovery(panicHandler)

    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()

    // Should not crash
    handler.ServeHTTP(w, req)

    resp := w.Result()

    // Should redirect to 500 page
    if resp.StatusCode != http.StatusSeeOther {
        t.Errorf("Expected redirect to error page, got %d", resp.StatusCode)
    }
}
```

### 5. E2E Tests (Optional - Phase 6)

```go
// tests/e2e/user_flow_test.go
package e2e

import (
    "testing"
)

// Using a headless browser library like chromedp or playwright
func TestCompleteUserFlow(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping E2E test")
    }

    // 1. Visit landing page
    // 2. Click "Login with Spotify"
    // 3. Complete OAuth (mock or test credentials)
    // 4. Verify redirect to discovery page
    // 5. Search for an artist
    // 6. Click artist card
    // 7. Verify artist detail page loads
    // 8. Click "Find Concerts"
    // 9. Verify globe page loads
    // 10. Click concert marker
    // 11. Verify concert detail appears
}
```

### Testing Checklist by Phase

**Phase 1 (MVP):**

- [ ] Unit tests for database functions
- [ ] Unit tests for search logic (case-insensitive, filtering)
- [ ] Unit tests for filter logic (range, checkbox combinations)
- [ ] HTTP handler tests (all routes)
- [ ] Middleware tests (logging, recovery, auth)
- [ ] Error page tests (404, 500)

**Phase 2 (Auth + Spotify):**

- [ ] OAuth flow tests (mocked)
- [ ] Session management tests
- [ ] Spotify API client tests (mocked)
- [ ] Integration tests (DB + auth)
- [ ] Follow/unfollow functionality tests

**Phase 3 (Concert APIs):**

- [ ] Ticketmaster client tests
- [ ] Geocoding service tests
- [ ] Concert filtering tests
- [ ] Globe data transformation tests

**Phase 4+ (AI, Performance):**

- [ ] OpenAI integration tests (mocked)
- [ ] Go routine tests (concurrent operations)
- [ ] Performance/load tests
- [ ] E2E user flow tests

### Running Tests

```makefile
# Add to Makefile
.PHONY: test test-unit test-integration test-coverage

test:
	go test -v ./...

test-unit:
	go test -short -v ./...

test-integration:
	go test -v ./tests/integration/...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-race:
	go test -race ./...
```

**Test Commands:**

```bash
# Run all tests
make test

# Run only unit tests (skip integration)
make test-unit
# or: go test -short ./...

# Run with coverage report
make test-coverage

# Check for race conditions
make test-race

# Run specific test
go test -v -run TestGetArtist ./internal/services/spotify

# Run tests in specific package
go test ./internal/handlers/...
```

**Coverage Goals:**

- Phase 1: 60%+ coverage
- Phase 2: 70%+ coverage
- Phase 3+: 75%+ coverage
- Critical paths (auth, payments): 90%+ coverage

---

## 🔹 PHASE 1: MVP (2-3 Weeks) + 01 Core Features

**Goal**: Portfolio-ready, deployable application meeting all 01 Founders requirements.

### Week 1: Foundation + Search

#### **Days 1-2: Project Setup & Docker** 🐳

**Features:**

- [ ] Project structure with Go 1.22+ patterns
- [ ] Docker Compose setup (PostgreSQL + Go app)
- [ ] Environment variable management (.env)
- [ ] Makefile with new testing commands
- [ ] Hot reload with Air
- [ ] Database migrations system
- [ ] Middleware stack (recovery, logging, CORS)
- [ ] Custom error pages (404, 401, 500)
- [ ] Fetch data from 01 Groupie Tracker API
- [ ] Cache API data in PostgreSQL

**Project Structure:**

```
Groupie-Tracker/
├── cmd/
│   └── api/
│       └── main.go (Go 1.22 routing)
├── internal/
│   ├── middleware/ (new!)
│   │   └── middleware.go
│   ├── handlers/
│   │   ├── artist.go
│   │   ├── discover.go
│   │   ├── concert.go
│   │   └── error.go (new!)
│   ├── services/
│   │   ├── spotify/
│   │   └── ticketmaster/
│   ├── models/
│   └── database/
├── tests/ (new!)
│   ├── integration/
│   └── e2e/
├── web/
│   ├── templates/
│   │   ├── landing.html
│   │   ├── login.html
│   │   ├── discover.html
│   │   ├── artist-detail.html
│   │   ├── concerts.html
│   │   ├── error-404.html (new!)
│   │   ├── error-401.html (new!)
│   │   └── error-500.html (new!)
│   └── static/
├── docker-compose.yml
├── Dockerfile
├── Makefile
└── README.md
```

**Tech Stack:**

- **Backend**: Go 1.22+ (new routing patterns)
- **Database**: PostgreSQL 16
- **External Packages**: `lib/pq`, `godotenv`, `golang.org/x/oauth2`
- **Containerization**: Docker + Docker Compose
- **Hot Reload**: Air

**Learning Checkpoints:**

- ✅ Go 1.22 routing patterns (45 mins)
- ✅ Middleware chaining (30 mins)
- ✅ PostgreSQL basics vs SQLite (20 mins)
- ✅ Docker Compose fundamentals (30 mins)

---

#### **Days 3-5: Search Bar with HTMX** 🔍

**01 Brief**: `search-bar_README.md` ✅

**Objectives:**

- Artist Discovery page as main authenticated landing
- Real-time search with HTMX type-ahead
- Meet all 01 search requirements

**Features:**

- [ ] **Artist Discovery page** (main page after login)
  - Wildcard grid layout (responsive)
  - Artist cards with embedded track players
  - Search bar (top-right corner)
- [ ] Search bar component with HTMX type-ahead
- [ ] Search categories:
  - Artist/band name
  - Members
  - Locations
  - First album date
  - Creation date
- [ ] Category labeling (e.g., "Phil Collins - member")
- [ ] Case-insensitive search
- [ ] Debounced suggestions (300ms delay)
- [ ] Keyboard navigation (arrow keys, enter, escape)
- [ ] Empty state handling
- [ ] Unit tests for search logic
- [ ] HTTP handler tests

**Artist Card Components:**

```html
<div class="artist-card">
  <img src="artist-image.jpg" alt="Artist Name" />
  <h3>Artist Name</h3>
  <div class="genre-tags">
    <span>Rock</span>
    <span>Alternative</span>
  </div>
  <div class="mini-player">
    <div class="track">
      <button class="play-btn">▶</button>
      <span>Track Name</span>
      <span class="duration">3:45</span>
    </div>
    <!-- 2 more tracks -->
  </div>
  <button class="btn-view-profile">View Profile</button>
  <button class="btn-follow">Follow</button>
</div>
```

**HTMX Implementation:**

```html
<input
  type="text"
  name="q"
  hx-get="/api/search"
  hx-trigger="keyup changed delay:300ms"
  hx-target="#search-results"
  hx-indicator="#search-spinner"
  placeholder="Search artists, members, locations..."
/>
<div id="search-results"></div>
```

**Backend Endpoints (Go 1.22):**

```go
mux.HandleFunc("GET /api/search", searchAPIHandler)
mux.HandleFunc("GET /discover", discoverHandler) // Main page
```

**Learning Checkpoint:**

- ✅ HTMX core concepts (45 mins)
- ✅ `hx-get`, `hx-post`, `hx-target`, `hx-swap` (30 mins)
- ✅ Debouncing and performance (15 mins)
- ✅ Writing unit tests (30 mins)

---

#### **Days 6-7: Artist Detail Page**

**Features:**

- [ ] Artist detail page layout
- [ ] Profile picture (top-left, large circular)
- [ ] Stats card (top-right):
  - Follower count
  - Popularity score
  - Genre tags
  - "Follow" button (dynamic state with HTMX)
- [ ] Tabs: Tracks, Albums, Concerts
- [ ] Tracks tab with Spotify embeds
- [ ] Albums grid
- [ ] "Find Concerts" button in Concerts tab
- [ ] Responsive design (mobile stacks vertically)
- [ ] Unit tests for artist data fetching
- [ ] HTTP handler tests

**Route (Go 1.22):**

```go
mux.HandleFunc("GET /artist/{id}", getArtistHandler)
mux.HandleFunc("POST /artist/{id}/follow", followArtistHandler)
```

---

### Week 2: Globe + Filters

#### **Days 8-10: Geolocalization with 3D Globe** 🌍

**01 Brief**: `Geolocalization.md` ✅

**Objectives:**

- Interactive 3D globe for concert visualization
- Geocoding integration
- Concert filtering on globe

**Features:**

- [ ] Geocoding service integration (Nominatim)
- [ ] Convert concert locations to coordinates
- [ ] Store coordinates in database
- [ ] 3D globe with Globe.gl
- [ ] Concert markers (color-coded)
- [ ] Click marker → show concert details
- [ ] Interactive rotation, zoom, controls
- [ ] Sidebar with concert list
- [ ] Filter by artist, date range
- [ ] Upcoming/past toggle
- [ ] Mobile-responsive (drawer layout)
- [ ] Unit tests for geocoding
- [ ] Integration tests for concert data

**Route (Go 1.22):**

```go
mux.HandleFunc("GET /concerts", concertsGlobeHandler)
mux.HandleFunc("GET /api/concerts/coordinates", concertCoordinatesAPIHandler)
```

**Learning Checkpoint:**

- ✅ Three.js basics (1 hour)
- ✅ Globe.gl library (30 mins)
- ✅ Geocoding APIs (30 mins)

---

#### **Days 11-13: Filters** 🔧

**01 Brief**: `Filters_README.md` ✅

**Objectives:**

- Multi-criteria filtering with HTMX
- Range and checkbox filters
- Dynamic updates

**Features:**

- [ ] Filter sidebar on Discovery page
- [ ] Creation date (range slider)
- [ ] First album date (range slider)
- [ ] Number of members (checkboxes)
- [ ] Concert locations (multi-select checkboxes)
- [ ] Active/inactive toggle
- [ ] HTMX dynamic filtering
- [ ] URL state preservation (shareable links)
- [ ] Active filter chips display
- [ ] Result count
- [ ] "Clear Filters" button
- [ ] Mobile drawer/modal layout
- [ ] Unit tests for filter logic
- [ ] Integration tests

**Route (Go 1.22):**

```go
mux.HandleFunc("GET /api/artists/filter", artistFilterAPIHandler)
```

**Learning Checkpoint:**

- ✅ HTMX form handling (45 mins)
- ✅ URL state management (30 mins)
- ✅ Complex SQL queries (45 mins)

---

#### **Day 14: UI Polish & Error Pages** 🎨

**01 Brief**: `visualizations_README.md` ✅

**Features:**

- [ ] Consistent design system implementation
- [ ] Custom 404 page (friendly, helpful)
- [ ] Custom 401 page (redirect to login)
- [ ] Custom 500 page (error recovery)
- [ ] Loading states (skeletons)
- [ ] Empty states
- [ ] Success toasts
- [ ] Smooth transitions
- [ ] Accessibility audit (ARIA, keyboard nav)
- [ ] Mobile responsiveness testing
- [ ] Cross-browser testing

**Error Page Routes (Go 1.22):**

```go
mux.HandleFunc("GET /404", notFoundHandler)
mux.HandleFunc("GET /401", unauthorizedHandler)
mux.HandleFunc("GET /500", serverErrorHandler)
```

---

### Week 3: Docker Production + Deployment

#### **Days 15-17: Production Containerization** 🐳

**Features:**

- [ ] Multi-stage Dockerfile optimized
- [ ] Docker Compose production config
- [ ] Environment management (dev/prod)
- [ ] Health check endpoints
- [ ] Graceful shutdown
- [ ] Log aggregation
- [ ] Database backup strategy
- [ ] Security hardening
- [ ] Load testing preparation
- [ ] Performance benchmarks

**Health Check:**

```go
mux.HandleFunc("GET /health", healthCheckHandler)
mux.HandleFunc("GET /ready", readinessCheckHandler)
```

---

#### **Days 18-19: Deployment** 🚀

**Platform**: Render / Railway / Fly.io

**Features:**

- [ ] Deploy PostgreSQL database
- [ ] Deploy Go application container
- [ ] Environment variables configuration
- [ ] SSL certificates (automatic)
- [ ] Custom domain (optional)
- [ ] Monitoring setup
- [ ] Error tracking
- [ ] Uptime monitoring
- [ ] Performance monitoring
- [ ] Test all routes in production

---

#### **Days 20-21: Documentation & Testing** 📝

**Features:**

- [ ] Comprehensive README
  - Architecture diagram
  - API documentation
  - Setup instructions
  - Screenshots/GIFs
  - Live demo link
- [ ] Code comments and documentation
- [ ] Demo video (2-3 minutes)
- [ ] Database schema documentation
- [ ] Test coverage report (aim for 70%+)
- [ ] Integration test suite
- [ ] Performance benchmarks documented

---

## 🔹 PHASE 2: Authentication + Spotify Integration (Weeks 4-6)

### 2.1: Multi-Provider OAuth

**Features:**

- [ ] OAuth 2.0 implementation (Spotify, Google, Apple)
- [ ] Session management
- [ ] User profile creation
- [ ] "Remember me" functionality
- [ ] Logout with session cleanup
- [ ] Account linking
- [ ] OAuth tests (mocked)
- [ ] Session tests

**Routes (Go 1.22):**

```go
mux.HandleFunc("GET /login", loginPageHandler)
mux.HandleFunc("GET /auth/spotify/callback", spotifyCallbackHandler)
mux.HandleFunc("GET /auth/google/callback", googleCallbackHandler)
mux.HandleFunc("GET /auth/apple/callback", appleCallbackHandler)
mux.HandleFunc("POST /logout", logoutHandler)
```

**Learning Checkpoint:**

- ✅ OAuth 2.0 flow (1 hour)
- ✅ Token management (30 mins)
- ✅ Session vs JWT (30 mins)

---

### 2.2: Spotify API Integration with Go Routines

**Features:**

- [ ] Spotify API service layer
- [ ] **Concurrent API calls with goroutines**
- [ ] Artist details fetching
- [ ] Top tracks fetching
- [ ] Albums fetching
- [ ] Related artists fetching
- [ ] Spotify embed player integration
- [ ] Data caching in PostgreSQL
- [ ] Rate limiting handling
- [ ] Error recovery
- [ ] Unit tests with mocks
- [ ] Integration tests

**Goroutine Implementation:**

```go
// Fetch multiple data points concurrently
var wg sync.WaitGroup
wg.Add(3)

go fetchArtistDetails(&wg, artistID)
go fetchTopTracks(&wg, artistID)
go fetchAlbums(&wg, artistID)

wg.Wait()
```

**Learning Checkpoint:**

- ✅ **Go routines and channels** (1.5 hours)
- ✅ Worker pools (45 mins)
- ✅ Race conditions (30 mins)

---

### 2.3: User Features

**Features:**

- [ ] Follow/unfollow artists (with goroutine background tasks)
- [ ] User dashboard
- [ ] My Library page
- [ ] Saved concerts
- [ ] User settings page
- [ ] Profile customization
- [ ] Notification preferences
- [ ] Search history
- [ ] Recently viewed
- [ ] Unit tests
- [ ] HTTP handler tests

**Routes (Go 1.22):**

```go
mux.HandleFunc("GET /profile", profileHandler)
mux.HandleFunc("GET /settings", settingsHandler)
mux.HandleFunc("GET /library", libraryHandler)
```

---

## 🔹 PHASE 3-6: (Same as before, with testing integrated)

[Keep all the Phase 3-6 content from previous version]

---

## 🎨 v0 Prompt (Complete Redesign)

```
Create a modern music discovery web application with Apple Music aesthetic and this specific structure:

DESIGN STYLE:
- Apple Music inspired
- Colors: White/light (#FFFFFF, #F5F5F7), red (#FC3C44), orange (#FF9500), purple (#BF5AF2)
- SF Pro Display font (Apple system fonts)
- Clean, premium, fully responsive

USER FLOW:
Landing Page → Login (OAuth) → Discovery Page (main) → Artist Detail → Concerts Globe

PAGES:

1. LANDING PAGE (/)
- Gradient hero (red to orange)
- "Discover Your Music Universe"
- Two CTAs: "Sign Up" and "Login"
- 3 feature cards
- Footer with social links

2. LOGIN PAGE (/login)
- Centered card
- OAuth buttons: Spotify (green), Google, Apple (black)
- Error banner for invalid credentials
- Link to signup

3. DISCOVERY PAGE (/discover) - MAIN AFTER LOGIN
Top nav: Logo | Search bar (top-right) | User avatar
Left sidebar (collapsible): Filters
Main area: Grid of artist cards (4 cols desktop, 2 tablet, 1 mobile)

Artist Card:
- Square image (rounded)
- Artist name
- Genre tags (pills)
- Mini player (3 tracks with play buttons)
- "View Profile" button
- "Follow" button
- Hover: lift effect

4. ARTIST DETAIL (/artist/{id})
Top: Large circular profile (left) | Stats card (right: followers, popularity, genres, follow button)
Tabs: Tracks | Albums | Concerts

Tracks tab:
- List of tracks with Spotify embeds
- Play buttons
- Like icons

Concerts tab:
- Large "Find Concerts" button (red gradient)

5. CONCERT GLOBE (/concerts)
Layout: 3D globe (70% width) | Sidebar (30%)
Globe: Red markers, click to show details
Sidebar: Artist filter, date range, concert list with "Get Tickets"

6. ERROR PAGES
404: "404" + message + "Go Home"
401: Lock icon + "Please log in"
500: Error icon + "Something went wrong"

TECHNICAL:
- Mobile-first responsive
- Smooth transitions
- Loading skeletons
- ARIA labels
- Keyboard navigation
```

---

## 📊 Complete Technology Stack

### Backend

- **Language**: Go 1.22+ (new routing patterns)
- **Router**: `net/http` with method-specific handlers
- **Database Driver**: `github.com/lib/pq`
- **OAuth**: `golang.org/x/oauth2`
- **Environment**: `github.com/joho/godotenv`
- **Testing**: Go standard library + table-driven tests

### Middleware

- Recovery (panic handling)
- Logging (request/response)
- Authentication (session validation)
- CORS (cross-origin requests)
- Rate limiting (in-memory or Redis)
- Custom 404 handler

### Concurrency

- Go routines for parallel API calls
- Worker pools for bulk operations
- Channels for communication
- WaitGroups for synchronization
- Context for timeouts

### Frontend

- HTML5, CSS3, Vanilla JavaScript
- HTMX (progressive enhancement)
- Globe.gl (Three.js)
- Spotify embed players

### Testing

- Unit tests (services, handlers)
- Integration tests (database, API)
- HTTP handler tests
- Middleware tests
- E2E tests (optional)
- Coverage target: 70%+

---

## 🎯 Success Criteria

### 01 Founders Requirements ✅

- ✅ All briefs completed (search, filters, geo, viz)
- ✅ Standard Go library focus
- ✅ Professional error handling

### Technical Excellence ✅

- ✅ Go 1.22+ routing patterns
- ✅ Comprehensive middleware
- ✅ Go routines for performance
- ✅ 70%+ test coverage
- ✅ Responsive design
- ✅ Professional error pages

### Portfolio Quality ✅

- ✅ Clean, documented code
- ✅ Live demo deployed
- ✅ Architecture diagram
- ✅ Comprehensive README

---

**Last Updated**: January 2025  
**Version**: 2.0.0  
**Status**: Ready for Development  
**Maintainer**: Ibraheem Youssef (@IbsYoussef)
