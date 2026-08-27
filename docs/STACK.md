# Qubit — Technology Stack

This document distinguishes what is **currently implemented** from what is **planned**.

---

# Frontend — Planned

| Technology | Purpose | Status |
|---|---|---|
| React | UI | Planned |
| TypeScript | Type safety | Planned |
| Vite | Build tool | Planned |
| Tailwind CSS | Styling | Planned |
| shadcn/ui | UI components | Planned |
| TanStack Query | Server state | Planned |
| Zustand | Client state | Planned |
| React Router | Routing | Planned |

---

# Backend — Current

| Technology | Purpose | Status |
|---|---|---|
| Go | Backend language | Implemented |
| Gin | HTTP framework | Implemented |
| PostgreSQL | Primary database | Implemented |
| pgx / pgxpool | PostgreSQL driver/pool | Implemented |
| golang-migrate | Database migrations | Implemented |
| Argon2id | Password hashing | Implemented |

---

# Backend — Planned

| Technology | Purpose | Status |
|---|---|---|
| sqlc | Type-safe SQL generation | Planned / evaluate before adoption |
| Redis | Caching/session-related infrastructure | Planned |
| WebSockets | Realtime features | Planned |

---

# Storage — Planned

| Technology | Purpose | Status |
|---|---|---|
| MinIO | Local object storage | Planned |
| S3 API | Storage abstraction | Planned |

---

# Infrastructure — Planned

| Technology | Purpose | Status |
|---|---|---|
| Docker | Containerization | Planned / local hardening later |
| Docker Compose | Local multi-service environment | Planned |
| GitHub Actions | CI | Planned |
| Structured logging / `slog` | Application logging | Planned |
| OpenAPI / Swagger | API documentation | Planned |
| Nginx | Reverse proxy | Later, only if deployment requires it |

---

# Testing

| Technology | Purpose | Status |
|---|---|---|
| Go testing | Backend unit/integration tests | Implemented |
| Testcontainers | Isolated integration environments | Planned / evaluate |
| Vitest | Frontend tests | Planned |
| Playwright | Browser E2E tests | Planned |

---

# Authentication

Current:

- Email/password
- Argon2id
- Password verification
- Separate password credential table

Planned:

- Short-lived access tokens
- Refresh tokens
- Session management
- Session revocation
- Google OAuth
- GitHub OAuth
- Email verification
- Password reset

---

# API

Current:

- REST
- JSON
- Versioned `/api/v1` paths

Planned:

- OpenAPI / Swagger

---

# Technology Selection Principles

Choose technologies based on:

- Simplicity
- Correctness
- Security
- Maintainability
- Learning value
- Demonstrable engineering depth
- Actual project requirements

Do not add infrastructure merely to create the appearance of scale.
