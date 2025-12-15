# FSS Demo - Modern Stock Dashboard

A full-stack financial stock dashboard application demonstrating modern development practices using **SvelteKit** (frontend), **Go with Fibers** (backend), and **Docker** containerization.

## 🎯 Project Overview

This project showcases a complete, production-ready web application built with:

- **Frontend**: SvelteKit with Server-Side Rendering (SSR), TailwindCSS, and Flowbite components
- **Backend**: Go with Fiber framework for high-performance REST APIs
- **Infrastructure**: Docker and Docker Compose for containerization and orchestration
- **Data Source**: Alpha Vantage API for real-time stock data

## 📋 Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    User's Browser                           │
│                  (SvelteKit Frontend)                       │
└─────────────────────────────────────────────────────────────┘
                           ↕
              HTTP Requests / JSON Responses
                           ↕
┌─────────────────────────────────────────────────────────────┐
│              Go Fibers API Backend (Port 8080)              │
│     - Stock data endpoints                                  │
│     - Data processing                                       │
│     - CORS support                                          │
└─────────────────────────────────────────────────────────────┘
                           ↕
┌─────────────────────────────────────────────────────────────┐
│          Alpha Vantage API (External Data Source)           │
└─────────────────────────────────────────────────────────────┘
```

## 🚀 Quick Start

### Prerequisites

- **Docker** and **Docker Compose** (recommended)
- OR
- **Go** 1.21+ (for backend)
- **Node.js** 18+ with **Bun** (for frontend)
- **Alpha Vantage API Key** (get free key at https://www.alphavantage.co/api/)

### Using Docker Compose (Recommended)

1. **Clone the repository:**

```bash
git clone <repository-url>
cd fss-demo
```

2. **Setup Environment Variables and Configuration:**

#### Frontend Setup

```bash
# Copy example env file
cp frontend/.env.example frontend/.env.local

# Edit the file and set your API URL
# nano frontend/.env.local
```

#### Backend Setup

```bash
# Copy example config file
cp stock-ingestor/config.example.yaml stock-ingestor/config.yaml

# Edit the config file and add your Alpha Vantage API key
# nano stock-ingestor/config.yaml
```

Update `stock-ingestor/config.yaml`:

```yaml
---
server:
  port: 8080
  api_key: your-server-api-key-here <-- Change this for api-key for frontend to backend. Can be anything.

database:
  host: tsdb
  port: 5432
  user: postgres
  password: password
  dbname: screener

alphavantage:
  apiKey: YOUR_ALPHAVANTAGE_API_KEY_HERE <-- Get Alphavantage API key from https://www.alphavantage.co/support/#api-key
  uri: https://www.alphavantage.co/query?
```

Also set environment variables for Docker:

```bash
# In compose.yml, update:
# GO_API_KEY should match the api_key in config.yaml server section
# GO_API_KEY=your-server-api-key-here
```

**Important Note:** `GO_API_KEY` environment variable in `compose.yml` must match the `api_key` in `stock-ingestor/config.yaml`:

```yaml
# config.yaml
server:
  port: 8080
  api_key: your-server-api-key-here # ← This value
```

```yaml
# compose.yml
services:
  api:
    environment:
      - GO_API_KEY=your-server-api-key-here # ← Must be the SAME as above
```

3. **Start the application:**

```bash
docker-compose up
```

3. **Access the application:**

- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

4. **Stop the application:**

```bash
docker-compose down
```

### Manual Setup (Development)

#### Backend Setup

1. **Navigate to stock-ingestor directory:**

```bash
cd stock-ingestor
```

2. **Copy and configure the configuration file:**

```bash
# Copy example config
cp config.example.yaml config.yaml

# Edit config.yaml with your Alpha Vantage API key
nano config.yaml
```

3. **Update config.yaml:**

```yaml
server:
  port: 8080
  host: "0.0.0.0"

api:
  key: "YOUR_ALPHA_VANTAGE_API_KEY" # ← Replace with your API key
  base_url: "https://www.alphavantage.co"

alphavantage:
  apiKey: "YOUR_ALPHA_VANTAGE_API_KEY" # ← Replace with your API key
```

**Where to get an API key:**

- Visit https://www.alphavantage.co/api/
- Sign up for free
- Copy your API key from the dashboard

4. **Run the backend:**

```bash
go run cmd/api/main.go
```

The backend will start on `http://localhost:8080`

#### Frontend Setup

1. **Navigate to frontend directory:**

```bash
cd frontend
```

2. **Copy and configure the environment file:**

```bash
# Copy example env file
cp .env.example .env.local

# Edit .env.local file
nano .env.local
```

3. **Update .env.local:**

