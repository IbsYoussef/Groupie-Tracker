.PHONY: help setup install-tools deps env gen-secret docker-start docker-stop docker-restart db-shell db-query db-users db-sessions db-tables db-reset dev run build test clean verify

# Default target
help:
	@echo "🎵 Groupie Tracker v2 - Development Commands"
	@echo ""
	@echo "Setup (run once on new machine):"
	@echo "  make setup        - Complete setup for new machine"
	@echo "  make install-tools - Install Air and other dev tools"
	@echo "  make deps         - Download Go dependencies"
	@echo "  make env          - Copy .env.example to .env"
	@echo "  make gen-secret   - Generate SESSION_SECRET in .env"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-start   - Start Docker containers with proper permissions"
	@echo "  make docker-stop    - Stop Docker containers"
	@echo "  make docker-restart - Restart Docker containers"
	@echo ""
	@echo "Database:"
	@echo "  make db-shell     - Open PostgreSQL shell"
	@echo "  make db-query     - Run SQL query (usage: make db-query Q='...')"
	@echo "  make db-users     - List all users"
	@echo "  make db-sessions  - List active sessions"
	@echo "  make db-tables    - List all tables"
	@echo "  make db-reset     - ⚠️  Reset database (deletes all data)"
	@echo ""
	@echo "Development:"
	@echo "  make dev          - Start development server with hot reload"
	@echo "  make run          - Run without hot reload"
	@echo "  make build        - Build the application"
	@echo ""
	@echo "Utilities:"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make verify       - Verify setup is correct"
	@echo ""
	@echo "💡 Tip: Stop dev server with Ctrl+C"

# Complete setup for new machine
setup: install-tools deps env
	@echo ""
	@echo "✅ Setup complete!"
	@echo ""
	@echo "📝 Next steps:"
	@echo "  1. Run 'make gen-secret' to generate SESSION_SECRET"
	@echo "  2. Run 'make docker-start' to start database"
	@echo "  3. Run 'make dev' to start development server"
	@echo "  4. Visit http://localhost:8080 and register"
	@echo ""

# Install development tools
install-tools:
	@echo "📦 Installing development tools..."
	@command -v air > /dev/null 2>&1 || \
		(echo "Installing Air..." && go install github.com/air-verse/air@latest)
	@echo "✅ Air installed"
	@echo "🔍 Checking PATH..."
	@which air > /dev/null 2>&1 && echo "✅ Air found in PATH" || \
		(echo "⚠️  Air not in PATH. Add this to your ~/.bashrc:" && \
		 echo "   export PATH=\$$PATH:\$$(go env GOPATH)/bin" && \
		 echo "   Then run: source ~/.bashrc")

# Download Go dependencies
deps:
	@echo "📥 Downloading Go dependencies..."
	@go mod download
	@go mod verify
	@echo "✅ Dependencies downloaded and verified"

# Copy .env.example to .env if it doesn't exist
env:
	@if [ ! -f .env ]; then \
		echo "📝 Creating .env from .env.example..."; \
		cp .env.example .env; \
		echo "✅ .env created"; \
		echo "ℹ️  Run 'make gen-secret' to generate SESSION_SECRET"; \
	else \
		echo "ℹ️  .env already exists, skipping..."; \
	fi

# Generate SESSION_SECRET in .env file
gen-secret:
	@if [ ! -f .env ]; then \
		echo "❌ .env file not found. Run 'make env' first"; \
		exit 1; \
	fi
	@echo "🔑 Generating SESSION_SECRET..."
	@SECRET=$$(openssl rand -hex 32); \
	if echo "$$OSTYPE" | grep -q "darwin"; then \
		sed -i '' "s/SESSION_SECRET=your_session_secret/SESSION_SECRET=$$SECRET/" .env; \
		sed -i '' "s/SESSION_SECRET=.*/SESSION_SECRET=$$SECRET/" .env; \
	else \
		sed -i "s/SESSION_SECRET=your_session_secret/SESSION_SECRET=$$SECRET/" .env; \
		sed -i "s/SESSION_SECRET=.*/SESSION_SECRET=$$SECRET/" .env; \
	fi
	@echo "✅ SESSION_SECRET generated and added to .env"

# Docker commands
docker-start:
	@echo "🐳 Starting Docker with current user permissions..."
	@UID=$$(id -u) GID=$$(id -g) docker compose up -d
	@echo "⏳ Waiting for container to start..."
	@sleep 5
	@echo "🔍 Checking for permission issues..."
	@if docker compose logs postgres 2>&1 | grep -q "Permission denied"; then \
		echo "⚠️  Permission denied detected!"; \
		echo "🔧 Fixing permissions..."; \
		make docker-stop; \
		sudo chown -R $$(id -u):$$(id -g) postgres-data/ 2>/dev/null || \
			(echo "📁 Removing postgres-data folder..." && sudo rm -rf postgres-data/); \
		echo "🔄 Restarting Docker..."; \
		UID=$$(id -u) GID=$$(id -g) docker compose up -d; \
		sleep 10; \
	fi
	@echo "✅ Checking container health..."
	@STATUS=$$(docker ps --filter "name=groupie-tracker-db" --format "{{.Status}}"); \
	if echo "$$STATUS" | grep -qE "Up|healthy"; then \
		echo "✅ Docker containers started successfully ($$STATUS)"; \
	else \
		echo "⚠️  Container status: $$STATUS"; \
		echo "💡 Run 'docker ps' to check or see TROUBLESHOOTING.md"; \
	fi

