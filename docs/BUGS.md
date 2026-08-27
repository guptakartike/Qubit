# Qubit — Bug Log

Track meaningful bugs from discovery through verification. Do not use this file for ordinary typos or temporary local mistakes.

---

## BUG-001 — Migration CLI SSL Connection

**Status:** Verified Fixed

### Symptoms

`golang-migrate` failed against local PostgreSQL with an SSL-related error.

### Root Cause

The local PostgreSQL server does not require SSL while the connection string did not explicitly disable it.

### Fix

Use `sslmode=disable` for local PostgreSQL CLI commands.

### Verification

Migration version reached `2` and PostgreSQL showed:

```text
schema_migrations
users
password_credentials
```

### Production Note

`sslmode=disable` is a local-development setting and must not automatically be used in production.

---

## BUG-002 — DATABASE_URL Empty in Migration CLI

**Status:** Expected Development Configuration

### Symptoms

Running a migration command with:

```bash
migrate -path migrations -database "$DATABASE_URL" version
```

returned an empty-URL error.

### Root Cause

`.env` is loaded by the Go application through `godotenv`; loading `.env` does not automatically export variables into the shell process used by external CLI tools.

### Current Approach

Use an explicit local database URL for migration CLI commands when necessary.

### Future Improvement

Document a single, consistent local-development configuration workflow.

---

## BUG-003 — PostgreSQL Duplicate Email Was Not Translated

**Status:** Verified Fixed

### Symptoms

The duplicate-registration integration test received the raw PostgreSQL error:

```text
ERROR: duplicate key value violates unique constraint "users_email_key"
SQLSTATE 23505
```

instead of:

```text
auth.ErrEmailAlreadyExists
```

### Root Cause

The repository wrapped the PostgreSQL error but did not inspect the PostgreSQL error code and translate the duplicate-email constraint violation into the domain error expected by the service/handler layers.

### Fix

Translate the relevant PostgreSQL unique-constraint violation into `auth.ErrEmailAlreadyExists` while preserving unrelated database errors as internal failures.

### Verification

The repository duplicate-email test passed and the full suite returned green.

---

## BUG-004 — Handler Invalid Input Returned 500

**Status:** Verified Fixed

### Symptoms

The handler test for invalid registration input returned `500` instead of the expected `400`.

### Root Cause

The handler's error mapping did not recognize the validation/domain error returned by the service correctly.

### Fix

Align service error propagation and handler `errors.Is()` matching with `auth.ErrInvalidInput`.

### Verification

The handler test suite passed, including invalid-input behavior, and the full `go test ./...` suite passed.

---

## Open Issues

No known failing registration tests remain at the current milestone.

The registration architecture review may identify additional improvements; new issues should be recorded here only if they are meaningful engineering issues.
