# 🎵 Groupie Tracker v2 – Complete Feature Roadmap

## 🎯 Project Vision

A modern music discovery platform combining concert information, artist discovery, and interactive visualizations. Built progressively from 01 Founders requirements to a full-featured Spotify-integrated application.

**Timeline**: 12-16 weeks total (MVP in 2-3 weeks)  
**Status**: Planning → Development

---

## 🎨 Design System

### Color Palette (Apple Music Inspired + Orange Accents)

```css
/* Primary Backgrounds */
--bg-primary: #FFFFFF;          /* Clean white (light mode) */
--bg-secondary: #F5F5F7;        /* Light gray */
--bg-dark-primary: #000000;     /* Pure black (dark mode) */
--bg-dark-secondary: #1C1C1E;   /* Dark gray */
--bg-card: #2C2C2E;             /* Card backgrounds (dark) */

/* Accent Colors */
--accent-red: #FC3C44;          /* Apple Music red (primary CTA) */
--accent-orange: #FF9500;       /* Orange (highlights, energy) */
--accent-purple: #BF5AF2;       /* Purple (premium features) */
--accent-pink: #FF2D55;         /* Pink (favorites, likes) */

/* Gradients */
--gradient-hero: linear-gradient(135deg, #FC3C44, #FF9500);
--gradient-cta: linear-gradient(135deg, #FF9500, #FC3C44);
--gradient-premium: linear-gradient(135deg, #BF5AF2, #FC3C44);

/* Text */
--text-primary: #1D1D1F;        /* Dark text (light mode) */
--text-secondary: #86868B;      /* Gray text */
--text-dark-primary: #FFFFFF;   /* White text (dark mode) */
--text-dark-secondary: #98989D; /* Gray text (dark mode) */

/* Semantic Colors */
--success: #34C759;             /* Green for success states */
--warning: #FF9500;             /* Orange for warnings */
--error: #FF3B30;               /* Red for errors */
--info: #007AFF;                /* Blue for info */
```

### Typography

```css
--font-primary: -apple-system, BlinkMacSystemFont, 'SF Pro Display', 'Segoe UI', system-ui, sans-serif;
--font-mono: 'SF Mono', 'Monaco', 'Courier New', monospace;

--heading-1: 48px / 600;        /* Hero headings */
--heading-2: 36px / 600;        /* Section headings */
--heading-3: 24px / 600;        /* Card titles */
--body-large: 18px / 400;       /* Large body text */
--body: 16px / 400;             /* Regular body text */
--body-small: 14px / 400;       /* Small text, captions */
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

## 🔹 PHASE 1: MVP (2-3 Weeks) + 01 Core Features

**Goal**: Portfolio-ready, deployable application meeting all 01 Founders requirements.

### Week 1: Foundation + Search

#### **Days 1-2: Project Setup & Docker** 🐳

**Features:**
- [ ] Project structure (Go modules, directory layout)
- [ ] Docker Compose setup (PostgreSQL + Go app)
- [ ] Environment variable management (.env)
- [ ] Makefile for development commands
- [ ] Hot reload with Air
- [ ] Database migrations system
- [ ] Fetch data from 01 Groupie Tracker API
- [ ] Cache API data in PostgreSQL

**Docker Setup:**
```yaml
services:
  - postgres:16-alpine (database)
  - app:golang:1.22-alpine (Go application with hot reload)
volumes:
  - postgres_data (persistent database)
```

**Database Schema (Initial):**
```sql
-- Core tables
CREATE TABLE artists (
    id SERIAL PRIMARY KEY,
    api_id INTEGER UNIQUE,
    name VARCHAR(255),
    image_url VARCHAR(500),
    members TEXT[],
    creation_date INTEGER,
    first_album DATE,
    data JSONB
);

CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    artist_id INTEGER REFERENCES artists(id),
    location VARCHAR(255),
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8)
);

