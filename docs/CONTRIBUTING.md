# Contributing to Qubit

Qubit follows a lightweight professional development workflow.

---

## Git Workflow

Prefer feature branches rather than committing directly to `main`.

Examples:

```text
feature/auth-registration
feature/auth-login
feature/workspaces
feature/task-board
```

---

## Commit Convention

Use conventional-style commit messages:

```text
feat(auth): implement user registration
feat(auth): implement login
fix(auth): translate duplicate email error
refactor(auth): simplify registration service
test(auth): add registration rollback coverage
docs: update authentication roadmap
chore: update dependencies
```

---

## Feature Workflow

```text
Plan
 ↓
Design
 ↓
Implement
 ↓
Test
 ↓
Review
 ↓
Document
 ↓
Commit
```

---

## Pull Request Checklist

- [ ] Code compiles
- [ ] `go test ./...` passes
- [ ] Relevant integration tests pass
- [ ] No hardcoded secrets
- [ ] Error handling reviewed
- [ ] Security implications reviewed
- [ ] Documentation updated
- [ ] Architecture decisions recorded if changed
- [ ] Changes are focused and explainable
