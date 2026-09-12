# Session 091 — Registry System Lifecycle

## Goal

Add explicit lifecycle operations to the registry `System` domain.

## Changes

Added lifecycle methods to `registry.System`:

- `Activate()`
- `Deactivate()`

The lifecycle rules are:

| Current status | Operation | Result |
|---|---|---|
| active | Activate | error; remains active |
| active | Deactivate | becomes inactive |
| inactive | Activate | becomes active |
| inactive | Deactivate | error; remains inactive |

## Design

Lifecycle transitions are owned by the `System` domain rather than allowing callers to directly manage status changes.

This keeps lifecycle rules close to the domain model and provides a controlled foundation for future registry services, APIs, persistence, and execution integration.

## Tests

Added coverage for:

- Activating an inactive system
- Attempting to activate an already-active system
- Deactivating an active system
- Attempting to deactivate an already-inactive system
- Verifying that failed transitions do not change the current status

## Verification

The following commands pass:

```bash
go test ./internal/registry
go test ./...
```

## Outcome

The registry `System` domain now supports controlled active/inactive lifecycle transitions with explicit error handling and test coverage.