CREATE TABLE concerts (
    id SERIAL PRIMARY KEY,
    artist_id INTEGER REFERENCES artists(id),
    location_id INTEGER REFERENCES locations(id),
    concert_date DATE
);
```

**Tech Stack:**
- **Backend**: Go 1.22+ (standard library)
- **Database**: PostgreSQL 16
- **External Packages**: `lib/pq`, `godotenv`
- **Containerization**: Docker + Docker Compose
- **Hot Reload**: Air

**Learning Checkpoint:**
- ✅ PostgreSQL basics vs SQLite (20 mins)
- ✅ Docker Compose fundamentals (30 mins)
- ✅ Go database/sql package (30 mins)

---

#### **Days 3-5: Search Bar with HTMX** 🔍

**01 Brief**: `search-bar_README.md` ✅

**Objectives:**
- Implement intelligent search with real-time suggestions
- Introduce HTMX for dynamic updates
- Meet all 01 search requirements

**Features:**
- [ ] Artist listing page with cards/grid layout
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
- [ ] Search results page
- [ ] Empty state handling

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

**Backend Endpoints:**
- `GET /api/search?q={query}` - Returns HTML snippet with suggestions
- `GET /artists?q={query}` - Full search results page

**Learning Checkpoint:**
- ✅ HTMX core concepts (45 mins)
- ✅ `hx-get`, `hx-post`, `hx-target`, `hx-swap` (30 mins)
- ✅ Debouncing and performance (15 mins)
- ✅ Returning HTML partials from Go handlers (30 mins)

---

#### **Days 6-7: Artist Detail Page**

**Features:**
- [ ] Artist detail page layout
- [ ] Display artist information (name, image, members, dates)
- [ ] List of concerts/tour dates
- [ ] Concert locations
- [ ] Related information from API
- [ ] Responsive design (mobile + desktop)
- [ ] Back navigation
- [ ] Share button (future: social sharing)

**Route:**
- `GET /artist/{id}` - Artist detail page

**UI Components:**
- Hero section with artist image
- Info cards (members, dates)
- Concert list/timeline
- Location tags

---

### Week 2: Globe + Filters

#### **Days 8-10: Geolocalization with 3D Globe** 🌍

**01 Brief**: `Geolocalization.md` ✅

**Objectives:**
- Map concert locations on interactive 3D globe
- Convert addresses to coordinates
- Create portfolio-standout visual feature

**Features:**
- [ ] Geocoding service integration (Nominatim or Google)
- [ ] Convert concert locations to lat/long coordinates
- [ ] 3D globe visualization with Globe.gl
- [ ] Concert location markers (color-coded by date)
- [ ] Click marker → show concert details modal/panel
- [ ] Interactive rotation and zoom
- [ ] Globe controls (rotate, zoom, reset view)
- [ ] Filter concerts by artist on globe
- [ ] Upcoming vs past concerts toggle
- [ ] Smooth animations and transitions
- [ ] Mobile-responsive globe (touch gestures)

**Tech Stack:**
- **3D Library**: Globe.gl (Three.js wrapper)
- **Geocoding**: Nominatim (free, no API key) or Google Geocoding API
- **Backend**: Go endpoints for coordinate data

**Globe Features:**
```javascript
const globe = Globe()
  .globeImageUrl('/static/images/earth-texture.jpg')
  .pointsData(concerts)
  .pointLat(d => d.latitude)
  .pointLng(d => d.longitude)
  .pointColor(d => d.isPast ? '#86868B' : '#FC3C44')
  .pointAltitude(0.01)
  .pointRadius(0.6)
  .pointLabel(d => `${d.artist} - ${d.location}`)
  .onPointClick(concert => showConcertDetails(concert))
  .backgroundColor('#000000');
```

**UI Layout:**
```
┌─────────────────────────────────────┐
│  🌍 3D Globe (full width)           │
│                                      │
│  [Sidebar]          [Globe Canvas]  │
│  - Artist filter                    │
│  - Date range                       │
│  - Upcoming/Past                    │
│  - Concert list                     │
└─────────────────────────────────────┘
```

**Backend Endpoints:**
- `GET /api/concerts/coordinates` - Returns concert data with lat/long
- `GET /api/geocode?location={address}` - Geocoding service

**Learning Checkpoint:**
- ✅ Three.js basics (1 hour)
- ✅ Globe.gl library usage (30 mins)
- ✅ Geocoding APIs (30 mins)
- ✅ Coordinate systems (lat/long) (15 mins)

---

#### **Days 11-13: Filters** 🔧

**01 Brief**: `Filters_README.md` ✅

**Objectives:**
- Multi-criteria filtering with HTMX
- Range filters and checkbox filters
- Dynamic result updates without page reload

**Required Filters:**
- [ ] Creation date (range slider: 1960-2024)
- [ ] First album date (range slider)
- [ ] Number of members (checkboxes: 1-2, 3-4, 5+, 6+)
- [ ] Concert locations (multi-select checkboxes by country/region)

**Additional Filters:**
- [ ] Active/inactive bands (checkbox)
- [ ] Has upcoming concerts (checkbox)

**Filter Types:**
1. **Range Filter** - Dual-handle slider for date ranges
2. **Checkbox Filter** - Multiple selection for categories

**HTMX Implementation:**
```html
<form hx-get="/api/artists/filter" 
      hx-target="#artist-results" 
      hx-trigger="change, search from:#filter-form"
      hx-swap="innerHTML"
      id="filter-form">
  
  <!-- Range Filter: Creation Date -->
  <div class="filter-group">
    <label>Creation Date</label>
    <input type="range" name="creation_from" min="1960" max="2024" value="1960">
    <input type="range" name="creation_to" min="1960" max="2024" value="2024">
    <span id="creation-range-display">1960 - 2024</span>
  </div>
  
  <!-- Checkbox Filter: Members -->
  <div class="filter-group">
    <label>Number of Members</label>
    <label><input type="checkbox" name="members" value="1-2"> 1-2</label>
    <label><input type="checkbox" name="members" value="3-4"> 3-4</label>
    <label><input type="checkbox" name="members" value="5+"> 5+</label>
  </div>
  
  <!-- Location Filter -->
  <div class="filter-group">
    <label>Concert Locations</label>
    <label><input type="checkbox" name="location" value="USA"> USA</label>
    <label><input type="checkbox" name="location" value="UK"> UK</label>
    <label><input type="checkbox" name="location" value="France"> France</label>
    <!-- ... more locations ... -->
  </div>
  
  <button type="button" hx-get="/api/artists" hx-target="#artist-results">
    Clear Filters
  </button>
