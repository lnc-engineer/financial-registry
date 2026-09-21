# Session 100 — Registry Service Validation Boundary

## Overview

Session 100 establishes explicit validation at the registry service boundary.

`SystemService.Register` now validates a system through `System.Validate()` before delegating registration to `SystemRepository.Register`.

## Objective

The service layer now:

* validates systems before repository access,
* returns domain validation errors unchanged,
* prevents invalid systems from reaching the repository, and
* preserves existing repository error propagation.

## Architecture

```text
SystemService.Register
        |
        v
System.Validate
        |
        v
SystemRepository.Register
        |
        v
Registry.Register
```

`Registry.Register` retains its existing validation so invalid systems cannot be stored even when the registry is accessed directly.

## Tests

The service test suite verifies:

* missing ID is rejected before repository access,
* missing name is rejected before repository access,
* invalid status is rejected before repository access,
* valid registration continues to delegate to the repository, and
* repository errors continue to propagate.

## Design Decision

The service layer does not duplicate validation rules.

It reuses `System.Validate()` and establishes the application-level boundary where domain validation occurs before repository access.

The repository implementation retains its own validation as a storage invariant.

## Verification

The implementation was formatted with `gofmt`.

The registry tests passed:

```bash
go test ./internal/registry
```

The full test suite passed:

```bash
go test ./...
```

`git diff --check` reported no whitespace errors.

## Outcome

Session 100 strengthens the separation between domain validation, application service operations, and repository operations.

This provides a clearer foundation for future API and orchestration layers without introducing persistence or transport concerns into the current registry service.
