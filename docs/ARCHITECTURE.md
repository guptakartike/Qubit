# Architecture

## Architecture Style

Qubit follows a **Modular Monolith** architecture.

The system is designed so that each module is independently maintainable while remaining inside a single deployable application.

Future migration to microservices should require minimal code changes.

---

# Backend Modules

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

---

# High Level System

```
Browser
      │
      ▼
React Frontend
      │
REST API
WebSockets
      │
Go Backend
      │
 ┌────┴───────────┐
 │                │
PostgreSQL     Redis
 │
MinIO
```

---

# Guiding Principles

- Modular Design
- Separation of Concerns
- Clean Code
- SOLID Principles
- Secure by Default
- Testability
- Scalability