</form>

<div id="artist-results">
  <!-- Filtered artists appear here -->
</div>
```

**Filter Features:**
- HTMX automatic updates on filter change
- URL state preservation (shareable filtered URLs: `/artists?creation_from=1980&members=3-4`)
- Active filter chips/tags display
- Filter result count ("Showing 42 of 156 artists")
- Smooth transitions when results update
- Mobile-friendly filter drawer/panel

**Backend Endpoints:**
- `GET /api/artists/filter?creation_from=1980&members=3-4&location=USA`

**Learning Checkpoint:**
- ✅ HTMX form handling (45 mins)
- ✅ URL state management (30 mins)
- ✅ Multi-criteria database queries (45 mins)

---

#### **Day 14: UI Polish & Visualizations** 🎨

**01 Brief**: `visualizations_README.md` ✅

**Objectives:**
- Professional, consistent design system
- Follow Shneiderman's 8 Golden Rules
- Responsive, accessible, polished UI

**Shneiderman's 8 Rules Implementation:**
1. ✅ **Consistency**: Unified color scheme, button styles, spacing
2. ✅ **Shortcuts**: Keyboard navigation, quick filters
3. ✅ **Feedback**: Loading states, success messages, hover effects
4. ✅ **Closure**: Clear task completion (search → results → detail)
5. ✅ **Error Handling**: Friendly error messages, empty states
6. ✅ **Reversibility**: Clear filters, back navigation
7. ✅ **Control**: User-driven interactions, no auto-play
8. ✅ **Memory Load**: Visible state, breadcrumbs, clear labels

**Features:**
- [ ] Consistent design system (colors, typography, spacing)
- [ ] Reusable CSS components (buttons, cards, inputs, badges)
- [ ] Responsive grid layouts (mobile, tablet, desktop)
- [ ] Smooth animations and transitions
- [ ] Loading states (skeletons, spinners)
- [ ] Empty states ("No results found")
- [ ] Error states with helpful messages
- [ ] Success feedback (toast notifications)
- [ ] Accessibility (ARIA labels, keyboard navigation, focus states)
- [ ] Mobile-first responsive design
- [ ] Dark mode support

**Key Components:**
```css
/* Button Component */
.btn-primary {
  background: var(--gradient-cta);
  color: white;
  padding: 12px 24px;
  border-radius: 12px;
  transition: transform 0.2s;
}
.btn-primary:hover {
  transform: scale(1.05);
}

/* Card Component */
.card {
  background: var(--bg-card);
  border-radius: 16px;
  padding: var(--space-lg);
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
  transition: transform 0.3s, box-shadow 0.3s;
}
.card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px rgba(0,0,0,0.15);
}
```

**Pages to Polish:**
- Landing page (hero section)
- Artist listing page
- Artist detail page
- Search results page
- Globe visualization page
- 404/Error pages

---

### Week 3: Docker Production + Deployment

#### **Days 15-17: Production Containerization** 🐳

**Features:**
- [ ] Multi-stage Dockerfile (dev + production)
- [ ] Docker Compose production config
- [ ] Environment variable management (dev vs prod)
- [ ] Health checks
- [ ] Log management
- [ ] Database backup strategy
- [ ] Security hardening (non-root user, minimal base image)

**Production Dockerfile:**
```dockerfile
# Production stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/web ./web

EXPOSE 8080

CMD ["./main"]
```

**Production docker-compose.yml:**
```yaml
version: '3.8'

services:
  postgres:
    image: postgres:16-alpine
    restart: always
    environment:
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password
    secrets:
      - db_password
    volumes:
      - prod_data:/var/lib/postgresql/data

  app:
    image: groupie-tracker:latest
    restart: always
    ports:
      - "8080:8080"
    environment:
      - ENV=production
    depends_on:
      - postgres

secrets:
  db_password:
    file: ./secrets/db_password.txt

volumes:
  prod_data:
