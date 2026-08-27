# Qubit — Engineering Principles

Qubit is built to demonstrate professional engineering practices without unnecessary complexity.

---

# Core Principles

Every meaningful feature should satisfy:

- Correctness
- Security
- Testability
- Maintainability
- Clear ownership
- Appropriate observability
- Reasonable scalability
- Documentation

---

# Before Writing Code

Answer:

1. What problem are we solving?
2. What is explicitly out of scope?
3. Which module owns the behavior?
4. What data is required?
5. What are the failure modes?
6. What are the security implications?
7. What database constraints are required?
8. How will it be tested?
9. Does this introduce unnecessary complexity?
10. How will the decision affect future features?

---

# Layer Responsibilities

## Handler

Responsible for:

- HTTP parsing/binding
- Request-shape validation
- Calling application services
- HTTP status codes
- Response serialization

Handlers should not contain core business logic or SQL.

## Service

Responsible for:

- Business/application rules
- Validation coordination
- Normalization
- Security-sensitive orchestration
- Coordinating repositories

## Repository

Responsible for:

- SQL/database access
- Persistence-specific error translation where needed
- Transactional database operations

Repositories should not know about HTTP.

---

# Testing Strategy

Use the cheapest test that proves the behavior:

```text
Unit Tests
    ↓
Pure logic / service behavior

Handler Tests
    ↓
HTTP contract and error mapping

Integration Tests
    ↓
Actual PostgreSQL behavior

End-to-End Verification
    ↓
Real request through the application
```

Do not mock database behavior when the purpose of the test is to prove actual PostgreSQL behavior.

---

# Security Standards

- Never store plaintext passwords.
- Use Argon2id for passwords.
- Never log passwords or password hashes.
- Use parameterized SQL.
- Enforce critical uniqueness and integrity in PostgreSQL.
- Do not expose internal database errors to API clients.
- Use constant-time comparison for password hashes.
- Treat authentication and authorization as backend-enforced concerns.

---

# Complexity Rule

Do not add technology merely because it sounds production-grade.

Avoid premature:

- Microservices
- Kubernetes
- Kafka
- Event sourcing
- CQRS
- Multi-region architecture
- Multiple databases
- Service meshes

A simple architecture that is correct and well-tested is preferable.

---

# Documentation Rule

When implementation changes architecture or behavior, update the relevant documentation in the same milestone.
