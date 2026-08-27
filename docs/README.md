# Qubit

> A production-oriented collaborative workspace platform built to demonstrate strong backend, architecture, security, testing, and system-design skills.

Qubit aims to unify communication, documentation, project management, file sharing, and collaboration in one workspace.

The project is intentionally being developed as a **high-quality resume/portfolio project first**, rather than as a live SaaS product. The engineering standards remain production-oriented so the system can be extended later without rebuilding its foundations.

---

## Current Status

🚧 **Under Active Development**

### Current milestone

> **Authentication — Registration implemented and verified; registration architecture review is the current task.**

The registration flow is working end-to-end locally and the full Go test suite is currently green.

---

## Current Backend

- Go
- Gin
- PostgreSQL
- pgx / pgxpool
- golang-migrate
- Argon2id password hashing

Current backend layering:

```text
HTTP / Gin
    ↓
Handler
    ↓
Service
    ↓
Repository
    ↓
PostgreSQL
```

Qubit uses a **Modular Monolith** rather than microservices.

---

## Implemented So Far

### Foundation

- Go backend initialized
- HTTP server running
- Gin integrated
- Health endpoint implemented
- Environment configuration started
- PostgreSQL connection established
- pgxpool connection management established
- Database migrations configured
- `users` table created
- `password_credentials` table created

### Authentication / Registration

- Registration DTOs
- Input validation
- Email normalization
- Argon2id password hashing
- Argon2id password verification
- Authentication domain errors
- PostgreSQL repository
- Transactional account creation
- Rollback behavior
- Duplicate-email error translation
- Registration service
- Gin registration handler
- Service tests
- PostgreSQL repository integration tests
- HTTP handler tests
- Manual end-to-end registration verification

---

## API

### Health

```text
GET /api/v1/health
```

### Registration

```text
POST /api/v1/auth/register
```

Successful registration returns `201 Created` with safe user information only.

---

## Product Direction

Qubit is planned around these business capabilities:

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

Only the features that meaningfully strengthen the MVP should be implemented. **Quality > feature count.**

---

## Documentation

| Document | Purpose |
|---|---|
| `PRD.md` | Product requirements and scope |
| `ROADMAP.md` | Current and future implementation plan |
| `ARCHITECTURE.md` | System architecture and boundaries |
| `PROJECT_STRUCTURE.md` | Repository and module structure |
| `STACK.md` | Technology choices and status |
| `DECISIONS.md` | Architectural decisions and rationale |
| `FLOW.md` | Request and execution flows |
| `FEATURES.md` | Feature implementation/verification log |
| `BUGS.md` | Meaningful bug history |
| `CHANGELOG.md` | Notable project changes |
| `ENGINEERING.md` | Engineering standards |
| `CONTRIBUTING.md` | Git and contribution workflow |
| `HANDOVER.md` | Starting point for future AI/development sessions |

---

## Development Principle

Every meaningful feature should be:

- Correct
- Secure
- Testable
- Documented
- Maintainable
- Explainable in an interview
- Designed with reasonable future scalability

Avoid premature microservices, infrastructure complexity, and abstractions that do not solve a real problem.

---

## License

MIT
