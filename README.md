# Go REST CRUD API

A minimal REST API in Go with a clean, modular layout:

- `models/` – domain structs and DTOs
- `handlers/` – HTTP handlers (business logic + validation)
- `routes/` – router wiring
- `db/` – database init and helpers
- `main.go` – app entrypoint
- `items.db` – sample SQLite database for local testing
- `Dockerfile` – container build

> This project is a compact template you can extend with auth, pagination, filtering, etc.

---

## Features

- CRUD endpoints for an `Item` resource
- SQLite for zero-setup local runs (`items.db` included)
- Idiomatic Go project structure (handlers/routes/models split)
- Ready-to-containerize with the included `Dockerfile`

---

## Quick Start

### Prerequisites
- Go (1.21+ recommended)
- `make` (optional)
- SQLite (optional; not required if using the bundled `items.db`)

### Run locally

```bash
# 1) Get deps
go mod download

# 2) Run
go run ./...
# or
go run main.go




---