```env
VITE_API_URL=http://localhost:8080         # Backend API URL
VITE_DEV_PORT=5173                         # Development port
VITE_CORS_ORIGIN=https://localhost         # CORS origin
```

4. **Install dependencies:**

```bash
bun install
# or
npm install
```

5. **Run development server:**

```bash
bun run dev
# or
npm run dev
```

The frontend will start on `http://localhost:5173`

## 📚 Available Commands

### Frontend Commands

```bash
# Start development server
bun run dev
npm run dev

# Build for production
bun run build
npm run build

# Preview production build
bun run preview
npm run preview

# Type checking
bun run check
npm run check

# Format code
bun run format
npm run format

# Lint code
bun run lint
npm run lint
```

### Backend Commands

```bash
# Run the server
go run cmd/api/main.go

# Build binary
go build -o stock-ingestor cmd/api/main.go

# Run tests
go test ./...

# Run with hot reload (requires: go install github.com/cosmtrek/air@latest)
air
```

## 🔗 API Endpoints

### Stock Data

**Get Time Series Data**

```http
GET /api/stocks/:symbol?period=daily|weekly|monthly
```

Query Parameters:

- `symbol`: Stock symbol (e.g., AAPL, GOOGL)
- `period`: Time series period - `daily`, `weekly`, or `monthly` (default: daily)

Example:

```bash
curl http://localhost:8080/api/stocks/AAPL?period=weekly
```

**Get Company Overview**

```http
GET /api/stocks/:symbol/overview
```

Example:

```bash
curl http://localhost:8080/api/stocks/AAPL/overview
```

## 📁 Project Structure

```
fss-demo/
├── frontend/                 # SvelteKit frontend application
│   ├── src/
│   │   ├── routes/          # Page routes and API routes
│   │   ├── lib/
│   │   │   ├── components/  # Reusable Svelte components
│   │   │   ├── stores/      # Svelte stores (state management)
│   │   │   └── types/       # TypeScript type definitions
│   │   └── app.html         # HTML shell
│   ├── package.json
│   ├── svelte.config.js
│   ├── vite.config.ts
│   └── tsconfig.json
│
├── stock-ingestor/           # Go Fiber backend API
│   ├── cmd/
│   │   └── api/
│   │       └── main.go      # Application entry point
│   ├── internal/
│   │   ├── application/     # Use cases and DTOs
│   │   ├── domain/          # Entities and interfaces
│   │   ├── infrastructure/  # External API adapters
│   │   └── interfaces/      # HTTP handlers
│   ├── config/              # Configuration management
│   ├── go.mod
│   ├── go.sum
│   └── README.md
│
├── docs/                     # Documentation
│   ├── presentation/
│   │   └── presentation.tex # Beamer presentation
│   └── tools.json          # Tool documentation
│
├── example-data/            # Sample JSON data files
│   ├── balance_sheet_aapl.json
│   ├── earnings_aapl.json
│   ├── time_series_daily_aapl.json
│   └── ...
│
├── scripts/                  # Installation scripts
│   ├── install_dev_macos.sh
│   ├── install_dev_linux.sh
│   └── install_dev_windows.ps1
│
├── compose.yml              # Docker Compose configuration
├── README.md                # This file
└── .gitignore
```

## 🐳 Docker Setup

### Docker Compose Services

**api** - Go Fibers backend

- Port: 8080
- Environment: development
- Auto-reload on file changes

**frontend** - SvelteKit frontend

- Port: 5173
- Depends on: api service
- Environment: development
- Auto-reload on file changes

### Build Docker Images

```bash
# Build all services
docker-compose build

# Build specific service
docker-compose build api
docker-compose build frontend
```

### Rebuild Production Images

```bash
# Build production images using Dockerfile.prod
docker build -f frontend/Dockerfile.prod -t fss-demo-frontend:latest frontend/
docker build -f stock-ingestor/Dockerfile.prod -t fss-demo-api:latest stock-ingestor/
```

## 🔧 Configuration

### Getting API Keys

**Alpha Vantage API Key (Required for Backend):**

1. Visit https://www.alphavantage.co/api/
2. Enter your email address
3. Click "GET FREE API KEY"
4. Copy your API key
5. Use it in both `config.yaml` and environment variables

### Backend Configuration (stock-ingestor)

**Step 1: Copy example config**

```bash
cd stock-ingestor
cp config.example.yaml config.yaml
```

**Step 2: Edit `stock-ingestor/config.yaml`:**

```yaml
server:
  port: 8080
  host: "0.0.0.0"

api:
  key: "YOUR_ACTUAL_API_KEY_HERE" # ← Paste your API key
  base_url: "https://www.alphavantage.co"

alphavantage:
  apiKey: "YOUR_ACTUAL_API_KEY_HERE" # ← Paste your API key here too
```

