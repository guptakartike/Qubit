# Qubit — Handover

> Read this file first when starting a new AI-assisted development session.

---

# Project

Qubit is a collaborative workspace platform being built as a **strong resume/portfolio project with production-oriented engineering standards**.

It is currently a local development project, not a live SaaS product.

Primary goal:

> Demonstrate strong software engineering through a smaller number of well-designed, secure, tested, and explainable features.

---

# Current Architecture

```text
Modular Monolith

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

Do not introduce microservices unless an actual requirement justifies them.

---

# Current Stack

### Backend

- Go
- Gin
- PostgreSQL
- pgx / pgxpool
- golang-migrate
- Argon2id

### Planned / Later

- Redis
- WebSockets
- MinIO / S3-compatible storage
- Docker Compose hardening
- GitHub Actions
- OpenAPI / Swagger
- Structured logging

---

# Completed Foundation

- [x] Go backend
- [x] Gin server
- [x] Health endpoint
- [x] Environment configuration started
- [x] PostgreSQL connection
- [x] pgxpool
- [x] Database migrations
- [x] `users` table
- [x] `password_credentials` table

---

# Completed Authentication Work

## Registration

Status:

> **Implemented and verified. Final architecture review pending.**

Implemented:

- Register request/response DTOs
- Validation
- Email normalization
- Argon2id hashing
- Argon2id verification
- Authentication errors
- Registration service
- PostgreSQL repository
- Transactional user + credential creation
- Rollback handling
- Duplicate-email translation
- Gin handler
- Service tests
- Repository integration tests
- Handler tests
- Manual HTTP/database verification

---

# Current Test State

The full suite is green:

```bash
go test ./...
```

Relevant test areas:

```text
internal/auth
internal/auth/service
internal/auth/repository
internal/auth/handler
```

---

# Current Database Model

```text
users
├── id
├── name
├── email
├── status
├── email_verified_at
├── deleted_at
├── created_at
└── updated_at

password_credentials
├── user_id
├── password_hash
├── created_at
└── updated_at
```

Relationship:

```text
password_credentials.user_id
            ↓
         users.id
```

Foreign key uses `ON DELETE CASCADE`.

---

# Important Architecture Correction

Earlier documentation claimed that UUIDv7 was implemented.

That claim is currently **not verified**.

The current code uses the UUID library's `uuid.New()`.

Therefore:

- Do not document UUIDv7 as implemented.
- Review the identifier strategy before finalizing the architecture milestone.
- Do not change it silently.

See `DECISIONS.md` ADR-003.

---

# Current Task

## Registration Architecture Review

Before Login:

1. Inspect registration code and tests.
2. Review handler/service/repository boundaries.
3. Review domain-error propagation.
4. Review PostgreSQL `23505` translation.
5. Review validation and normalization.
6. Review Argon2id implementation and parameters.
7. Review transaction/rollback semantics.
8. Review database schema.
9. Review API contract.
10. Review test coverage.
11. Identify missing high-value tests.
12. Identify security issues.
13. Identify production-readiness issues.
14. Fix only justified issues.
15. Update documentation.
16. Commit the completed registration milestone.

Do **not** implement Login during this review.

---

# Next Feature

After the registration review:

> **Login**

Login should first be designed around:

- Credential lookup
- Account status
- Password verification
- Authentication failure behavior
- Session model
- Access token strategy
- Refresh token strategy

Do not jump directly into OAuth or unrelated modules.

---

# Authentication Roadmap

```text
Registration                 ✅
      ↓
Architecture Review          ← CURRENT
      ↓
Login                        ⏳
      ↓
Sessions / Tokens            ⏳
      ↓
Logout                       ⏳
      ↓
Email Verification           ⏳
      ↓
Password Reset               ⏳
      ↓
Google OAuth                 ⏳
      ↓
GitHub OAuth                 ⏳
```

---

# AI Development Rules

- Read this file and the other project docs before modifying code.
- Inspect actual source code rather than trusting documentation blindly.
- Do not dump code without explaining the design.
- Prefer simple, correct solutions.
- Do not introduce unnecessary abstractions.
- Keep HTTP logic out of repositories.
- Keep SQL/database logic out of handlers.
- Keep business logic out of handlers.
- Never expose passwords or password hashes.
- Never hardcode secrets.
- Update `DECISIONS.md` when a significant architectural decision changes.
- Update `FLOW.md` when execution paths change.
- Update `FEATURES.md`, `BUGS.md`, `CHANGELOG.md`, and `ROADMAP.md` when relevant.

---

# Git Milestone

Do not commit registration as complete until the architecture review is finished.

Expected milestone commit:

```text
feat(auth): implement user registration
```
