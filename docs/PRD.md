# Qubit — Product Requirements Document

## 1. Problem Statement

Small and medium-sized software teams often use multiple disconnected tools for communication, documentation, task management, and project coordination. Context becomes fragmented across Slack, Notion, Jira, Trello, Google Drive, and similar products.

Qubit aims to provide a unified collaborative workspace where teams can communicate, manage projects, create documents, and organize knowledge without constantly switching between tools.

---

## 2. Target Users

- Small and medium-sized software teams
- Startups
- Engineering teams
- Product teams
- Project managers
- Team leads
- Developers and knowledge workers

---

## 3. Product Vision

Qubit should connect collaboration context rather than simply place unrelated tools inside one UI.

Example:

```text
Chat discussion
      ↓
Task
      ↓
Document
      ↓
Project activity
```

The long-term differentiator is **context-aware AI** operating on structured workspace data.

---

## 4. MVP Scope

The MVP should focus on a small set of capabilities that form one coherent team workflow:

### Authentication

- Email/password registration
- Login
- Sessions/tokens
- Email verification
- Password reset
- OAuth as an extension

### Organizations / Workspaces

- Organization creation
- Membership
- Invitations
- Roles and permissions
- Workspace creation
- Workspace membership

### Collaboration

- Channels
- Messaging
- Documents
- Tasks
- Basic project activity

### Supporting capabilities

- File storage
- Notifications
- Search

Advanced versions of these features remain later milestones.

---

## 5. Explicitly Out of Scope for the First Production-Oriented MVP

- Native mobile applications
- Desktop applications
- Google Docs-level collaborative editing
- Full Jira replacement
- Full Slack replacement
- Large-scale video conferencing
- Complex enterprise billing
- Marketplace/app ecosystem
- Advanced AI agent workflows
- Complex workflow automation
- Microservice architecture
- Multi-region deployment
- Globally distributed database architecture

---

## 6. Product Differentiation

Qubit should initially differentiate through **unified context**.

Users should be able to move between related conversations, tasks, documents, and project activity without losing context.

Long-term AI examples:

```text
"Summarize what happened on Project X this week."
"Create tasks from this discussion."
"What decisions were made about authentication?"
"What are the blockers for the current sprint?"
```

AI is a future capability built on top of workspace data, not the first milestone.

---

## 7. Technical Requirements

### Frontend

- React
- TypeScript
- Vite
- Tailwind CSS
- shadcn/ui
- TanStack Query
- React Router
- Zustand

### Backend

- Go
- Gin
- PostgreSQL
- Redis later
- Modular monolith

### Database

PostgreSQL with:

- Foreign keys
- Unique constraints
- Appropriate indexes
- Transactions for multi-step operations
- Versioned migrations
- Soft deletion where historical relationships require preservation

### Authentication

- Email/password
- Argon2id
- Short-lived access tokens
- Refresh-token/session management
- Email verification
- Session revocation
- Google OAuth
- GitHub OAuth
- Backend-enforced authorization

### Realtime

WebSockets for features that actually require live communication.

### Storage

MinIO / S3-compatible storage for development and early deployment, behind an abstraction.

### Infrastructure

- Docker / Docker Compose
- GitHub Actions
- Environment configuration
- Structured logging
- Automated tests
- OpenAPI / Swagger

---

## 8. Current Implementation Status

### Authentication

```text
Registration        ✅ Implemented + verified
Architecture review ⏳ Current
Login               ⏳
Sessions/tokens     ⏳
Email verification  ⏳
Password reset      ⏳
OAuth               ⏳
```

### Current Registration API

```text
POST /api/v1/auth/register
```

The current registration flow validates input, normalizes email, hashes the password with Argon2id, creates the user and password credential atomically, and returns safe user data.

---

## 9. Success Criteria

The project succeeds as a resume/portfolio project when it demonstrates:

- Strong modular architecture
- Secure authentication
- Meaningful automated tests
- Relational database design
- Transaction handling
- Clear API contracts
- Authorization design
- Real-time capability
- Practical system design
- Clean documentation

For the eventual product MVP, reliability targets include:

- Critical backend flows covered by automated tests
- Reproducible database migrations
- No critical authentication/security vulnerabilities
- Containerized deployment when deployment becomes a goal

---

## 10. Open Product Decisions

Still to be finalized when the relevant features are designed:

1. Organization vs workspace permission inheritance
2. Company-domain behavior and organization association
3. Document real-time collaboration model
4. Search architecture beyond PostgreSQL
5. AI provider/retrieval/tenant-isolation architecture
6. Exact session and token model

These decisions should be made close to implementation rather than prematurely.

---

## Product Principle

> **Quality > Feature Count**

Every major feature should be:

- Properly architected
- Secure
- Tested
- Documented
- Maintainable
- Observable where appropriate
- Defensible in a technical interview
