# Changelog

All notable Qubit changes are documented here.

---

## Unreleased

### Added

- Registration request/response DTOs
- Registration validation
- Email normalization
- Argon2id password hashing and verification
- Authentication domain errors
- PostgreSQL user repository
- Transactional user + password credential creation
- Duplicate-email error translation
- Registration service
- Gin registration handler
- Registration service tests
- PostgreSQL repository integration tests
- Registration handler tests
- Manual registration verification

### Fixed

- PostgreSQL duplicate-email errors now map to the authentication domain error
- Invalid registration input now maps to HTTP `400 Bad Request`
- Credential creation failure correctly rolls back the user insert

### Documentation

- Updated project architecture to reflect the modular monolith
- Updated authentication flow and registration state
- Corrected UUIDv7 documentation to reflect the actual current implementation
- Added current handover and milestone state

---

## v0.1.0 — Initial Foundation

### Added

- Initial project structure
- Development roadmap
- Architecture documentation
- Backend setup
- Frontend direction
- PostgreSQL foundation
- Initial migration structure

---

## Next

The next milestone after the registration architecture review is Login.
