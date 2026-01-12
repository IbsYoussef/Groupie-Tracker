.PHONY: help setup install-tools deps env dev build run test clean

# Default target
help:
	@echo "🎵 Groupie Tracker v2 - Development Commands"
	@echo ""
	@echo "Setup (run once on new machine):"
	@echo "  make setup        - Complete setup for new machine"
	@echo "  make install-tools - Install Air and other dev tools"
	@echo "  make deps         - Download Go dependencies"
	@echo "  make env          - Copy .env.example to .env"
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

# Complete setup for new machine
setup: install-tools deps env
	@echo ""
	@echo "✅ Setup complete!"
	@echo ""
	@echo "📝 Next steps:"
	@echo "  1. Edit .env and add your API keys"
	@echo "  2. Run 'make dev' to start development"
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
		echo "ℹ️  Edit .env to add your API keys when ready"; \
	else \
		echo "ℹ️  .env already exists, skipping..."; \
	fi

# Start development server with hot reload
dev:
	@if ! command -v air > /dev/null 2>&1; then \
		echo "❌ Air not found. Run 'make install-tools' first"; \
		exit 1; \
	fi
	@if [ ! -f .env ]; then \
		echo "❌ .env file not found. Run 'make env' first"; \
		exit 1; \
	fi
	@echo "🚀 Starting development server..."
	air

# Run without hot reload
run:
	@if [ ! -f .env ]; then \
		echo "❌ .env file not found. Run 'make env' first"; \
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
	@echo "Checking .env file..."
	@test -f .env && echo "✅ .env exists" || echo "⚠️  .env missing (run 'make env')"
	@echo ""
	@echo "Checking dependencies..."
	@go mod verify > /dev/null 2>&1 && echo "✅ Dependencies verified" || echo "⚠️  Dependencies issue (run 'make deps')"
	@echo ""
	@echo "Checking project structure..."
	@test -f cmd/api/main.go && echo "✅ main.go exists" || echo "❌ main.go missing"
	@test -f .air.toml && echo "✅ .air.toml exists" || echo "❌ .air.toml missing"
	@echo ""
	@echo "Setup verification complete!"