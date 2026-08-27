# Qubit — Architecture

## Architecture Style

Qubit follows a **Modular Monolith** architecture.

The application remains one deployable backend while business capabilities are separated into modules with explicit boundaries.

We are deliberately **not using microservices** at this stage. A module may be extracted later only if actual operational or scaling requirements justify it.

---

# High-Level Architecture

```text
                    Browser
                       │
                       ▼
              React / TypeScript
                       │
                 REST / WebSocket
                       │
                       ▼
                 Go + Gin API
                       │
        ┌──────────────┼──────────────┐
        │              │              │
        ▼              ▼              ▼
 Authentication   Future Modules   Cross-cutting
        │                              │
        ▼                              ▼
   PostgreSQL                       Redis (planned)
        │
        └──────────────┐
                       ▼
                Object Storage
                 (MinIO planned)
```

---

# Backend Dependency Direction

The current backend uses this conceptual direction:

```text
HTTP / Router
      ↓
Handler
      ↓
Service
      ↓
Repository abstraction
      ↓
PostgreSQL implementation
      ↓
pgx / PostgreSQL
```

Dependencies must not flow upward. In particular:

- Repositories do not call services.
- Repositories do not know about HTTP.
- Services do not construct HTTP responses.
- Handlers do not contain database operations.

---

# Current Authentication Architecture

```text
internal/auth/
├── DTOs / domain types
├── validation
├── normalization
├── password hashing / verification
├── errors
├── service/
│   └── registration application logic
├── repository/
│   └── PostgreSQL persistence
└── handler/
    └── Gin HTTP adapter
```

Current registration path:

```text
POST /api/v1/auth/register
        ↓
Gin Handler
        ↓
Registration Service
        ├── Validate input
        ├── Normalize email
        ├── Generate user ID
        ├── Hash password
        └── Persist account
              ↓
       PostgreSQL Repository
              ↓
          Transaction
          ├── users
          └── password_credentials
```

---

# Authentication Data Model

Identity and credentials are intentionally separated:

```text
users
  │
  ├── password_credentials      implemented
  ├── oauth_accounts             planned
  └── sessions                   planned
```

This avoids putting provider-specific authentication data into the core user table.

---

# Current Database Architecture

PostgreSQL is the system of record for relational application data.

Current authentication tables:

```text
users
  │
  └── password_credentials
```

The password credential has a foreign key to `users(id)` with `ON DELETE CASCADE`.

---

# Planned Backend Modules

- Authentication
- Organizations
- Workspaces
- Documents
- Tasks
- Chat
- Calendar
- Storage
- Search
- Notifications
- AI

These are **planned business boundaries**, not a requirement to implement every module immediately.

---

# Cross-Cutting Concerns

Planned cross-cutting capabilities include:

- Authentication middleware
- Request IDs
- Structured logging
- Panic recovery
- CORS
- Rate limiting
- Configuration management
- Health/readiness checks
- Graceful shutdown

Only introduce them when the relevant milestone requires them.

---

# Architecture Principles

1. **Modular monolith first.**
2. **Organize around business capabilities.**
3. **Keep HTTP, business logic, and persistence responsibilities separate.**
4. **Use PostgreSQL constraints as final data-integrity enforcement.**
5. **Use transactions for multi-step operations that must be atomic.**
6. **Prefer explicit, testable dependencies over global state.**
7. **Avoid abstraction for abstraction's sake.**
8. **Design for future scale without implementing distributed-system complexity prematurely.**
