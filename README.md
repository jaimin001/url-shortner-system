# URL Shortener

Backend-heavy URL shortener. The frontend is just a form and a table.

## Architecture

```
Browser → Go Backend → MongoDB + Valkey
```

## Backend (Go)

**Stack:** Gin · MongoDB · Valkey (Redis fork)

**Why this setup:**
- MongoDB is the source of truth — stores every shortened link
- Valkey handles two things:
  1. Caches redirects (fast path)
  2. Pre-generated key pool (no collision handling needed)

**How a shorten works:**
1. Pop a 6-char key from Valkey list
2. Save `{id, original_url, created_at, expires_at}` to MongoDB
3. Cache `{id → original_url}` in Valkey with TTL
4. Return the short key

**How a redirect works:**
1. Check Valkey — return if found (fast)
2. Check MongoDB — return if found, backfill Valkey cache
3. Return 404 if neither exists

**Key pool:**
- Valkey holds a list of ~1000 pre-generated 6-char keys
- Background goroutine replenishes when below threshold
- No collision logic needed — keys are guaranteed unique

**Link expiration:**
- `LINK_TTL` env var sets lifetime (default 24h)
- Valkey TTL handles auto-expiry for cache
- Background goroutine cleans MongoDB hourly

**Security:**
- Rate limiting per IP via Valkey (30 req/min default)
- API key required for `/shorten` endpoint
- URL validation: only http/https, max 2048 chars
- Security headers on all responses
- CORS whitelist only

## Frontend

React + Vite + Tailwind. Form to shorten, table to view recent links. That's it.

## Config

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | 8080 | Server port |
| `MONGODB_URI` | mongodb://localhost:27017 | MongoDB |
| `VALKEY_URL` | localhost:6379 | Valkey |
| `ALLOWED_ORIGIN` | http://localhost:5173 | CORS |
| `RATE_LIMIT` | 30 | req/min per IP |
| `API_KEY` | — | for `/shorten` |
| `LINK_TTL` | 24 | hours |

## Run

```bash
# Backend
cd src/backend
go run main.go

# Frontend
cd frontend
npm install && npm run dev
```

Set env vars via `.env` or system environment.