```

---

#### **Days 18-19: Deployment** 🚀

**Platform**: Render / Railway / Fly.io

**Features:**
- [ ] Deploy PostgreSQL database
- [ ] Deploy Go application
- [ ] Environment variables configuration
- [ ] Domain setup (optional: custom domain)
- [ ] SSL certificates (automatic on Render)
- [ ] Health check endpoint (`GET /health`)
- [ ] Monitoring setup
- [ ] Test production deployment

**Render Setup:**
1. Create PostgreSQL database
2. Create Web Service (Docker)
3. Connect to GitHub repo
4. Set environment variables
5. Deploy and test

**Health Check Endpoint:**
```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
        "version": "1.0.0",
    })
}
```

---

#### **Days 20-21: Documentation & Demo** 📝

**Features:**
- [ ] Comprehensive README.md
  - Project description
  - Features list
  - Tech stack
  - Setup instructions
  - API documentation
  - Screenshots
  - Live demo link
- [ ] Code comments and documentation
- [ ] Demo video or GIF
- [ ] Architecture diagram
- [ ] Database schema documentation
- [ ] Deployment guide
- [ ] Contribution guide (if open source)

**README Structure:**
```markdown
# 🎵 Groupie Tracker v2

[Hero screenshot]

## 🎯 Overview
[Brief description]

## ✨ Features
- Search with type-ahead suggestions
- Interactive 3D concert globe
- Advanced filtering
- Responsive design

## 🛠 Tech Stack
- Go 1.22+
- PostgreSQL 16
- HTMX
- Globe.gl (Three.js)
- Docker

## 🚀 Quick Start
[Setup instructions]

## 📸 Screenshots
[Multiple screenshots]

## 🌐 Live Demo
[Link to deployed app]

## 📄 License
MIT
```

---

## 🔹 PHASE 2: Authentication + Spotify Integration (Weeks 4-6)

### 2.1: Multi-Provider OAuth

**Objectives:**
- Implement OAuth 2.0 authentication
- Support multiple login providers
- User session management

**OAuth Providers:**
- [ ] Spotify (music data + authentication)
- [ ] Google (universal login)
- [ ] Apple Sign-In (iOS users, iCloud)
- [ ] GitHub (developer-friendly, optional)

**Features:**
- [ ] OAuth configuration for all providers
- [ ] "Sign Up" and "Login" pages
- [ ] Social login buttons
- [ ] OAuth callback handlers
- [ ] Token management (access + refresh tokens)
- [ ] User profile creation/update
- [ ] Session cookies or JWT tokens
- [ ] Logout functionality
- [ ] Account linking (connect multiple providers)

**OAuth Flow:**
```
1. User clicks "Login with Spotify"
2. Redirect to provider authorization page
3. User approves permissions
4. Provider redirects back with authorization code
5. Exchange code for access token + refresh token
6. Fetch user profile from provider
7. Create/update user in database
8. Start session (cookie or JWT)
9. Redirect to dashboard
```

**Database Schema:**
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    display_name VARCHAR(255),
    profile_image VARCHAR(500),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE oauth_accounts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL, -- 'spotify', 'google', 'apple', 'github'
    provider_user_id VARCHAR(255) NOT NULL,
    access_token TEXT,
    refresh_token TEXT,
    token_expiry TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(provider, provider_user_id)
);

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    session_token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**Learning Checkpoint:**
- ✅ OAuth 2.0 flow deep dive (1 hour)
- ✅ Token management best practices (30 mins)
- ✅ Session vs JWT authentication (30 mins)
- ✅ golang.org/x/oauth2 library (45 mins)

---

### 2.2: Spotify API Integration

**Objectives:**
- Fetch rich artist data from Spotify
- Display top tracks with playback
- Enhance artist pages with Spotify data

**Features:**
- [ ] Spotify API service layer
- [ ] Fetch artist details (bio, images, genres, popularity)
- [ ] Fetch top tracks with preview URLs
- [ ] Fetch albums with release dates and cover art
- [ ] Fetch related artists
- [ ] Spotify embed player for 30-second previews
- [ ] Artist popularity and follower metrics
- [ ] Genre tags and categorization
- [ ] Artist recommendations

**Hybrid Data Strategy:**
```
Artist Detail Page:
├── Basic info from 01 API
│   ├── Concert dates
│   ├── Concert locations
│   └── Relations
└── Enhanced info from Spotify API
    ├── Biography
    ├── Top 5 tracks (playable previews)
    ├── Albums (grid with covers)
    ├── Popularity score
    ├── Genres
    ├── Follower count
    └── Related artists (carousel)
```

**Spotify API Endpoints to Use:**
- `GET /v1/artists/{id}` - Artist details
- `GET /v1/artists/{id}/top-tracks` - Top tracks
- `GET /v1/artists/{id}/albums` - Albums
- `GET /v1/artists/{id}/related-artists` - Related artists
- `GET /v1/search` - Search artists

**Database Schema Addition:**
```sql
CREATE TABLE spotify_artists (
    id SERIAL PRIMARY KEY,
    artist_id INTEGER REFERENCES artists(id),
    spotify_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
    popularity INTEGER,
    genres TEXT[],
    followers INTEGER,
    image_url VARCHAR(500),
    spotify_data JSONB,
    last_updated TIMESTAMP DEFAULT NOW()
);

