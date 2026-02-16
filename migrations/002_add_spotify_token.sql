-- Migration 002: Add spotify_access_token column to users table
-- Required for storing the user's Spotify OAuth token after login
-- so it can be reused for API calls (e.g. fetching top artists)

ALTER TABLE users ADD COLUMN IF NOT EXISTS spotify_access_token TEXT DEFAULT '';