# Session 096 — Registry System Service

## Overview

Session 096 introduces the `SystemService` application layer for the Financial Systems Registry.

Session 095 established the `SystemRepository` interface as the storage boundary for registered systems. This session builds on that boundary by introducing a service that depends on the repository contract rather than directly on the `Registry` implementation.

The service layer provides an application-facing entry point for registry operations and creates a clean boundary for future API and orchestration work.

## Changes

### SystemService

Added `internal/registry/system_service.go`.

The `SystemService` contains a `SystemRepository` dependency and exposes the current registry operations through an application-level interface:

* `Register`
* `Get`
* `List`
* `ListByStatus`

The service delegates these operations to the configured repository.

This keeps the service independent of the current in-memory `Registry` implementation.

### Constructor

Added `NewSystemService`.

The constructor accepts a `SystemRepository`:

```go
func NewSystemService(repository SystemRepository) *SystemService
```

This allows the service to work with any implementation that satisfies the repository contract.

The service therefore does not require callers to provide a concrete `*Registry`.

## Dependency Direction

The registry architecture now separates domain, storage, and application concerns:

```text
System Domain
     |
     v
SystemRepository
     |
     v
SystemService
     |
     v
Future API / Application Entry Points
```

The concrete storage implementation remains:

```text
SystemRepository
     |
     v
Registry
(in-memory storage)
```

The service depends on the `SystemRepository` abstraction rather than on `Registry` directly.

This creates the following boundary:

```text
                 +----------------+
                 |  SystemService |
                 +----------------+
                         |
                         | depends on
                         v
                 +----------------+
                 |SystemRepository|
                 +----------------+
                         |
                         v
                 +----------------+
                 |    Registry    |
                 | in-memory store|
                 +----------------+
```

A future database-backed repository can therefore be introduced without requiring the service to change its dependency type.

## Testing

Added `internal/registry/system_service_test.go`.

The tests use a fake implementation of `SystemRepository` rather than the concrete `Registry`.

The tests verify that:

* `SystemService` is created with the supplied repository.
* `Register` delegates to the repository.
* repository registration errors are returned by the service.
* `Get` delegates to the repository.
* `Get` returns the repository result.
* `List` delegates to the repository.
* `ListByStatus` delegates to the repository.
* the fake repository satisfies the `SystemRepository` contract.

This confirms that the service is tested against the repository abstraction rather than a specific storage implementation.

## Architectural Direction

The registry subsystem now has three distinct layers:

```text
System
Domain model and validation
        |
        v
SystemRepository
Storage contract
        |
        v
SystemService
Application/use-case boundary
        |
        v
Future API
```

The current concrete storage remains:

```text
SystemService
      |
      v
SystemRepository
      |
      v
Registry
```

Future persistence can replace the repository implementation without changing the service dependency:

```text
                  SystemService
                       |
                       v
               SystemRepository
                  /          \
                 /            \
                v              v
           Registry       DatabaseRepository
          (in-memory)        (future)
```

The session intentionally does not introduce HTTP handlers, database code, or additional persistence infrastructure.

## Testing

The registry package tests pass with:

```text
go test ./internal/registry
```

The complete project test suite passes with:

```text
go test ./...
```

## Result

The Financial Systems Registry now has an explicit application-level service boundary above the repository layer.

The architecture has progressed from:

```text
System
  |
  v
Registry
```

to:

```text
System
  |
  v
SystemRepository
  |
  v
SystemService
```

This establishes a foundation for future API endpoints, application use cases, and alternative persistence implementations while keeping the current implementation small and testable.
