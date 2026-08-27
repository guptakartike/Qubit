# Qubit — Architecture & Engineering Decisions

This file records meaningful technical decisions and why they were made.

---

## ADR-001 — Modular Monolith

**Status:** Accepted

### Decision

Qubit will initially use a modular monolith instead of microservices.

### Why

The project is still being built and microservice operational complexity is not justified. Clear module boundaries provide most of the architectural benefits needed now while keeping local development, testing, and deployment simple.

### Trade-off

Modules do not scale or deploy independently in the short term. That trade-off is intentional.

---

## ADR-002 — PostgreSQL

**Status:** Accepted

### Decision

PostgreSQL is the primary relational database.

### Why

Qubit contains strongly relational data such as users, organizations, memberships, permissions, tasks, documents, and messages. PostgreSQL provides transactions, constraints, indexing, and mature tooling.

---

## ADR-003 — Application-Generated UUID Identifiers

**Status:** **Needs Review**

### Current Implementation

The current Go implementation generates IDs with the UUID library's `uuid.New()`.

### Documentation Correction

Earlier documentation stated that UUIDv7 was already implemented. That is not currently verified by the code.

### Decision

Do **not** claim UUIDv7 support until the implementation explicitly generates UUIDv7 and tests verify it.

### Next Action

During the registration architecture review, decide whether to keep the current UUID strategy or deliberately move to UUIDv7.

---

## ADR-004 — Canonical Lowercase Emails

**Status:** Accepted

### Decision

Email addresses are normalized to lowercase before persistence.

### Why

Authentication should operate on a canonical representation. PostgreSQL also enforces `UNIQUE(email)` as the final integrity boundary.

---

## ADR-005 — Separate Password Credentials

**Status:** Accepted

### Decision

Password hashes are stored in a separate one-to-one `password_credentials` table rather than directly in `users`.

### Why

`users` represents identity; credential tables represent authentication mechanisms. This makes future OAuth and session-related data easier to model without polluting the identity table.

---

## ADR-006 — Argon2id Password Hashing

**Status:** Accepted

### Decision

Passwords are hashed using Argon2id.

### Why

Passwords require a one-way password hashing algorithm designed to be expensive for attackers. The implementation uses a random salt and an encoded Argon2id representation.

Plaintext passwords must never be stored or logged.

---

## ADR-007 — Soft Account Deletion

**Status:** Accepted

### Decision

Normal account deletion uses `deleted_at` rather than immediately physically deleting the user.

### Why

Qubit will contain historical relationships such as tasks, messages, documents, and activity logs. Historical identity may need to remain referentially useful.

Permanent deletion can be treated as a separate operation if required.

---

## ADR-008 — Authentication Providers

**Status:** Accepted

Qubit will support:

- Email/password
- Google OAuth
- GitHub OAuth

OAuth identities will be modeled separately from password credentials.

---

## ADR-009 — Atomic Account Creation

**Status:** Accepted

### Decision

Creating a user and the associated password credential happens inside one PostgreSQL transaction.

### Why

The account must not exist without its required credential. If credential creation fails, user creation must roll back.

---

## ADR-010 — Registration API Layering

**Status:** Accepted

### Decision

Registration follows:

```text
Handler → Service → Repository → PostgreSQL
```

### Why

HTTP concerns, application logic, and persistence concerns have different responsibilities and test requirements. Keeping them separate allows handler and service behavior to be tested without a real database while still using integration tests for PostgreSQL-specific behavior.

---

## ADR-011 — PostgreSQL Is Final Authority for Email Uniqueness

**Status:** Accepted

### Decision

The application may normalize and validate email input, but PostgreSQL's unique constraint remains the final authority for duplicate prevention.

### Why

Application-level pre-checks alone are vulnerable to race conditions. The database constraint provides atomic enforcement under concurrency.

Duplicate-key SQLSTATE `23505` is translated into the domain error `auth.ErrEmailAlreadyExists` where appropriate.

---

## ADR-012 — Resume-First, Production-Oriented Scope

**Status:** Accepted

### Decision

Qubit is currently being built as a local/resume project rather than a live multi-user SaaS product.

### Why

The objective is to demonstrate a small number of strong engineering capabilities. We will still use production-oriented architecture, security, testing, and documentation, but will not add infrastructure merely for the appearance of scale.

### Consequence

Prioritize:

- Correctness
- Security
- Testing
- Architecture
- Demonstrable engineering depth

Avoid premature:

- Microservices
- Kubernetes
- Kafka
- Event sourcing
- CQRS
- Multi-region infrastructure
- Multiple databases
