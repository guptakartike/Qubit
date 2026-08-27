# Qubit — Feature Development Log

Track meaningful features from planning through verification.

---

## FEAT-001 — User Identity & Account Schema

**Status:** Verified

### Goal

Create the foundational user identity model required for authentication and future platform features.

### Implemented

- User identity
- Canonical email
- Account status
- Email verification timestamp
- Soft deletion field
- Creation/update timestamps
- UUID primary key
- Unique email constraint

### Database

```text
001_create_users
```

Table:

```text
users
```

### Verification

PostgreSQL schema was inspected successfully.

---

## FEAT-002 — Password Credentials

**Status:** Verified

### Goal

Separate authentication credentials from general user identity.

### Implemented

- One-to-one password credential relationship
- Password hash storage
- Foreign key to users
- Creation/update timestamps
- `ON DELETE CASCADE`

### Database

```text
002_create_password_credentials
```

Table:

```text
password_credentials
```

### Security

- Argon2id password hashing
- Random salt
- No plaintext password persistence

### Verification

Migration applied successfully and PostgreSQL schema inspected.

---

## FEAT-003 — User Registration

**Status:** Implemented and Verified — architecture review pending

### Goal

Allow a user to create a Qubit account securely through an HTTP API.

### Flow

```text
HTTP Request
    ↓
Gin Handler
    ↓
Registration Service
    ↓
Validate input
    ↓
Normalize email
    ↓
Generate user ID
    ↓
Hash password
    ↓
Repository
    ↓
BEGIN
    ↓
INSERT users
    ↓
INSERT password_credentials
    ↓
COMMIT
    ↓
Safe HTTP response
```

### API

```text
POST /api/v1/auth/register
```

### Success

```text
201 Created
```

Response contains:

- ID
- Name
- Email
- Status

The password and password hash are never returned.

### Validation

Implemented:

- Required name
- Required email
- Name length limit
- Email length limit
- Email syntax validation
- Required password
- Password length validation

### Error Mapping

```text
Invalid JSON       → 400
Invalid input      → 400
Duplicate email    → 409
Unexpected failure → 500
```

### Database Behavior

User and password credential creation occur in one transaction.

A credential failure causes the user insert to roll back.

PostgreSQL duplicate-key error `23505` is translated to `ErrEmailAlreadyExists` where applicable.

### Tests

#### Service

- Success
- Validation failure
- Hashing failure
- Repository failure
- Unique ID generation

#### Repository

- Successful creation
- Rollback on credential failure
- Duplicate email

#### Handler

- Success
- Invalid JSON
- Invalid input
- Duplicate email
- Internal error

### Verification

The full suite currently passes:

```bash
go test ./...
```

The registration endpoint was also manually tested against the running API and the resulting PostgreSQL records were inspected.

### Current Next Step

Perform the final registration architecture/security/code-quality review before starting Login.

---

# Feature Development Rule

Every meaningful feature should move through:

```text
Planned
   ↓
Designed
   ↓
Implemented
   ↓
Tested
   ↓
Verified
   ↓
Committed
```

Do not mark a feature complete merely because the code compiles.
