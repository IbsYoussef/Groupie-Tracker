# 🎵 Groupie Tracker – Full Feature Roadmap

## 🔹 PHASE 1: MVP – Core Functionality

Your foundation. Clean, useful, portfolio-ready.

### ✅ Essential Features:

- [ ] User Authentication (Signup/Login/Logout)
- [ ] Artist Search (via Spotify API)
- [ ] Listen to Top Tracks (Spotify embed or playback)
- [ ] Concert Discovery (events via Ticketmaster/Songkick API)
- [ ] Basic User Profile (display followed/saved artists)

### 🛠 Tech Stack – MVP:

- Go (backend)
- PostgreSQL (database)
- HTML / CSS / Vanilla JS (frontend)
- Docker (containerization & deployment)

> 📌 **HTMX not yet included** — keep frontend minimal and manageable.

---

## 🔹 PHASE 2: Follow & Personalize

### 🎯 Engagement Features:

- [ ] Follow Artists
- [ ] Create a Personal Music Library
- [ ] Concert Notifications (or simple list)
- [ ] Track Listening History

### 🔧 Tech Stack – Phase 2:

- Same as MVP
- You **may optionally begin introducing HTMX here** for:
  - Progressive rendering of artist cards
  - Real-time updates of follows
  - Live search or pagination without full JS SPA setup

> ✅ This is the best time to start exploring HTMX — not critical, but nice to enhance UX if you're ready.

---

## 🔹 PHASE 3: AI-Driven Personalization

### 🧠 AI Features:

- [ ] AI-generated Artist Recommendations
- [ ] Mood-based Tagging (e.g., "Chill", "Workout")
- [ ] Personalized Suggestions Dashboard

### 🧠 Tech Stack – Phase 3:

- Add OpenAI or other ML backend service
- Possibly store embeddings or tag mappings
- Optional: HTMX for AI-enhanced UI interactions

---

## 🔹 PHASE 4: Indie & Community Tools

### 🚀 Empower Indie Artists:

- [ ] Indie Band Submission Portal
- [ ] RSVP System for Small Gigs
- [ ] Local Discovery Map

---

## 🔹 PHASE 5: Music Profile Visuals & UX

### 🎨 Visual Flair:

- [ ] Hexagonal SVG Chart of Music Personality
- [ ] Animated Music Graphs
- [ ] Custom Theming (by genre/mood)

### 🎨 Tech:

- D3.js or SVG rendering tools
- HTMX or Alpine.js for lightweight interactivity

---

## 🔹 PHASE 6: Stretch & Social

### 💡 Bonus Features:

- [ ] Shareable Music Profiles
- [ ] User-to-User Recs
- [ ] Playlist Generator
- [ ] Commenting on Events/Artists

---

## 🧭 Deployment Checklist:

- [ ] Fully Dockerized
- [ ] Hosted on Render/Railway/Fly.io
- [ ] Polished README with demo
- [ ] Responsive design
