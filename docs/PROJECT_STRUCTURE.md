# Qubit — Project Structure

> The structure is organized around business capabilities rather than generic technical layers.

---

# Repository Structure — Target

```text
qubit/
├── backend/
├── frontend/
├── docs/
├── .github/
├── README.md
├── docker-compose.yml
└── .env.example
```

---

# Backend — Current / Evolving Structure

```text
backend/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── dto.go
│   │   ├── validation.go
│   │   ├── normalise.go
│   │   ├── password.go
│   │   ├── errors.go
│   │   ├── password_test.go
│   │   ├── service/
│   │   │   ├── register.go
│   │   │   └── register_test.go
│   │   ├── repository/
│   │   │   ├── repository.go
│   │   │   ├── postgres.go
│   │   │   └── postgres_test.go
│   │   └── handler/
│   │       └── register.go
│   ├── database/
│   └── server/
├── migrations/
└── go.mod
```

The exact file list may evolve as the authentication module gains login, sessions, OAuth, and other capabilities.

---

# Planned Backend Modules

```text
auth
organization
workspace
document
task
chat
calendar
storage
notification
search
ai
```

Only create a module when its feature is actually being implemented.

---

# Module Responsibilities

## Handler

HTTP adapter.

- Parse HTTP input
- Call service
- Map errors to HTTP
- Serialize response

## Service

Application/business logic.

- Rules
- Validation coordination
- Orchestration
- Security-sensitive workflows

## Repository

Persistence boundary.

- SQL
- Database queries
- Transactions
- Persistence-specific error translation

## DTOs

API request/response shapes.

## Domain / model types

Business entities and values where needed.

## Errors

Module-specific domain/application errors.

---

# Authentication Module Direction

The authentication module is allowed to evolve beyond the initial registration structure.

Expected conceptual structure:

```text
auth/
├── handler/
├── service/
├── repository/
├── DTO/domain types
├── validation
├── password
└── errors
```

Do not force every module into an identical file structure if the real domain does not require it.

---

# Frontend — Planned

```text
frontend/
├── public/
├── src/
│   ├── app/
│   ├── pages/
│   ├── features/
│   ├── components/
│   ├── layouts/
│   ├── hooks/
│   ├── services/
│   ├── lib/
│   ├── store/
│   ├── utils/
│   ├── types/
│   └── assets/
└── package.json
```

---

# Frontend Rules

- UI components should not contain raw API calls.
- Server state should use TanStack Query where appropriate.
- Local state should remain local unless it has cross-component value.
- Feature-specific code should remain close to its feature.

---

# Module Communication

Modules should not reach directly into another module's repository.

Preferred:

```text
Workspace Service
       ↓
Organization application interface
```

Avoid:

```text
Workspace Repository
       ↓
Organization Repository
```

This keeps business capabilities independently understandable.

---

# Dependency Direction

```text
Handler
  ↓
Service
  ↓
Repository
  ↓
Database
```

Do not reverse this direction.

---

# Guiding Principle

> **Organize by business capability, not by technology.**
