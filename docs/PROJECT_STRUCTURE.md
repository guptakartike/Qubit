# 🏗️ Qubit Project Structure

> This document defines the directory structure, module boundaries, and organizational principles used throughout the Qubit codebase.

---

# Philosophy

Qubit follows a **Modular Monolith** architecture.

The codebase is organized around **business domains**, not technologies.

Each module owns its:

- API
- Business Logic
- Database Access
- Models
- Tests

Modules should communicate through well-defined interfaces instead of directly accessing each other's internals.

---

# Repository Structure

```
qubit/

├── frontend/
├── backend/

├── docs/
│   ├── adr/
│   ├── api/
│   ├── database/
│   ├── diagrams/
│   └── meetings/

├── .github/
│   ├── workflows/
│   ├── ISSUE_TEMPLATE/
│   └── PULL_REQUEST_TEMPLATE.md

├── README.md
├── ROADMAP.md
├── ARCHITECTURE.md
├── PROJECT_STRUCTURE.md
├── TECH_STACK.md
├── CONTRIBUTING.md
├── CHANGELOG.md
├── ENGINEERING.md
├── docker-compose.yml
└── .env.example
```

---

# Backend Structure

```
backend/

├── cmd/
│   └── api/
│       └── main.go

├── internal/

│   ├── auth/
│   ├── organization/
│   ├── workspace/
│   ├── document/
│   ├── task/
│   ├── chat/
│   ├── storage/
│   ├── calendar/
│   ├── notification/
│   ├── search/
│   └── ai/

├── pkg/

├── configs/

├── database/

├── migrations/

├── scripts/

├── tests/

└── go.mod
```

---

# Module Structure

Every module follows exactly the same layout.

```
workspace/

├── handler.go
├── service.go
├── repository.go
├── model.go
├── dto.go
├── routes.go
├── errors.go
└── tests/
```

---

## Responsibility

### handler.go

- HTTP Requests
- HTTP Responses
- Validation
- Status Codes

No business logic.

---

### service.go

Contains business rules.

Examples:

- Create Workspace
- Rename Workspace
- Archive Workspace

Business logic belongs here.

---

### repository.go

Responsible for database operations.

Examples:

- Find Workspace
- Save Workspace
- Delete Workspace

Only SQL/database access.

---

### model.go

Contains domain models.

Example:

```
Workspace
Organization
User
```

---

### dto.go

Defines request and response objects.

Example:

```
CreateWorkspaceRequest

UpdateWorkspaceRequest

WorkspaceResponse
```

---

### routes.go

Registers HTTP routes.

Example:

```
POST /workspaces

GET /workspaces/:id
```

---

### errors.go

Module-specific errors.

Example:

```
WorkspaceNotFound

WorkspaceAlreadyExists
```

---

### tests/

Module-specific tests.

---

# Frontend Structure

```
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
│   ├── styles/
│   └── assets/

├── package.json

└── vite.config.ts
```

---

# Feature Structure (Frontend)

Every feature should remain self-contained.

Example:

```
features/

authentication/

workspace/

documents/

tasks/

chat/

calendar/
```

Each feature may contain:

```
authentication/

components/

hooks/

api/

types/

pages/

utils/
```

---

# Shared Components

Global reusable components belong here.

```
components/

Button

Modal

Input

Avatar

Navbar

Sidebar

Loader
```

---

# API Layer

The frontend never calls the backend directly from UI components.

```
services/

auth.ts

workspace.ts

tasks.ts

documents.ts
```

All HTTP requests originate from the service layer.

---

# State Management

Global application state:

```
store/
```

Server state:

```
TanStack Query
```

Local component state:

```
React useState()
```

---

# Documentation Structure

```
docs/

adr/
database/
api/
diagrams/
meetings/
```

---

# Naming Conventions

## Go

Packages

```
workspace
organization
notification
```

Use lowercase.

---

## Files

```
workspace_service.go

workspace_repository.go

workspace_handler.go
```

Use descriptive names.

---

## React Components

```
WorkspaceCard.tsx

Sidebar.tsx

TaskModal.tsx
```

PascalCase.

---

## Variables

camelCase

```
workspaceID

userRole

taskStatus
```

---

## Constants

UPPER_SNAKE_CASE

```
MAX_FILE_SIZE

JWT_EXPIRATION
```

---

# Module Communication

Modules should never access another module's repository directly.

Correct:

```
Workspace Service

↓

Organization Service Interface
```

Incorrect:

```
Workspace Repository

↓

Organization Repository
```

This keeps modules loosely coupled.

---

# Dependency Direction

```
Handler

↓

Service

↓

Repository

↓

Database
```

Never reverse this dependency.

Repositories should never call services.

---

# General Rules

- Keep modules independent.
- One responsibility per package.
- Business logic belongs in services.
- Database logic belongs in repositories.
- HTTP logic belongs in handlers.
- Never expose database models directly to clients.
- Prefer composition over inheritance.
- Write tests alongside features.

---

# Future Scalability

The project structure is intentionally designed so that individual modules can later be extracted into independent microservices with minimal refactoring.

---

# Guiding Principle

> **Organize by business capability, not by technology.**

The structure should make it immediately obvious **what the software does**, not just **how it is implemented**.