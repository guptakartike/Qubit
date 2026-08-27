# Qubit — Execution Flow

This document records how requests travel through the application.

---

# General Backend Flow

```text
HTTP Request
     ↓
Gin Router
     ↓
Middleware (when configured)
     ↓
Handler
     ↓
Service
     ↓
Repository
     ↓
PostgreSQL
     ↓
Repository
     ↓
Service
     ↓
Handler
     ↓
HTTP Response
```

---

# Registration Flow — Current

```text
POST /api/v1/auth/register
             ↓
      Gin RegistrationHandler
             ↓
       ShouldBindJSON
             ↓
    RegistrationService.Register
             ↓
       Validate request
             ↓
       Normalize email
             ↓
       Generate user ID
             ↓
       Hash password
             ↓
       Repository.CreateUser
             ↓
          BEGIN
          /    \
         /      \
   users     password_credentials
         \      /
          \    /
          COMMIT
             ↓
      Registration response
             ↓
         201 Created
```

---

# Registration Failure Flow

```text
Repository.CreateUser
        ↓
BEGIN
        ↓
INSERT users
        ↓
INSERT password_credentials fails
        ↓
ROLLBACK
        ↓
Return wrapped error
        ↓
Service propagates domain/application error
        ↓
Handler maps to HTTP response
```

The database must not retain a partially-created account.

---

# Duplicate Email Flow

```text
Register request
      ↓
Normalize email
      ↓
INSERT users
      ↓
PostgreSQL UNIQUE(email)
      ↓
SQLSTATE 23505
      ↓
Repository translates duplicate-email condition
      ↓
auth.ErrEmailAlreadyExists
      ↓
Handler
      ↓
409 Conflict
```

The database constraint remains the final authority for uniqueness, avoiding race-condition-prone application-only checks.

---

# Password Flow

```text
Registration password
        ↓
Argon2id
        ↓
Random salt
        ↓
Encoded password hash
        ↓
password_credentials.password_hash
```

Verification later follows:

```text
Login password
      ↓
Parse encoded Argon2id hash
      ↓
Derive candidate hash
      ↓
Constant-time comparison
      ↓
Authenticated / rejected
```

---

# Planned Login Flow

```text
POST /api/v1/auth/login
        ↓
Validate request
        ↓
Normalize email
        ↓
Find user
        ↓
Check account status
        ↓
Find password credential
        ↓
Verify Argon2id hash
        ↓
Create session
        ↓
Issue access token
        ↓
Issue/store refresh token
        ↓
Return authenticated response
```

This flow is **planned, not implemented yet**.

---

# Planned Authorization Flow

```text
Request
  ↓
Access Token
  ↓
Authenticate User
  ↓
Resolve Organization / Workspace Membership
  ↓
Resolve Role / Permission
  ↓
Allow or Reject
```

Authentication answers:

```text
Who are you?
```

Authorization answers:

```text
Are you allowed to do this?
```

---

# Rule

When a new module or important endpoint is implemented, update this document with the **actual** execution path. Planned behavior must be clearly labeled as planned.