CREATE TABLE spotify_tracks (
    id SERIAL PRIMARY KEY,
    spotify_artist_id INTEGER REFERENCES spotify_artists(id),
    spotify_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
    preview_url VARCHAR(500),
    duration_ms INTEGER,
    track_data JSONB
);
```

**Spotify Embed Player:**
```html
<iframe 
  src="https://open.spotify.com/embed/track/{track_id}" 
  width="100%" 
  height="80" 
  frameborder="0" 
  allow="encrypted-media">
</iframe>
```

---

### 2.3: User Features & Personalization

**Objectives:**
- Personal music library
- Follow/save functionality
- User dashboard

**Features:**
- [ ] Follow/unfollow artists
- [ ] Save concerts to personal calendar
- [ ] User dashboard (personalized homepage)
- [ ] "My Library" page (followed artists)
- [ ] "Saved Concerts" list
- [ ] Search history
- [ ] Recently viewed artists
- [ ] Notification preferences
- [ ] Email notifications toggle
- [ ] User settings page
- [ ] Profile customization (avatar, display name)

**User Dashboard:**
```
Dashboard Layout:
├── Welcome section ("Welcome back, {name}!")
├── Upcoming concerts from followed artists
├── Recently viewed artists
├── Recommended artists (based on follows)
├── Concert alerts (new concerts added)
└── Quick stats (X artists followed, Y concerts saved)
```

**Database Schema:**
```sql
CREATE TABLE user_follows (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    artist_id INTEGER REFERENCES artists(id) ON DELETE CASCADE,
    followed_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, artist_id)
);

CREATE TABLE user_saved_concerts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    concert_id INTEGER REFERENCES concerts(id) ON DELETE CASCADE,
    saved_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, concert_id)
);

