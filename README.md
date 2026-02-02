# GOKOZYY

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Built with Bubble Tea](https://img.shields.io/badge/Built%20with-Bubble%20Tea-blueviolet)](https://github.com/charmbracelet/bubbletea)

A modern, interactive CLI tool for scaffolding full-stack web applications with a **Vite + React + TypeScript** frontend and a **Go** backend. Generate production-ready projects in seconds with your choice of backend frameworks, databases, and UI libraries.

<div align="center">

```
 ██████╗  ██████╗ ██╗  ██╗ ██████╗ ███████╗██╗   ██╗██╗   ██╗
██╔════╝ ██╔═══██╗██║ ██╔╝██╔═══██╗╚══███╔╝╚██╗ ██╔╝╚██╗ ██╔╝
██║  ███╗██║   ██║█████╔╝ ██║   ██║  ███╔╝  ╚████╔╝  ╚████╔╝
██║   ██║██║   ██║██╔═██╗ ██║   ██║ ███╔╝    ╚██╔╝    ╚██╔╝
╚██████╔╝╚██████╔╝██║  ██╗╚██████╔╝███████╗   ██║      ██║
 ╚═════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝   ╚═╝      ╚═╝
```

</div>

## Features

### Backend Options

- **3 Framework Choices:**
  - **Standard Library** - Pure Go `net/http` package (zero dependencies)
  - **Chi** - Lightweight, idiomatic HTTP router
  - **Gin** - High-performance web framework with rich middleware

### Database Support

- **PostgreSQL** - Production-ready with `pgx/v5` driver and Docker support
- **SQLite** - Embedded database with `go-sqlite3` driver
- **None** - Start without a database and add it later

### Frontend Stack

- **Vite** - Lightning-fast build tool with HMR
- **React 18+** - Modern React with TypeScript
- **Tailwind CSS v4** - Latest version with Vite plugin
- **Bun** - Ultra-fast JavaScript runtime and package manager
- **shadcn/ui** (optional) - Beautiful, accessible component library built on Radix UI

### Developer Experience

- **Interactive TUI** - Beautiful terminal wizard powered by Bubble Tea
- **Hot Reload** - Air configuration for backend auto-restart
- **Docker Ready** - Multi-stage Dockerfile and docker-compose setup
- **TypeScript** - Full type safety with path aliases (`@/` imports)
- **Makefile** - Common development tasks preconfigured
- **Environment Config** - Sensible defaults in `.env` file

## Installation

### Option 1: Install from Source

Requires Go 1.23 or higher:

```bash
go install github.com/kozykoding/gokozyy@latest
```

Make sure your `$GOPATH/bin` is in your `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Option 2: Build from Source

Clone the repository and build:

```bash
git clone https://github.com/kozykoding/gokozyy.git
cd gokozyy
make build
sudo mv ./main /usr/local/bin/gokozyy
```

Or use it directly:

```bash
go run main.go create
```

### Verify Installation

```bash
gokozyy --help
```

## Quick Start

Create your first project in 3 steps:

```bash
# 1. Start the interactive wizard
gokozyy create

# 2. Follow the prompts to configure your project
# (See Interactive Wizard Guide below)

# 3. Start developing!
cd my-awesome-project
make docker-run         # (if using Docker + PostgreSQL)
make watch              # Start backend with hot-reload
cd frontend && bun dev  # Start React frontend
```

Your backend will run on `http://localhost:8080` and frontend on `http://localhost:5173`.

## Interactive Wizard Guide

The CLI features a beautiful, keyboard-driven wizard that walks you through project setup:

### Step 1: Project Name

```
╭────────────────────────────────────────────────────────────────╮
│ Enter your project name:                                       │
│ ▸ my-awesome-project                                           │
│                                                                │
│ [Enter to continue • q to quit]                                │
╰────────────────────────────────────────────────────────────────╯
```

Enter your desired project name (lowercase, hyphens recommended).

**Controls:** `Enter` to continue, `q` to quit

---

### Step 2: Backend Framework

```
╭────────────────────────────────────────────────────────────────╮
│ Select your Go backend framework:                              │
│                                                                │
│   ○ Standard-library                                           │
│   ● Chi                                                        │
│   ○ Gin                                                        │
│                                                                │
│ [↑/↓ to navigate • space to select • y to confirm]             │
╰────────────────────────────────────────────────────────────────╯
```

Choose your preferred backend framework:

| Framework            | Description                                            | Best For                                               |
| -------------------- | ------------------------------------------------------ | ------------------------------------------------------ |
| **Standard Library** | Pure Go `net/http` with `http.NewServeMux()`           | Simple APIs, learning Go, minimal dependencies         |
| **Chi**              | Lightweight, idiomatic router built on stdlib concepts | RESTful APIs, middleware-heavy apps, Go purists        |
| **Gin**              | High-performance framework with rich ecosystem         | Complex apps, need for speed, comprehensive middleware |

All frameworks come preconfigured with a `/api/health` endpoint returning `{"status":"ok"}`.

**Controls:** `↑/↓` to navigate, `space` to select, `y` to confirm

---

### Step 3: Database Driver

```
╭────────────────────────────────────────────────────────────────╮
│ Choose your database driver:                                   │
│                                                                │
│   ○ None                                                       │
│   ● Postgres                                                   │
│   ○ SQLite                                                     │
│                                                                │
│ [↑/↓ to navigate • space to select • y to confirm]             │
╰────────────────────────────────────────────────────────────────╯
```

Select your database option:

#### None

- No database driver or configuration
- Perfect for API gateways, static content servers, or projects that will use external services
- Add a database later as needed

#### PostgreSQL

- **Driver:** `github.com/jackc/pgx/v5/stdlib`
- **Docker Support:** Includes `docker-compose.yml` with PostgreSQL service
- **Connection:** Configured via environment variables in `.env`
- **Production Ready:** Health checks, volume persistence, optimized settings
- **Generated File:** `backend/internal/database/database.go` with `NewPostgres()` function

**Default PostgreSQL Configuration (.env):**

```bash
GOKOZYY_DB_HOST=localhost
GOKOZYY_DB_PORT=5432
GOKOZYY_DB_DATABASE=gokozyy
GOKOZYY_DB_USERNAME=sammy
GOKOZYY_DB_PW=thisismypassword
GOKOZYY_DB_SCHEMA=public
```

#### SQLite

- **Driver:** `github.com/mattn/go-sqlite3`
- **Embedded:** No external database server required
- **Lightweight:** Perfect for development, small apps, or edge deployments
- **Generated File:** `backend/internal/database/database.go` with `NewSQLite(path string)` function

**Controls:** `↑/↓` to navigate, `space` to select, `y` to confirm

---

### Step 4: Docker Support

```
╭────────────────────────────────────────────────────────────────╮
│ Do you want Docker/docker-compose files for your backend/DB?   │
│                                                                │
│ [y for Yes • n for No • h to go back]                          │
╰────────────────────────────────────────────────────────────────╯
```

Choose whether to include Docker configuration:

#### Yes (Recommended for PostgreSQL)

Generates:

- **Multi-stage Dockerfile** with 4 optimized stages:
  1. `backend-builder` - Compiles Go binary (golang:1.23-alpine)
  2. `frontend-builder` - Builds Vite app (oven/bun:latest)
  3. `prod` - Production API server (alpine:latest, ~20MB)
  4. `frontend` - Development frontend server (oven/bun:latest)

- **docker-compose.yml** with 3 services:
  - `app` - Backend API server with environment configuration
  - `frontend` - Frontend development server on port 5173
  - `psql_gokozyy` - PostgreSQL with health checks and volume persistence

#### No

- Develop locally without Docker
- Manually install PostgreSQL or use SQLite
- Smaller project footprint

**Controls:** `y` for Yes, `n` for No, `h` to go back

---

### Step 5: Frontend Stack

```
╭────────────────────────────────────────────────────────────────╮
│ Select your frontend stack:                                    │
│                                                                │
│   ● Vite + React + Tailwind + Bun                              │
│   ○ Vite + React + Tailwind + shadcn/ui + Bun                  │
│                                                                │
│ [↑/↓ to navigate • space to select • y to confirm]             │
╰────────────────────────────────────────────────────────────────╯
```

Choose your frontend configuration:

#### Option 1: Vite + React + Tailwind + Bun (Basic)

**What You Get:**

- React 18+ with TypeScript
- Vite for blazing-fast HMR and builds
- Tailwind CSS v4 with Vite plugin
- Bun as runtime and package manager
- TypeScript path aliases (`@/` for clean imports)
- Hot module replacement enabled

**Generated Structure:**

```
frontend/
├── src/
│   ├── App.tsx           # Main app component
│   ├── index.css         # Tailwind imports
│   └── main.tsx          # Entry point
├── vite.config.ts        # Vite + Tailwind plugin + @ alias
├── tailwind.config.ts    # Tailwind v4 configuration
├── tsconfig.json         # Base TypeScript config
├── tsconfig.app.json     # App TypeScript config with paths
└── package.json          # Bun dependencies
```

**Perfect for:**

- Quick prototypes
- Custom design systems
- Learning React + Tailwind
- Full control over component architecture

---

#### Option 2: Vite + React + Tailwind + shadcn/ui + Bun (Enhanced)

**Everything from Option 1, PLUS:**

**Additional Dependencies:**

- `lucide-react` - Beautiful icon library (800+ icons)
- `class-variance-authority` - CVA for component variants
- `clsx` - Conditional className utility
- `tailwind-merge` - Intelligent Tailwind class merging
- `tailwindcss-animate` - Animation utilities
- `@radix-ui/react-slot` - Radix UI primitives

**Additional Files:**

```
frontend/
├── components.json              # shadcn/ui configuration
├── src/
│   ├── components/ui/
│   │   └── button.tsx          # Pre-built Button component
│   └── lib/
│       └── utils.ts            # cn() utility for class merging
└── [... all files from Option 1]
```

**Example Button Component:**

```tsx
import { Button } from "@/components/ui/button"

<Button variant="default">Click me</Button>
<Button variant="destructive">Delete</Button>
<Button variant="outline" size="lg">Large Outline</Button>
```

**Ready to Add More Components:**

```bash
# Add any shadcn/ui component with one command
bunx shadcn@latest add card
bunx shadcn@latest add dialog
bunx shadcn@latest add form
```

**Perfect for:**

- Production applications
- Rapid UI development
- Accessible, polished interfaces
- Enterprise-grade design systems

**Controls:** `↑/↓` to navigate, `space` to select, `y` to confirm

---

### Step 6: Summary & Confirmation

```
╭────────────────────────────────────────────────────────────────╮
│ Summary of your project configuration:                         │
│                                                                │
│ Project Name: my-awesome-project                               │
│ Framework:    Chi                                              │
│ Database:     Postgres                                         │
│ Docker:       Yes                                              │
│ Frontend:     Vite + React + Tailwind + shadcn/ui + Bun        │
│                                                                │
│ [Enter/y to create • h to go back • q to quit]                 │
╰────────────────────────────────────────────────────────────────╯
```

Review your selections before generating the project.

**Controls:** `Enter` or `y` to create, `h` to go back, `q` to quit

---

### Success

```
✅ Project "my-awesome-project" created successfully!

🚀 Next steps to start nerding out:
  1. cd my-awesome-project && nvim .
  2. make docker-run         # Start your Postgres database
  3. make watch              # Start backend with hot-reload (Air)
  4. cd frontend && bun dev  # Start React frontend environment

Happy coding!
```

## Generated Project Structure

After running the wizard, you'll get a fully configured project:

```
my-awesome-project/
├── backend/                          # Go backend
│   ├── cmd/
│   │   └── server/
│   │       └── main.go              # Backend entry point
│   ├── internal/
│   │   └── database/
│   │       └── database.go          # DB connection (if selected)
│   ├── go.mod                       # Go dependencies
│   └── go.sum
├── frontend/                         # React frontend
│   ├── src/
│   │   ├── components/ui/           # shadcn components (if selected)
│   │   ├── lib/
│   │   │   └── utils.ts            # Utility functions
│   │   ├── App.tsx
│   │   ├── index.css               # Tailwind imports
│   │   └── main.tsx
│   ├── components.json              # shadcn config (if selected)
│   ├── vite.config.ts
│   ├── tailwind.config.ts
│   ├── tsconfig.json
│   └── package.json
├── Dockerfile                        # Multi-stage build (if Docker selected)
├── docker-compose.yml               # Services config (if Docker selected)
├── .air.toml                        # Hot reload config
├── Makefile                         # Development commands
├── .env                             # Environment variables
└── .gitignore                       # Git ignore patterns
```

## Development Workflow

### Backend Development

#### Start with Hot Reload (Recommended)

```bash
make watch
```

Uses [Air](https://github.com/air-verse/air) for automatic restarts on code changes. Air is auto-installed if not present.

**Watches:**

- `*.go` files
- `*.sql` files
- Template files (`.tpl`, `.tmpl`, `.html`)

**Excludes:**

- `frontend/node_modules`
- `frontend/dist`
- Test files

#### Traditional Run

```bash
make run
# or
cd backend && go run cmd/server/main.go
```

#### Build Binary

```bash
make build
./main
```

#### Run Tests

```bash
make test
```

#### Clean Build Artifacts

```bash
make clean
```

---

### Frontend Development

```bash
cd frontend

# Start dev server (hot reload enabled)
bun dev

# Build for production
bun run build

# Preview production build
bun run preview

# Install new dependencies
bun add <package-name>
```

Frontend runs on `http://localhost:5173` by default.

---

### Database Management

#### PostgreSQL with Docker

**Start Database:**

```bash
make docker-run
# or
docker-compose up -d psql_gokozyy
```

**Stop Database:**

```bash
make docker-down
# or
docker-compose down
```

**Connect to Database:**

```bash
docker exec -it psql_gokozyy psql -U sammy -d gokozyy
```

**View Logs:**

```bash
docker-compose logs -f psql_gokozyy
```

#### SQLite

SQLite database is created automatically when your app runs. Location is specified in your code:

```go
db, err := database.NewSQLite("./app.db")
```

---

### Full Stack Development

Open 3 terminal windows:

```bash
# Terminal 1: Database (if using PostgreSQL + Docker)
make docker-run

# Terminal 2: Backend
make watch

# Terminal 3: Frontend
cd frontend && bun dev
```

Visit `http://localhost:5173` - Vite will proxy API requests to `http://localhost:8080/api/*`.

## Configuration

### Environment Variables

All projects include a `.env` file with sensible defaults:

```bash
# Application Configuration
PORT=42069                    # Backend API port
APP_ENV=local                 # Environment (local, dev, prod)

# PostgreSQL Configuration (if using Postgres)
GOKOZYY_DB_HOST=localhost
GOKOZYY_DB_PORT=5432
GOKOZYY_DB_DATABASE=gokozyy
GOKOZYY_DB_USERNAME=sammy
GOKOZYY_DB_PW=thisismypassword
GOKOZYY_DB_SCHEMA=public
```

**Security Note:** The `.env` file is automatically added to `.gitignore` and has `0600` permissions (owner read/write only).

### TypeScript Path Aliases

Both frontend options include `@/` path alias for clean imports:

```typescript
// Instead of this:
import { Button } from "../../components/ui/button";

// Write this:
import { Button } from "@/components/ui/button";
```

Configured in `tsconfig.json`:

```json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"]
    }
  }
}
```

### Vite Configuration

Vite is configured with:

- Tailwind CSS v4 plugin
- Path resolution for `@/` alias
- React plugin with Fast Refresh

```typescript
// vite.config.ts
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
import path from "path";

export default defineConfig({
  plugins: [react(), tailwindcss()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});
```

### Tailwind CSS v4

Configured with the new Vite plugin approach:

```typescript
// tailwind.config.ts
import type { Config } from "tailwindcss";

export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
} satisfies Config;
```

Import in CSS:

```css
/* src/index.css */
@import "tailwindcss";
```

### Air Configuration (Hot Reload)

The `.air.toml` file configures automatic backend restarts:

```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "make build"
  bin = "./main"
  exclude_dir = ["frontend/node_modules", "frontend/dist", "tmp", "testdata"]
  include_ext = ["go", "tpl", "tmpl", "html", "sql"]
  delay = 1000  # ms
  stop_on_error = true
```

## Docker Setup

### Using Docker Compose (Recommended)

Projects with Docker enabled include a complete `docker-compose.yml`:

#### Start All Services

```bash
docker-compose up -d
```

This starts:

- **Backend API** - Production build on port specified in `.env`
- **Frontend Dev Server** - Hot reload on port 5173
- **PostgreSQL** - Database with health checks on port 5432

#### View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f app
docker-compose logs -f frontend
docker-compose logs -f psql_gokozyy
```

#### Stop Services

```bash
docker-compose down

# Stop and remove volumes (deletes database data)
docker-compose down -v
```

#### Rebuild After Changes

```bash
docker-compose up -d --build
```

---

### Docker Services Explained

#### 1. Backend API (`app`)

```yaml
app:
  build:
    context: .
    dockerfile: Dockerfile
    target: prod
  env_file: .env
  depends_on:
    psql_gokozyy:
      condition: service_healthy
```

- **Multi-stage build** for optimal image size (~20MB)
- **Wait for database** health check before starting
- **Environment variables** from `.env` file
- **Alpine-based** production image with CA certificates

#### 2. Frontend Dev Server (`frontend`)

```yaml
frontend:
  build:
    context: .
    dockerfile: Dockerfile
    target: frontend
  ports:
    - "5173:5173"
  volumes:
    - ./frontend:/app/frontend
```

- **Hot reload** enabled with volume mounting
- **Bun runtime** for fast package installation
- **Port 5173** exposed for development

#### 3. PostgreSQL Database (`psql_gokozyy`)

```yaml
psql_gokozyy:
  image: postgres:latest
  environment:
    POSTGRES_USER: ${GOKOZYY_DB_USERNAME}
    POSTGRES_PASSWORD: ${GOKOZYY_DB_PW}
    POSTGRES_DB: ${GOKOZYY_DB_DATABASE}
  healthcheck:
    test: ["CMD-SHELL", "pg_isready -U ${GOKOZYY_DB_USERNAME}"]
    interval: 5s
    timeout: 5s
    retries: 5
  volumes:
    - psql_data_gokozyy:/var/lib/postgresql/data
```

- **Persistent storage** with named volume
- **Health checks** ensure database is ready
- **Environment-based** configuration from `.env`

---

### Dockerfile Overview

The multi-stage Dockerfile optimizes for both size and performance:

#### Stage 1: Backend Builder

```dockerfile
FROM golang:1.23-alpine AS backend-builder
# Compiles Go binary with CGO disabled
# Result: Single static binary
```

#### Stage 2: Frontend Builder

```dockerfile
FROM oven/bun:latest AS frontend-builder
# Runs bun install and bun run build
# Result: Static files in frontend/dist
```

#### Stage 3: Production API

```dockerfile
FROM alpine:latest AS prod
# Copies backend binary + frontend dist
# Installs CA certificates
# Result: Tiny image serving API + static frontend
```

#### Stage 4: Frontend Dev Server

```dockerfile
FROM oven/bun:latest AS frontend
# Development environment with hot reload
# Used by docker-compose for frontend service
```

---

### Production Deployment

Build and run the production image:

```bash
# Build production target
docker build -t my-app:latest --target prod .

# Run production container
docker run -d \
  --name my-app \
  -p 42069:42069 \
  --env-file .env \
  my-app:latest
```

The production image serves:

- **Backend API** on port from `.env` (default: 42069)
- **Frontend static files** from `/app/frontend/dist`

## Example Configurations

### Example 1: Simple REST API (No Database)

**Configuration:**

- Framework: Standard Library
- Database: None
- Docker: No
- Frontend: Vite + React + Tailwind + Bun

**Use Case:**

- API gateway
- Microservice that consumes other APIs
- Static content server
- Serverless function development

**Generated Backend:**

```go
// backend/cmd/server/main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

---

### Example 2: Full-Stack App with PostgreSQL

**Configuration:**

- Framework: Chi
- Database: PostgreSQL
- Docker: Yes
- Frontend: Vite + React + Tailwind + shadcn/ui + Bun

**Use Case:**

- SaaS application
- Content management system
- E-commerce platform
- Data-driven dashboards

**Start Development:**

```bash
cd my-saas-app
make docker-run   # Starts PostgreSQL
make watch        # Starts backend with hot-reload
cd frontend && bun dev  # Starts frontend
```

**Database Usage:**

```go
// backend/cmd/server/main.go
import "github.com/you/my-saas-app/backend/internal/database"

func main() {
    db, err := database.NewPostgres()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Use db for queries
    // ...
}
```

---

### Example 3: Lightweight Desktop App

**Configuration:**

- Framework: Gin
- Database: SQLite
- Docker: No
- Frontend: Vite + React + Tailwind + Bun

**Use Case:**

- Desktop applications (with Tauri/Electron)
- Local-first apps
- Development tools
- Personal productivity apps

**SQLite Usage:**

```go
// backend/cmd/server/main.go
import "github.com/you/my-desktop-app/backend/internal/database"

func main() {
    db, err := database.NewSQLite("./data/app.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Database is embedded in your app
    // ...
}
```

---

### Example 4: Rapid Prototype with shadcn/ui

**Configuration:**

- Framework: Chi
- Database: SQLite
- Docker: No
- Frontend: Vite + React + Tailwind + shadcn/ui + Bun

**Use Case:**

- Hackathons
- MVPs and prototypes
- Client demos
- Internal tools

**Quick Component Development:**

```bash
cd frontend

# Add shadcn components as needed
bunx shadcn@latest add card
bunx shadcn@latest add form
bunx shadcn@latest add dialog
bunx shadcn@latest add table
```

```tsx
// frontend/src/App.tsx
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

function App() {
  return (
    <div className="container mx-auto p-8">
      <Card>
        <CardHeader>
          <CardTitle>My Awesome MVP</CardTitle>
        </CardHeader>
        <CardContent>
          <Button onClick={() => fetch("/api/health")}>Check API Health</Button>
        </CardContent>
      </Card>
    </div>
  );
}
```

## Makefile Commands Reference

Every project includes these preconfigured commands:

| Command            | Description                                            |
| ------------------ | ------------------------------------------------------ |
| `make build`       | Compiles Go backend binary to `./main`                 |
| `make run`         | Runs backend with `go run cmd/server/main.go`          |
| `make test`        | Runs all Go tests with `go test ./...`                 |
| `make clean`       | Removes compiled binaries and artifacts                |
| `make watch`       | Starts backend with Air hot-reload (auto-installs Air) |
| `make docker-run`  | Starts PostgreSQL in Docker (if applicable)            |
| `make docker-down` | Stops Docker containers                                |

## Technology Stack

### CLI Tool

- **Language:** Go 1.25.5
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra) - Modern CLI applications
- **TUI Framework:** [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal user interfaces
- **Styling:** [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions for TUIs
- **UI Components:** [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components

### Generated Backend

- **Go:** 1.23+
- **Frameworks:**
  - Standard Library (`net/http`)
  - [Chi](https://github.com/go-chi/chi) v5
  - [Gin](https://github.com/gin-gonic/gin)
- **Database Drivers:**
  - [pgx](https://github.com/jackc/pgx) v5 (PostgreSQL)
  - [go-sqlite3](https://github.com/mattn/go-sqlite3) (SQLite)
- **Hot Reload:** [Air](https://github.com/air-verse/air)
- **Containerization:** Docker, docker-compose

### Generated Frontend

- **Runtime:** [Bun](https://bun.sh) - Fast all-in-one JavaScript runtime
- **Build Tool:** [Vite](https://vitejs.dev) - Next generation frontend tooling
- **Framework:** [React](https://react.dev) 18+ with TypeScript
- **Styling:** [Tailwind CSS](https://tailwindcss.com) v4
- **UI Library (Optional):** [shadcn/ui](https://ui.shadcn.com) - Re-usable components
- **UI Primitives:** [Radix UI](https://www.radix-ui.com) - Accessible component primitives
- **Icons:** [Lucide React](https://lucide.dev) - Beautiful icons
- **Utilities:**
  - [CVA](https://cva.style/docs) - Class variance authority
  - [clsx](https://github.com/lukeed/clsx) - Conditional classes
  - [tailwind-merge](https://github.com/dcastil/tailwind-merge) - Merge Tailwind classes

## Troubleshooting

### Installation Issues

#### Go Version Too Old

```
Error: requires go >= 1.23
```

**Solution:** Update Go:

```bash
# Linux/macOS
curl -OL https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz

# Verify
go version
```

#### $GOPATH/bin Not in PATH

```
Error: gokozyy: command not found
```

**Solution:** Add to your shell profile:

```bash
# For bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc

# For zsh
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

---

### Project Generation Issues

#### Bun Not Found

```
Error: exec: "bun": executable file not found in $PATH
```

**Solution:** Install Bun:

```bash
curl -fsSL https://bun.sh/install | bash
```

#### Permission Denied

```
Error: permission denied creating project directory
```

**Solution:** Run from a directory where you have write permissions or use sudo:

```bash
cd ~  # Move to home directory
gokozyy create
```

---

### Development Issues

#### Port Already in Use

```
Error: bind: address already in use
```

**Solution:** Change port in `.env` or kill the process:

```bash
# Find process using port 8080
lsof -i :8080

# Kill process
kill -9 <PID>

# Or change port in .env
PORT=8081
```

#### PostgreSQL Connection Failed

```
Error: failed to connect to postgres://...
```

**Solution:** Ensure PostgreSQL is running:

```bash
# Check if Docker container is running
docker ps | grep psql_gokozyy

# Start if not running
make docker-run

# Check logs
docker logs psql_gokozyy
```

#### Air Not Reloading

```
Air starts but doesn't reload on changes
```

**Solution:** Check `.air.toml` exclude patterns and file extensions:

```toml
include_ext = ["go", "tpl", "tmpl", "html", "sql"]
exclude_dir = ["frontend/node_modules", "frontend/dist"]
```

#### Frontend Build Errors

```
Error: Cannot find module '@/components/ui/button'
```

**Solution:** Ensure path aliases are configured:

```bash
# Check tsconfig.json has baseUrl and paths
cat frontend/tsconfig.json

# Restart TypeScript server in your editor
# VSCode: Cmd/Ctrl + Shift + P -> "TypeScript: Restart TS Server"
```

---

### Docker Issues

#### Docker Compose Database Not Healthy

```
Warning: dependency on psql_gokozyy not healthy
```

**Solution:** Check database logs and increase health check timeout:

```bash
# View database logs
docker-compose logs psql_gokozyy

# Check if PostgreSQL is ready
docker exec psql_gokozyy pg_isready -U sammy
```

#### Volume Permission Issues

```
Error: permission denied writing to volume
```

**Solution:** Fix volume permissions or use named volume:

```bash
# Remove and recreate volume
docker-compose down -v
docker-compose up -d
```

#### Build Cache Issues

```
Docker build uses stale files
```

**Solution:** Force rebuild without cache:

```bash
docker-compose build --no-cache
docker-compose up -d
```

---

### shadcn/ui Issues

#### Components Not Found

```
Error: Cannot resolve @/components/ui/button
```

**Solution:** Verify `components.json` exists and paths are correct:

```bash
cat frontend/components.json

# Should show:
{
  "aliases": {
    "components": "@/components",
    "utils": "@/lib/utils"
  }
}
```

#### Adding Components Fails

```
Error: Cannot find components.json
```

**Solution:** Ensure you're in frontend directory:

```bash
cd frontend
bunx shadcn@latest add button
```

---

### General Tips

1. **Always check logs first:**

   ```bash
   # Backend logs with Air
   # Displayed in terminal where you ran `make watch`

   # Frontend logs
   # Displayed in terminal where you ran `bun dev`

   # Docker logs
   docker-compose logs -f
   ```

2. **Clean rebuild often helps:**

   ```bash
   make clean
   make build

   cd frontend
   rm -rf node_modules
   bun install
   ```

3. **Verify environment variables:**

   ```bash
   cat .env
   # Check all values are correct
   ```

4. **Check port conflicts:**

   ```bash
   lsof -i :8080  # Backend
   lsof -i :5173  # Frontend
   lsof -i :5432  # PostgreSQL
   ```

## Roadmap

### Upcoming Features

- [ ] **Non-Interactive Mode** - Support CLI flags for CI/CD and automation
- [ ] **Additional Backend Frameworks** - Echo, Fiber, Gorilla Mux
- [ ] **More Database Options** - MySQL, MongoDB, Redis
- [ ] **Authentication Templates** - JWT, OAuth2, session-based auth
- [ ] **API Documentation** - Auto-generate Swagger/OpenAPI specs
- [ ] **Testing Setup** - Pre-configured testing frameworks and examples
- [ ] **CI/CD Templates** - GitHub Actions, GitLab CI, Jenkins pipelines
- [ ] **Deployment Configs** - Kubernetes, Railway, Fly.io, Vercel
- [ ] **Additional Frontend Options** - Vue, Svelte, Solid
- [ ] **Monorepo Support** - Turborepo, Nx integration
- [ ] **GraphQL Support** - gqlgen backend + Apollo client
- [ ] **WebSocket Support** - Real-time communication templates
- [ ] **Migration Tools** - Database migration generation (golang-migrate)
- [ ] **Logging & Monitoring** - Structured logging, OpenTelemetry
- [ ] **Configuration Management** - Viper, custom configs
- [ ] **Plugin System** - Community templates and extensions

### Recently Added

- ✅ Interactive TUI with Bubble Tea
- ✅ Tailwind CSS v4 support
- ✅ shadcn/ui integration
- ✅ Docker multi-stage builds
- ✅ Air hot-reload configuration
- ✅ Bun runtime support
- ✅ TypeScript path aliases

## Contributing

We welcome contributions! Here's how you can help:

### Reporting Bugs

Open an issue with:

- Your OS and Go version
- Steps to reproduce
- Expected vs actual behavior
- Error messages and logs

### Suggesting Features

Open an issue with:

- Clear description of the feature
- Use cases and benefits
- Proposed implementation (optional)

### Contributing Code

1. **Fork the repository**

   ```bash
   git clone https://github.com/yourusername/gokozyy.git
   cd gokozyy
   ```

2. **Create a feature branch**

   ```bash
   git checkout -b feature/amazing-feature
   ```

3. **Make your changes**
   - Follow Go conventions and formatting
   - Add comments for complex logic
   - Update documentation if needed

4. **Test your changes**

   ```bash
   # Test the CLI
   go run main.go create

   # Test generated project
   cd test-project
   make watch
   cd frontend && bun dev
   ```

5. **Commit with clear messages**

   ```bash
   git add .
   git commit -m "feat: add support for Echo framework"
   ```

6. **Push and create PR**

   ```bash
   git push origin feature/amazing-feature
   ```

### Development Setup

```bash
# Clone repository
git clone https://github.com/yourusername/gokozyy.git
cd gokozyy

# Install dependencies
go mod download

# Build
make build

# Run locally
go run main.go create

# Test changes
./main create
```

### Code Style

- Run `gofmt` before committing
- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Keep functions focused and small
- Add comments for exported functions
- Use meaningful variable names

### Adding New Templates

To add a new backend framework or feature:

1. Create generator function in `internal/generator/`

   ```go
   // generate_myframework.go
   func generateMyFramework(config Config) error {
       // Implementation
   }
   ```

2. Add to wizard options in `internal/ui/wizard.go`

   ```go
   frameworkOptions := []RadioOption{
       // ... existing options
       {Label: "MyFramework", Value: "myframework"},
   }
   ```

3. Wire up in `internal/generator/generator.go`

   ```go
   switch config.Framework {
   // ... existing cases
   case "myframework":
       if err := generateMyFramework(config); err != nil {
           return err
       }
   }
   ```

4. Test thoroughly with all combinations

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```
MIT License

Copyright (c) 2025 gokozyy contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Acknowledgments

This project is built with and inspired by amazing open-source tools:

- [Cobra](https://github.com/spf13/cobra) - For powerful CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - For delightful TUIs
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - For beautiful terminal styling
- [Vite](https://vitejs.dev) - For blazing-fast frontend tooling
- [Bun](https://bun.sh) - For incredible JavaScript performance
- [shadcn/ui](https://ui.shadcn.com) - For beautiful, accessible components
- [Tailwind CSS](https://tailwindcss.com) - For utility-first styling
- The entire Go and React communities

## Support

- **Documentation:** You're reading it!
- **Issues:** [GitHub Issues](https://github.com/yourusername/gokozyy/issues)
- **Discussions:** [GitHub Discussions](https://github.com/yourusername/gokozyy/discussions)

---

<div align="center">

**Built with ❤️ using Go and Bubble Tea**

[Report Bug](https://github.com/yourusername/gokozyy/issues) · [Request Feature](https://github.com/yourusername/gokozyy/issues) · [Contribute](https://github.com/yourusername/gokozyy/pulls)

**Star ⭐ this repo if you find it helpful!**

</div>
