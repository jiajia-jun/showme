# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
go mod download          # Install dependencies
go run main.go           # Start server on localhost:8080
go build -o main.exe     # Compile to binary
```

There are no tests yet.

## Architecture

Personal showcase website built with Gin (Go). Module name: `webproject`.

**Public site (no login required):**
- `/` — Single-page scrolling showcase (Hero → About → Contact)
- `/admin` — Admin login + profile editor (JWT-protected)

**Layers:**
- `router/` — Route initialization (`init.go`): static files, public API, protected API group
- `api/` — HTTP handlers (organized by feature)
  - `auth_handler.go` — `LoginUser`, `UpdateUserPassword`
  - `profile_handler.go` — `GetPublicProfile` (public GET), `UpdateProfile` (protected PUT), `CheckAdminAuth` (protected GET)
  - `message_handler.go` — `GetMessages`, `CreateMessage`, `LikeMessage`, `DeleteMessage`
- `middleware/` — JWT auth (`auth.go`), request logging (`logger.go`), static cache (`staticCache.go`)
- `dao/` — Data access layers:
  - `userdata.go` — User credentials: `map[string]string` persisted to `data/userdata.json`
  - `profiledata.go` — Profile info: `model.Profile` struct persisted to `./data/profile.json`
  - `messagedata.go` — Messages persisted to `./data/messages.json`
  - `imagedata.go` — Image file reader
- `model/` — Data models split by domain:
  - `authModel.go` — `User`, `UpdatePassword`
  - `profileModel.go` — `Profile`, `Skill`, `TimelineItem`, `Stat`
  - `messageModel.go` — `Message`
  - `image.go` — `ImageItem`, `ImagePath` constant
- `utils/` — JWT token generation/parsing (HS256, 2-hour expiry)

**Router structure:**
- Public HTML: `GET /` → `index.html`, `GET /admin` → `admin.html`
- Public API: `GET /api/profile`, `POST /api/login`, `POST /api/updatepassword`, `GET /api/messages`, `POST /api/messages`, `POST /api/messages/:id/like`
- Protected API (`/api` group with `AuthMiddleware()`): `PUT /api/profile`, `GET /api/admin/check`, `DELETE /api/messages/:id`

**JWT flow:** Login → server returns token → client stores in localStorage → `Authorization: Bearer <token>` → `AuthMiddleware` validates and injects `username` into Gin context.

## Important: Hardcoded Paths

- User data JSON: `data/userdata.json` (relative to working directory, auto-created on first run if missing)
- Profile data JSON: `./data/profile.json` (relative to working directory, auto-created with defaults)
- Message data JSON: `./data/messages.json` (relative to working directory, auto-created with defaults)

The admin account must be created manually by adding an entry to `data/userdata.json`, e.g., `{"admin": "yourpassword"}`.

## Frontend

- `static/index.html` + `home.js` + `showcase.css` — Public showcase (fetches `/api/profile`, typewriter effect, IntersectionObserver scroll animations)
- `static/admin.html` + `admin.js` — Admin panel: login view + profile editor view, toggled by JWT presence
- `static/docsy-styles.css` — Design system CSS variables shared across pages
- `static/styles.css` — Reset, animations, form/card styles used by admin page