CREATE TABLE user_search_history (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    query VARCHAR(255),
    searched_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_preferences (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE UNIQUE,
    email_notifications BOOLEAN DEFAULT TRUE,
    concert_alerts BOOLEAN DEFAULT TRUE,
    newsletter BOOLEAN DEFAULT FALSE,
    theme VARCHAR(20) DEFAULT 'dark' -- 'light' or 'dark'
);
```

---

## 🔹 PHASE 3: Concert APIs + Live Data (Weeks 7-8)

### 3.1: Ticketmaster & Songkick Integration

**Objectives:**
- Real, live concert data
- Ticket purchase links
- Expanded concert database

**Features:**
- [ ] Ticketmaster API integration
- [ ] Songkick API integration (if available)
- [ ] Fetch upcoming concerts by artist
- [ ] Fetch concerts by location/city
- [ ] Ticket pricing information
- [ ] Venue details (capacity, address, type)
- [ ] Purchase links (redirect to ticketing sites)
- [ ] Concert recommendations based on location
- [ ] On-sale date tracking
- [ ] Sold-out status indicators

**API Strategy:**
```
Concert Data Sources:
├── 01 Groupie Tracker API (historical, limited dataset)
├── Ticketmaster API (major venues, official ticketing)
└── Songkick API (indie/smaller venues, comprehensive)

Priority:
1. Fetch from Ticketmaster (official, best data)
2. Supplement with Songkick (broader coverage)
3. Keep 01 API as fallback/historical data
```

**Ticketmaster API Endpoints:**
- `GET /discovery/v2/events` - Search events
- `GET /discovery/v2/events/{id}` - Event details
- `GET /discovery/v2/attractions` - Artist info

**Database Schema Addition:**
```sql
CREATE TABLE external_concerts (
    id SERIAL PRIMARY KEY,
    source VARCHAR(50), -- 'ticketmaster', 'songkick', '01api'
    external_id VARCHAR(255),
    artist_id INTEGER REFERENCES artists(id),
    name VARCHAR(255),
    venue_name VARCHAR(255),
    venue_city VARCHAR(255),
    venue_country VARCHAR(255),
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    event_date TIMESTAMP,
    on_sale_date TIMESTAMP,
    ticket_url VARCHAR(500),
    min_price DECIMAL(10, 2),
    max_price DECIMAL(10, 2),
    currency VARCHAR(3),
    status VARCHAR(50), -- 'onsale', 'offsale', 'canceled', 'postponed'
    concert_data JSONB,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(source, external_id)
);
```

---

### 3.2: Enhanced Globe with Live Concerts

**Features:**
- [ ] Display live concert data on 3D globe
- [ ] Color-code markers by ticket availability
  - Red: Sold out
  - Orange: Few tickets left
  - Green: Available
  - Gray: Past concerts
- [ ] Size markers by venue capacity (small, medium, large)
- [ ] Animated marker additions (new concerts pulse)
- [ ] Heat map overlay (concert density visualization)
- [ ] Timeline scrubber (scroll through past → future)
- [ ] Filter by date range on globe
- [ ] Filter by multiple artists simultaneously
- [ ] Cluster markers for same city (expand on click)

**Globe Enhancements:**
```javascript
.pointColor(d => {
  if (d.status === 'soldout') return '#FF3B30';
  if (d.ticketsRemaining < 100) return '#FF9500';
  if (d.isPast) return '#86868B';
  return '#34C759';
})
.pointRadius(d => {
  if (d.venueCapacity > 50000) return 1.2;
  if (d.venueCapacity > 10000) return 0.8;
  return 0.5;
})
```

---

### 3.3: Notifications & Alerts

**Features:**
- [ ] Email notification system
- [ ] Concert alerts for followed artists
- [ ] In-app notification center
- [ ] Price drop alerts
- [ ] On-sale date reminders
- [ ] New concert announcements
- [ ] Newsletter (weekly digest)

**Notification Types:**
```
1. New concert announced for followed artist
2. Concert tickets going on sale soon (3 days before)
3. Concert happening soon (1 week before)
4. Price drop for saved concert
5. Concert sold out (remove from saved)
6. Concert canceled/postponed
```

**Email Templates:**
```html
<!-- New concert alert -->
<h2>🎵 New Concert Alert!</h2>
<p>{Artist Name} is coming to {City}!</p>
<p>Date: {Date}</p>
<p>Venue: {Venue}</p>
<a href="{concert_url}">Get Tickets</a>
```

---

## 🔹 PHASE 4: AI Personalization (Weeks 9-10)

### 4.1: OpenAI Integration

**Objectives:**
- AI chatbot for music insights
- Personalized recommendations
- Artist trivia and information

**Features:**
- [ ] AI chatbot interface (floating bubble + expanded view)
- [ ] "Tell me about {artist}" queries
- [ ] "Recommend artists similar to {artist}"
- [ ] Mood-based suggestions ("Suggest chill vibes artists")
- [ ] Genre exploration ("Explain the history of jazz")
- [ ] Concert recommendations based on taste
- [ ] Artist comparison ("Compare {artist1} vs {artist2}")
- [ ] Fun facts and trivia about artists

**OpenAI Prompts:**
```go
// Artist recommendation
prompt := fmt.Sprintf(
    `User follows these artists: %s.
    Recommend 5 similar artists with brief explanations.
    Format as JSON: [{name, reason, genres}]`,
    strings.Join(followedArtists, ", "),
)

// Artist insights
prompt := fmt.Sprintf(
    `Tell me an interesting fact about %s in 2-3 sentences.`,
    artistName,
)
```

**AI Chat Interface:**
- Draggable floating button (bottom-right)
- Expands to chat window
- Quick action chips ("Tell me about...", "Recommend similar")
- Chat history persistence
- Typing indicator
- Markdown support for responses

---

### 4.2: Music Profile Visualization

**Objectives:**
- Visualize user's music taste
- Interactive personality chart
- Shareable music DNA

**Features:**
- [ ] Hexagonal music personality radar chart
- [ ] Top genres breakdown (pie chart)
- [ ] Artist diversity score
- [ ] Listening habits timeline
- [ ] Concert attendance map
- [ ] Music personality type (AI-generated)
  - "Eclectic Explorer"
  - "Genre Loyalist"
  - "Live Music Fanatic"
- [ ] Share profile feature (image + link)
- [ ] Compare profiles with friends (Phase 5)

**Hexagonal Radar Chart:**
```javascript
// D3.js or Chart.js radar chart
const genres = ['Rock', 'Pop', 'Jazz', 'Hip-Hop', 'Electronic', 'Classical'];
const scores = [85, 60, 40, 75, 55, 30]; // Based on followed artists

// Hexagon visualization
drawHexagonalChart(genres, scores);
```

**Profile Stats:**
- Total artists followed
- Genres explored
- Concerts saved
- Most followed genre
- Discovery rate (new artists per month)
- Geographic diversity (concerts in X countries)

---

## 🔹 PHASE 5: Community & Indie Features (Weeks 11-13)

### 5.1: Indie Artist Platform

**Features:**
- [ ] Indie artist registration portal
- [ ] Submit local gigs/concerts
- [ ] Artist verification system
- [ ] Gig RSVP system
- [ ] Local discovery map (indie venues)
- [ ] Fan messaging (artist ↔ fans)
- [ ] Artist analytics dashboard
- [ ] Promote small shows

**Indie Artist Features:**
```
Artist Dashboard:
├── Profile management
├── Submit gig (date, venue, ticket link)
├── View RSVPs and attendees
├── Send announcements to followers
├── Analytics (views, saves, RSVPs)
└── Connect social media accounts
```

---

### 5.2: Social Features

**Features:**
- [ ] User-to-user artist recommendations
- [ ] Concert attendance planning ("Who's going?")
- [ ] Playlist generator (Spotify integration)
- [ ] Event comments and reviews
- [ ] Shareable concert memories (photos, reviews)
- [ ] Follow other users
- [ ] Activity feed (friends' follows, concert saves)
- [ ] Group concert planning

---

## 🔹 PHASE 6: Polish & Production (Weeks 14-16)

### 6.1: Performance Optimization

**Features:**
- [ ] Database query optimization (indexes, explain analyze)
- [ ] Caching layer (Redis optional, or in-memory)
- [ ] Go concurrency for API calls (goroutines, channels)
- [ ] Image optimization (WebP, lazy loading)
- [ ] Code splitting and minification
- [ ] CDN for static assets
- [ ] Database connection pooling
- [ ] Rate limiting for API endpoints
- [ ] Gzip compression

**Go Concurrency Example:**
```go
// Parallel API calls
var wg sync.WaitGroup
artistChan := make(chan SpotifyArtist)
tracksChan := make(chan []Track)

wg.Add(2)
go fetchArtist(artistID, artistChan, &wg)
go fetchTopTracks(artistID, tracksChan, &wg)

wg.Wait()
artist := <-artistChan
tracks := <-tracksChan
```

---

### 6.2: Production Deployment

**Features:**
- [ ] Fully Dockerized (multi-container production setup)
- [ ] CI/CD pipeline (GitHub Actions)
  - Automated testing
  - Build Docker image
  - Deploy to production
- [ ] Environment management (dev, staging, production)
- [ ] Secrets management (GitHub Secrets, environment variables)
- [ ] Monitoring and logging (structured logging)
- [ ] Error tracking (Sentry optional)
- [ ] Database backup automation
- [ ] SSL certificates (automatic with Render)
- [ ] Custom domain configuration
- [ ] Health checks and uptime monitoring

**GitHub Actions Workflow:**
```yaml
name: Deploy to Production

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Build Docker image
        run: docker build -t groupie-tracker .
      - name: Deploy to Render
        run: # Deploy script
```

---

### 6.3: Documentation & Portfolio Presentation

**Features:**
- [ ] Comprehensive README with screenshots
- [ ] Architecture diagrams (system design, database schema)
- [ ] API documentation (if exposing APIs)
- [ ] Setup and installation guide
- [ ] Developer contribution guide
- [ ] Demo video (2-3 minutes)
- [ ] Live demo link prominently displayed
- [ ] Portfolio case study write-up
  - Problem statement
  - Solution approach
  - Technical challenges and solutions
  - Results and metrics
- [ ] Code comments and documentation
- [ ] Unit tests for critical functions
- [ ] Integration tests

**Portfolio Highlights:**
- 🎯 Full-stack application (Go + PostgreSQL + HTMX)
- 🌍 3D visualization (Globe.gl)
- 🔐 Multi-provider OAuth (Spotify, Google, Apple)
- 🤖 AI integration (OpenAI)
- 📡 Multiple API integrations (Spotify, Ticketmaster, 01 API)
- 🐳 Docker containerization
- 🚀 Production deployment
- 📱 Responsive design
- ♿ Accessibility compliant
- 🧪 Test coverage

---

## 📊 Complete Technology Stack

### Backend
- **Language**: Go 1.22+
- **Router**: `net/http` (standard library)
- **Database Driver**: `github.com/lib/pq` (PostgreSQL)
- **OAuth**: `golang.org/x/oauth2`
- **Environment**: `github.com/joho/godotenv`
- **Hot Reload**: Air (development)
- **Testing**: Go standard testing package

### Database
- **Primary**: PostgreSQL 16
- **Caching**: In-memory or Redis (optional, Phase 6)
- **Migrations**: Custom Go scripts

### Frontend
- **Base**: HTML5, CSS3, Vanilla JavaScript
- **Interactivity**: HTMX (progressive enhancement)
- **3D Visualization**: Globe.gl (Three.js wrapper)
- **Charts**: Chart.js or D3.js (Phase 4)
- **Icons**: Feather Icons or Lucide Icons
- **Fonts**: SF Pro Display (Apple system font fallback)

### External APIs
- **01 Founders API**: Groupie Tracker API (artist/concert data)
- **Spotify API**: Artist data, tracks, albums, OAuth
- **Ticketmaster API**: Live concert data, ticketing
- **Songkick API**: Indie/smaller venue concerts
- **OpenAI API**: AI recommendations and insights
- **Geocoding**: Nominatim (free) or Google Geocoding

### Infrastructure
- **Containerization**: Docker + Docker Compose
- **Deployment**: Render / Railway / Fly.io
- **CI/CD**: GitHub Actions
- **Monitoring**: Logs, health checks
- **SSL**: Automatic (platform-provided)

### Development Tools
- **IDE**: GoLand / VSCode with Go extension
- **Version Control**: Git + GitHub
- **Environment**: WSL Ubuntu / macOS / Linux
- **API Testing**: Postman / Thunder Client
- **Database Client**: pgAdmin / DBeaver

---

## 🎯 Success Criteria

### 01 Founders Requirements ✅
- ✅ Core brief (README.md) - Artist/concert display
- ✅ Search bar brief - Type-ahead suggestions
- ✅ Filters brief - Range + checkbox filtering
- ✅ Geolocalization brief - 3D globe with markers
- ✅ Visualizations brief - Professional UI, Shneiderman's rules
- ✅ Standard Go library focus (minimal external packages)
- ✅ Proper error handling
- ✅ Clean code practices

### Portfolio Quality ✅
- ✅ Professional, polished UI (Apple Music aesthetic)
- ✅ Fully responsive (mobile, tablet, desktop)
- ✅ Fast performance (< 2s page loads)
- ✅ No bugs or crashes
- ✅ Comprehensive documentation
- ✅ Live demo deployed
- ✅ Unique features (3D globe, AI chat)

### Technical Demonstration ✅
- ✅ Multiple API integrations (5+ APIs)
- ✅ Database design (normalized schema)
- ✅ Authentication/authorization (OAuth 2.0)
- ✅ Real-time interactions (HTMX)
- ✅ 3D visualization (Globe.gl)
- ✅ Docker deployment
- ✅ CI/CD pipeline
- ✅ Security best practices

---

## 📅 Development Timeline

| Phase | Duration | Key Deliverables |
|-------|----------|------------------|
| **Phase 1** (MVP) | 2-3 weeks | Search, Globe, Filters, Docker, Deployed |
| **Phase 2** (Auth + Spotify) | 2-3 weeks | OAuth, User profiles, Spotify integration |
| **Phase 3** (Concert APIs) | 2 weeks | Ticketmaster, Live data, Notifications |
| **Phase 4** (AI) | 2 weeks | OpenAI chatbot, Music profile viz |
| **Phase 5** (Community) | 2-3 weeks | Indie platform, Social features |
| **Phase 6** (Polish) | 1-2 weeks | Optimization, Documentation |
| **Total** | **12-16 weeks** | **Production-ready application** |

**MVP Milestone**: End of Week 3 ✅  
**Feature-Complete**: End of Week 13  
**Production Launch**: End of Week 16

---

## 🚀 Getting Started

### Prerequisites
- Go 1.22+
- Docker Desktop
- Git
- GoLand or VSCode
- PostgreSQL client (optional, for manual queries)

### Initial Setup
```bash
# Clone repository
git clone https://github.com/IbsYoussef/Groupie-Tracker.git
cd Groupie-Tracker

# Switch to v2 development branch
git checkout v2-development

# Copy environment template
cp .env.example .env

# Edit .env with your API keys
nano .env

# Start development with Docker
make dev

# Access app
open http://localhost:8080
```

---

## 📝 Learning Checkpoints

| Checkpoint | Topic | Duration | Phase |
|------------|-------|----------|-------|
| 1 | PostgreSQL basics | 20 mins | 1.1 |
| 2 | Docker Compose | 30 mins | 1.1 |
| 3 | HTMX core concepts | 45 mins | 1.2 |
| 4 | HTMX search patterns | 1 hour | 1.2 |
| 5 | Three.js & Globe.gl | 1.5 hours | 1.3 |
| 6 | Geocoding APIs | 30 mins | 1.3 |
| 7 | HTMX forms & filters | 1 hour | 1.5 |
| 8 | OAuth 2.0 flow | 1 hour | 2.1 |
| 9 | Spotify API | 45 mins | 2.2 |
| 10 | OpenAI integration | 30 mins | 4.1 |

**Total Learning Time**: ~8-10 hours spread across phases

---

## 🎨 Design Principles

**Consistency**: Unified color palette, typography, spacing throughout  
**Simplicity**: Clean layouts, minimal cognitive load  
**Responsiveness**: Mobile-first, works beautifully on all devices  
**Accessibility**: WCAG 2.1 AA compliant, keyboard navigation  
**Performance**: Fast page loads, smooth animations  
**Delight**: Micro-interactions, thoughtful transitions  

---

## 📊 Metrics & KPIs (Optional)

**Technical Metrics:**
- Page load time: < 2 seconds
- Lighthouse score: > 90
- Test coverage: > 70%
- API response time: < 500ms

**User Metrics (Post-Launch):**
- Artists searched per session
- Concerts viewed
- Globe interactions
- User retention

---

## 🔄 Maintenance & Updates

**Regular Updates:**
- Weekly: Concert data refresh (Ticketmaster API)
- Monthly: Security patches, dependency updates
- Quarterly: Feature additions based on feedback

**Monitoring:**
- Uptime monitoring
- Error tracking
- Performance metrics
- User analytics

---

**Last Updated**: January 2025  
**Version**: 2.0.0  
**Status**: In Development  
**Maintainer**: Ibrahim Youssef (@IbsYoussef)

---

## 🙋 Questions or Feedback?

- GitHub Issues: [github.com/IbsYoussef/Groupie-Tracker/issues](https://github.com/IbsYoussef/Groupie-Tracker/issues)
- Portfolio: [your-portfolio.com]
