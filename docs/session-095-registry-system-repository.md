# Session 095 — Registry System Repository

## Overview

Session 095 introduces the `SystemRepository` abstraction for the Financial Systems Registry.

The Registry already provides in-memory storage and management of registered systems. Rather than creating a second storage implementation, this session establishes an interface around the existing Registry behavior.

This creates a persistence boundary that can later support alternative implementations without requiring callers to depend directly on the current in-memory storage mechanism.

## Changes

### SystemRepository interface

Added `internal/registry/system_repository.go`.

The `SystemRepository` interface defines the core storage operations currently provided by `Registry`:

* `Register`
* `Get`
* `List`
* `ListByStatus`

The interface represents the contract that future persistence implementations can satisfy.

### Registry implementation

The existing `Registry` remains responsible for the current in-memory implementation.

A compile-time assertion confirms that `Registry` satisfies `SystemRepository`:

```go
var _ SystemRepository = (*Registry)(nil)
```

No duplicate repository implementation or additional storage map was introduced.

### Repository contract tests

Added `internal/registry/system_repository_test.go`.

The tests verify that:

* `Registry` can be used through the `SystemRepository` interface.
* Systems can be registered through the interface.
* Systems can be retrieved through the interface.
* Systems can be listed through the interface.
* Systems can be filtered by status through the interface.

## Architectural Direction

The registry architecture now separates the domain-facing repository contract from its current storage implementation:

```text
System Domain
     |
     v
SystemRepository
     |
     v
Registry
(in-memory storage)
```

This establishes a boundary for future persistence work:

```text
SystemRepository
     |
     +-------------------------+
     |                         |
     v                         v
Registry                Future Database
(in-memory)             implementation
```

The current session intentionally does not introduce a database.

Keeping the implementation in memory allows the project to establish the interface, dependency boundary, and test behavior before introducing persistence infrastructure.

## Testing

The Registry package test suite passes:

```text
go test ./internal/registry
```

All existing Registry, System, lifecycle, validation, and repository contract tests pass.

## Result

The Registry now has an explicit repository contract while retaining its existing in-memory implementation.

This provides a foundation for future service-layer and persistence architecture without prematurely coupling the domain to a specific database technology.
