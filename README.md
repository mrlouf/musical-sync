# musical-sync
A backend oriented project using JavaScript and Golang to synchronize your playlists across platforms.

## Architecture

This application consists of three main components running in Docker containers:

1. **Nginx** - Reverse proxy routing requests to frontend and backend
2. **Frontend** - Simple JavaScript-based web interface for playlist synchronization
3. **Backend** - Golang server that polls Deezer and Spotify APIs

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/mrlouf/musical-sync.git
cd musical-sync
```

### 2. Configure environment variables (optional)

```bash
cp .env.example .env
# Edit .env with your API keys if you have them
```

### 3. Build and run with Docker Compose

**Using Make (recommended):**

```bash
make build    # Build all images
make up       # Start all services
make logs     # View logs
make down     # Stop all services
make help     # See all available commands
```

**Using Docker Compose directly:**

```bash
docker compose up --build
```

The application will be available at `http://localhost`

### 4. Stop the application

```bash
make down
# or
docker compose down
```

## Services

### Nginx (Port 80)
- Acts as a reverse proxy
- Routes `/` to the frontend
- Routes `/api/*` to the backend

### Frontend (Internal: Port 8080)
- JavaScript-based single-page application
- Provides UI for checking sync status
- Communicates with backend via fetch API

### Backend (Internal: Port 8081)
- Golang REST API server
- Polls external APIs (Deezer, Spotify) every 30 seconds
- Endpoints:
  - `GET /health` - Health check
  - `GET /sync-status` - Current synchronization status
  - `GET /poll/deezer` - Manually trigger Deezer poll
  - `GET /poll/spotify` - Manually trigger Spotify poll

## Development

### Backend Development

```bash
cd backend
go run main.go
```

### Frontend Development

The frontend is a static HTML file using plain JavaScript. Edit `frontend/index.html` and refresh your browser.

## Project Structure

```
musical-sync/
├── nginx/
│   ├── Dockerfile
│   └── nginx.conf
├── frontend/
│   ├── Dockerfile
│   └── index.html
├── backend/
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
├── docker-compose.yml
├── .env.example
└── README.md
```

## Future Enhancements

- Implement actual Deezer API integration
- Implement actual Spotify API integration
- Add OAuth authentication flow
- Add database for storing sync history
- Add user authentication
- Implement actual playlist synchronization logic