**Step 3: For Docker, set environment variable in `compose.yml`:**

```yaml
services:
  api:
    environment:
      - GO_API_KEY=YOUR_ACTUAL_API_KEY_HERE # ← Paste your API key
```

### Frontend Configuration

**Step 1: Copy example env file**

```bash
cd frontend
cp .env.example .env.local
```

**Step 2: Edit `frontend/.env.local`:**

```env
VITE_API_URL=http://localhost:8080      # Backend API URL
VITE_DEV_PORT=5173                      # Development server port
VITE_CORS_ORIGIN=https://localhost      # CORS origin
```

### Example Configuration Files

**Expected `config.example.yaml` structure:**

```yaml
server:
  port: 8080
  host: "0.0.0.0"
  environment: development

api:
  key: "" # Fill with your Alpha Vantage API key
  base_url: "https://www.alphavantage.co"

alphavantage:
  apiKey: "" # Fill with your Alpha Vantage API key
  timeout: 30
  max_requests_per_minute: 5
```

**Expected `frontend/.env.example` structure:**

```env
VITE_API_URL=http://localhost:8080
VITE_DEV_PORT=5173
VITE_CORS_ORIGIN=https://localhost
NODE_ENV=development
```

### Verifying Configuration

After setup, verify everything is working:

```bash
# Test backend API
curl http://localhost:8080/api/stocks/AAPL?period=daily

# Test frontend (open in browser)
http://localhost:5173
```

## 📊 Features

### Frontend (SvelteKit)

- ✅ Server-Side Rendering (SSR) for fast initial load
- ✅ Responsive UI with TailwindCSS
- ✅ Stock symbol search and selection
- ✅ Multiple chart visualizations:
  - Candlestick charts
  - Volume bar charts
  - Time series data
- ✅ Financial statement views:
  - Income statements
  - Balance sheets
  - Cash flow statements
- ✅ Company overview information
- ✅ Dark mode support (via theme store)

### Backend (Go + Fibers)

- ✅ Clean Architecture implementation
- ✅ RESTful API endpoints
- ✅ Alpha Vantage API integration
- ✅ CORS support
- ✅ Error handling and validation
- ✅ Configuration management
- ✅ Structured logging

## 🌟 Technology Stack

### Frontend

- **SvelteKit** - Modern web framework with SSR
- **Svelte 5** - Reactive JavaScript framework
- **TailwindCSS** - Utility-first CSS framework
- **Flowbite** - Component library
- **TypeScript** - Type-safe JavaScript
- **Vite** - Fast build tool

### Backend

- **Go** - High-performance compiled language
- **Fiber** - Fast and minimalist web framework
- **Clean Architecture** - Maintainable code structure

### Infrastructure

- **Docker** - Containerization
- **Docker Compose** - Multi-container orchestration

## 📝 Development Notes

### Hot Reloading

Both services support hot reloading:

- **Frontend**: Changes to `.svelte`, `.ts`, `.css` files trigger rebuild
- **Backend**: Changes to `.go` files trigger server restart (via Docker volumes)

### Debugging

- Frontend: Open browser DevTools (F12)
- Backend: Check Docker logs with `docker-compose logs api`

### Database

Currently uses in-memory storage. To add persistent database:

1. Add database service to `compose.yml`
2. Update backend configuration
3. Implement database repository adapters

## 🚨 Troubleshooting

### Port Already in Use

```bash
# Kill process on port 8080
lsof -ti:8080 | xargs kill -9

# Kill process on port 5173
lsof -ti:5173 | xargs kill -9
```

### Docker Issues

```bash
# Clean up containers and images
docker-compose down -v
docker-compose build --no-cache
docker-compose up
```

### Frontend Not Connecting to Backend

- Check API URL in environment variables
- Verify backend is running: `curl http://localhost:8080/readyz`
- Check CORS settings in backend

### Go Module Issues

```bash
# Update dependencies
go get -u ./...

# Clear module cache
go clean -modcache
```

## 📖 Learning Resources

### Understanding the Stack

- See `/docs/presentation/presentation.tex` for detailed explanation of:
  - Docker and containerization concepts
  - Server-Side Rendering (SSR)
  - SvelteKit framework
  - Go language and compilation
  - Modern development practices

### API Documentation

See `stock-ingestor/README.md` for detailed backend documentation.

## 🔐 Security Notes

- Never commit `.env` files with sensitive data
- Use environment variables for API keys
- Validate all user inputs in the backend
- Consider implementing authentication/authorization
