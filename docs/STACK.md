# ⚙️ Qubit Tech Stack

This document describes the technologies used throughout the project and the reasoning behind each choice.

---

# Frontend

| Technology | Purpose |
|------------|---------|
| React | UI Library |
| TypeScript | Type Safety |
| Vite | Build Tool |
| Tailwind CSS | Styling |
| shadcn/ui | UI Components |
| TanStack Query | Server State |
| Zustand | Client State |
| React Router | Routing |

---

# Backend

| Technology | Purpose |
|------------|---------|
| Go | Backend Language |
| Gin | HTTP Framework |
| PostgreSQL | Primary Database |
| Redis | Cache & Sessions |
| sqlc *(planned)* | Type-safe SQL |
| Goose | Database Migrations |

---

# Storage

| Technology | Purpose |
|------------|---------|
| MinIO | Object Storage |
| S3 API | Storage Interface |

---

# Infrastructure

| Technology | Purpose |
|------------|---------|
| Docker | Containerization |
| Docker Compose | Local Development |
| GitHub Actions | CI/CD |
| Nginx *(later)* | Reverse Proxy |

---

# Testing

| Technology | Purpose |
|------------|---------|
| Go Testing | Unit Testing |
| Testcontainers | Integration Tests |
| Vitest | Frontend Testing |
| Playwright | End-to-End Testing |

---

# API

- REST
- JSON
- OpenAPI / Swagger

---

# Authentication

- JWT Access Tokens
- Refresh Tokens
- OAuth (Google)
- Password Hashing (bcrypt/Argon2)

---

# Realtime

- WebSockets

---

# Logging & Monitoring

- slog (Go)
- OpenTelemetry *(later)*

---

# Development Tools

- Git
- GitHub
- VS Code
- Postman / Bruno
- Docker Desktop

---

# Guiding Principles

Technologies are chosen based on:

- Simplicity
- Scalability
- Maintainability
- Production Readiness
- Learning Value

Technology is never selected simply because it is popular.