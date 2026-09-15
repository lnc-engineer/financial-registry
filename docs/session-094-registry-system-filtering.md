# Session 094 — Registry System Filtering

## Overview

Session 094 extends the financial systems registry with system filtering by lifecycle status.

The registry previously supported registering systems, retrieving individual systems by ID, and listing all registered systems. This session introduces a focused discovery capability that allows callers to retrieve systems matching a specific `SystemStatus`.

## Changes

### Registry filtering

Added:

```go
func (r *Registry) ListByStatus(status SystemStatus) []System
```

The method returns all registered systems whose status matches the requested status.

Supported statuses currently include:

* `SystemStatusActive`
* `SystemStatusInactive`

### Empty-result behavior

`ListByStatus` returns a non-nil empty slice when no registered systems match the requested status.

This keeps collection behavior predictable for callers.

### Registry isolation

The filtering method returns system values in a newly allocated slice.

Changes made to the returned collection do not modify the systems stored inside the registry.

## Tests

Added coverage for:

* returning systems matching the requested status
* excluding systems with a different status
* returning a non-nil empty collection when no systems match
* preserving registry state when the returned collection is modified

## Design Direction

The registry is beginning to provide explicit discovery capabilities while keeping the domain API small and predictable.

Filtering is intentionally limited to system status in this session. More advanced discovery criteria can be introduced later when the registry domain requires them.

The current progression is:

```text
System Domain
    ↓
System Validation
    ↓
System Lifecycle
    ↓
Registry Collection
    ↓
System Management
    ↓
System Filtering
```

This maintains a small and reliable registry core while leaving room for future persistence, API, execution, observability, and trust-layer integrations.

## Verification

The complete Go test suite passes:

```text
go test ./...
```

All existing and new registry tests pass successfully.
