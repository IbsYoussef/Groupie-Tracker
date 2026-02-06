-- Migration: 001_init.sql
-- Initial database schema for Groupie Tracker v2
-- Creates users and sessions tables

-- Enable UUID extension (needed for gen_random_uuid())
    CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ========================================
-- USERS TABLE
-- ========================================
-- Stores all user accounts (email/password and OAuth users)

CREATE TABLE IF NOT EXISTS users (
    -- Primary key: Auto-generated UUID
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    -- Username: Unique identifier for the user
    username VARCHAR(50) UNIQUE NOT NULL,

    -- Email: Used for login and notifications
    email VARCHAR(255) UNIQUE NOT NULL,

    -- Password hash: NULL for OAuth users (Spotify, Google, Apple)
    password_hash TEXT,

    -- OAuth provider: 'spotify', 'google', 'apple', or NULL for email/password
    oauth_provider VARCHAR(50),

    -- OAuth ID: User's ID from the OAuth provider
    oauth_id TEXT,

    -- Timestamps
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- Constraints
    CHECK (
        (oauth_provider IS NULL AND oauth_id IS NULL) OR
        (oauth_provider IS NOT NULL AND oauth_id IS NOT NULL)
    )
);

-- ========================================
-- SESSIONS TABLE
-- ========================================
-- Stores active user sessions (login tokens)

CREATE TABLE IF NOT EXISTS sessions (
    -- Primary key: Auto-generated UUID
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    -- Foreign key: References users table
    -- ON DELETE CASCADE: If a user is deleted, their sessions auto-delete
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    -- Session token: Random string stored in cookie/header
    token TEXT UNIQUE NOT NULL,

    -- Expiration: When this session becomes invalid
    expires_at TIMESTAMP NOT NULL,

    -- Created timestamp
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ========================================
-- INDEXES
-- ========================================
-- Speed up common queries

-- Index on username for profile lookups
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Index on OAuth fields for OAuth login
CREATE INDEX IF NOT EXISTS idx_users_oauth ON users(oauth_provider, oauth_id);

-- Index on session token for authentication checks
CREATE INDEX IF NOT EXISTS idx_sessions_token ON sessions(token);

-- Index on user_id for fetching all user sessions
CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions(user_id);

-- Index on expires_at for cleanup queries
CREATE INDEX IF NOT EXISTS idx_sessions_expires_at ON sessions(expires_at);

-- ========================================
-- TRIGGERS
-- ========================================
-- Auto-update updated_at timestamp

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
