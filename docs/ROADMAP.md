# Qubit — Development Roadmap

> Build a small number of technically strong features rather than a large number of shallow features.

---

# Development Workflow

Every major feature follows:

```text
Requirements
    ↓
Architecture
    ↓
Database Design
    ↓
API Design
    ↓
Implementation
    ↓
Testing
    ↓
Review
    ↓
Documentation
    ↓
Git Commit
```

---

# Phase 0 — Foundation

### Status: 🟢 Core backend foundation complete; broader infrastructure remains planned

- [x] Backend repository initialized
- [x] Go backend
- [x] Gin HTTP server
- [x] Health endpoint
- [x] Environment configuration started
- [x] PostgreSQL
- [x] pgxpool
- [x] Database migrations
- [x] Initial authentication schema
- [ ] Docker/Compose hardening
- [ ] Structured logging
- [ ] Graceful shutdown
- [ ] CI pipeline

---

# Phase 1 — Authentication

## Milestone 1 — Registration

### Status: 🟢 Implemented and verified

- [x] Registration DTOs
- [x] Validation
- [x] Email normalization
- [x] Argon2id hashing
- [x] Domain errors
- [x] Repository
- [x] PostgreSQL transaction
- [x] Rollback behavior
- [x] Duplicate email handling
- [x] Registration service
- [x] Gin handler
- [x] Service tests
- [x] Repository integration tests
- [x] Handler tests
- [x] Manual end-to-end verification
- [ ] Final architecture/security review
- [ ] Documentation milestone update after review
- [ ] Git milestone commit

## Milestone 2 — Login

### Status: ⏳ Next after registration review

Before implementation, design:

- [ ] Credential lookup
- [ ] User/account status handling
- [ ] Password verification
- [ ] Authentication failure semantics
- [ ] Session model
- [ ] Access-token strategy
- [ ] Refresh-token strategy
- [ ] Token/session storage

Then implement and test login.

## Milestone 3 — Sessions & Tokens

- [ ] Session model
- [ ] Short-lived access tokens
- [ ] Refresh tokens
- [ ] Refresh-token rotation/revocation strategy
- [ ] Logout
- [ ] Multi-device sessions
- [ ] Session revocation

## Milestone 4 — Email Verification

- [ ] Verification token model
- [ ] Verification endpoint
- [ ] Token expiration
- [ ] Email delivery abstraction
- [ ] Verification state handling

## Milestone 5 — Password Recovery

- [ ] Forgot-password request
- [ ] Reset token/session
- [ ] Password reset
- [ ] Token expiration/revocation

## Milestone 6 — OAuth

- [ ] Google OAuth
- [ ] GitHub OAuth
- [ ] OAuth identity model
- [ ] Account linking rules

---

# Phase 2 — Organizations

- [ ] Organization creation
- [ ] Organization membership
- [ ] Invitations
- [ ] Roles
- [ ] Permissions
- [ ] Permission inheritance rules

---

# Phase 3 — Workspaces

- [ ] Workspace CRUD
- [ ] Workspace membership
- [ ] Workspace dashboard
- [ ] Access control

---

# Phase 4 — Documents

- [ ] Documents
- [ ] Markdown / rich text
- [ ] Comments
- [ ] Version history
- [ ] Basic collaboration

Advanced real-time editing is intentionally later.

---

# Phase 5 — Task Management

- [ ] Boards
- [ ] Tasks
- [ ] Labels
- [ ] Priority
- [ ] Due dates
- [ ] Activity log

---

# Phase 6 — Chat

- [ ] Channels
- [ ] Direct messages
- [ ] Messages
- [ ] Threads
- [ ] Mentions
- [ ] Reactions

---

# Phase 7 — Realtime

- [ ] WebSockets
- [ ] Presence
- [ ] Typing indicators
- [ ] Live updates

---

# Phase 8 — File Storage

- [ ] MinIO
- [ ] Uploads
- [ ] Folders
- [ ] Sharing
- [ ] Storage abstraction

---

# Phase 9 — Search

- [ ] PostgreSQL-backed search for MVP
- [ ] Global search
- [ ] Search across users/documents/tasks/files
- [ ] Dedicated search engine only if justified later

---

# Phase 10 — Notifications / Calendar

- [ ] Notifications
- [ ] Calendar
- [ ] Activity-driven notifications

---

# Phase 11 — AI

- [ ] Document summaries
- [ ] Workspace Q&A
- [ ] Meeting notes
- [ ] Task generation
- [ ] Context-aware workspace assistant

AI remains a differentiator built on top of structured workspace data, not the first MVP feature.

---

# Definition of Done

A meaningful feature is complete when:

- [ ] Implemented
- [ ] Tests pass
- [ ] High-value edge cases covered
- [ ] Security reviewed
- [ ] Documentation updated
- [ ] Code reviewed/refactored where needed
- [ ] Git milestone committed