docker-stop:
	@echo "🛑 Stopping Docker containers..."
	@docker compose down
	@echo "✅ Docker containers stopped"

docker-restart: docker-stop docker-start
	@echo "🔄 Docker containers restarted"

# Database commands
db-shell:
	@echo "🗄️  Opening PostgreSQL shell..."
	@echo "💡 Tip: Use \dt to list tables, \q to quit"
	@docker exec -it groupie-tracker-db psql -U groupie_user -d groupie_tracker

db-query:
	@if [ -z "$(Q)" ]; then \
		echo "❌ No query provided"; \
		echo "Usage: make db-query Q='SELECT * FROM users;'"; \
		exit 1; \
	fi
	@echo "🔍 Running query..."
	@docker exec -it groupie-tracker-db psql -U groupie_user -d groupie_tracker -c "$(Q)"

db-users:
	@echo "👥 Current users:"
	@docker exec -it groupie-tracker-db psql -U groupie_user -d groupie_tracker -c "SELECT username, email, created_at FROM users ORDER BY created_at DESC;"

db-sessions:
	@echo "🔐 Active sessions:"
	@docker exec -it groupie-tracker-db psql -U groupie_user -d groupie_tracker -c "SELECT user_id, expires_at, created_at FROM sessions WHERE expires_at > NOW() ORDER BY created_at DESC;"

db-tables:
	@echo "📋 Database tables:"
	@docker exec -it groupie-tracker-db psql -U groupie_user -d groupie_tracker -c "\dt"

db-reset:
	@echo "⚠️  WARNING: This will delete all data!"
	@echo "Press Ctrl+C to cancel, or Enter to continue..."
	@read confirm
	@echo "🗑️  Resetting database..."
	@make docker-stop
	@sudo rm -rf postgres-data/
	@make docker-start
	@echo "✅ Database reset complete. Register a new user to start fresh."

# Start development server with hot reload
dev:
	@if ! command -v air > /dev/null 2>&1; then \
		echo "❌ Air not found. Run 'make install-tools' first"; \
		exit 1; \
	fi
	@if [ ! -f .env ]; then \
		echo "❌ .env file not found. Run 'make setup' first"; \
		exit 1; \
	fi
	@echo "🚀 Starting development server..."
	air

# Run without hot reload
run:
	@if [ ! -f .env ]; then \
		echo "❌ .env file not found. Run 'make setup' first"; \
		exit 1; \
	fi
	@echo "🚀 Starting server..."
	go run cmd/api/main.go

# Build the application
build:
	@echo "🔨 Building application..."
	@mkdir -p bin
	@go build -o bin/groupie-tracker ./cmd/api
	@echo "✅ Build complete: bin/groupie-tracker"

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf bin/ tmp/
	@echo "✅ Clean complete"

# Verify setup
verify:
	@echo "🔍 Verifying setup..."
	@echo ""
	@echo "Checking Go installation..."
	@go version || (echo "❌ Go not installed" && exit 1)
	@echo "✅ Go installed"
	@echo ""
	@echo "Checking Air installation..."
	@which air > /dev/null 2>&1 && echo "✅ Air installed" || echo "⚠️  Air not found (run 'make install-tools')"
	@echo ""
	@echo "Checking Docker..."
	@docker --version > /dev/null 2>&1 && echo "✅ Docker installed" || echo "⚠️  Docker not found"
	@docker ps > /dev/null 2>&1 && echo "✅ Docker running" || echo "⚠️  Docker not running"
	@echo ""
	@echo "Checking .env file..."
	@test -f .env && echo "✅ .env exists" || echo "⚠️  .env missing (run 'make env')"
	@if [ -f .env ]; then \
		grep -q "SESSION_SECRET=your_session_secret" .env && \
			echo "⚠️  SESSION_SECRET not generated (run 'make gen-secret')" || \
			echo "✅ SESSION_SECRET is set"; \
	fi
	@echo ""
	@echo "Checking dependencies..."
	@go mod verify > /dev/null 2>&1 && echo "✅ Dependencies verified" || echo "⚠️  Dependencies issue (run 'make deps')"
	@echo ""
	@echo "Checking project structure..."
	@test -f cmd/api/main.go && echo "✅ main.go exists" || echo "❌ main.go missing"
	@test -f .air.toml && echo "✅ .air.toml exists" || echo "❌ .air.toml missing"
	@test -f docker-compose.yml && echo "✅ docker-compose.yml exists" || echo "❌ docker-compose.yml missing"
	@echo ""
	@echo "Setup verification complete!"