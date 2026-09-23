# Session 102 — Registry HTTP Service Boundary

## Objective

Refine the HTTP API boundary by decoupling the transport layer from the concrete `SystemService` implementation.

Session 101 established the first concrete HTTP API for the registry. Session 102 introduces an application-service interface within the API package so that the HTTP handler depends on the operations it requires rather than directly on the concrete registry service implementation.

## Architecture

The registry now follows this application flow:

```text
HTTP API
   │
   ▼
SystemHandler
   │
   ▼
systemService
   │
   ▼
SystemService
   │
   ▼
SystemRepository
   │
   ▼
Registry
   │
   ▼
System
```

The HTTP handler now depends on the `systemService` interface.

This keeps the transport layer focused on HTTP concerns while allowing the application service implementation to remain replaceable behind a small contract.

## HTTP Service Interface

Added:

```text
internal/api/system_service.go
```

The file defines:

```go
type systemService interface {
	Register(system registry.System) error
	Get(id string) (registry.System, bool)
	List() []registry.System
	ListByStatus(status registry.SystemStatus) []registry.System
}
```

The interface represents the application operations required by the HTTP API.

The HTTP package does not need to know how those operations are implemented.

## Handler Dependency

`SystemHandler` previously depended directly on:

```go
*registry.SystemService
```

It now depends on:

```go
systemService
```

The constructor was updated accordingly:

```go
func NewSystemHandler(service systemService) *SystemHandler
```

The existing HTTP behavior remains unchanged.

The handler continues to provide:

* `CreateSystem`
* `GetSystem`
* `ListSystems`
* `Routes`

## Compile-Time Verification

The API package includes a compile-time assertion:

```go
var _ systemService = (*registry.SystemService)(nil)
```

This verifies that `registry.SystemService` satisfies the service contract required by the HTTP layer.

If the registry service stops implementing one of the required operations, the project will fail to compile rather than silently violating the API boundary.

## Design Boundary

Session 102 does not introduce:

* a database
* authentication or authorization
* middleware changes
* external HTTP frameworks
* API versioning
* new HTTP endpoints
* changes to existing request or response formats
* changes to registry business rules
* orchestration logic

The purpose of this session is specifically to establish a small dependency boundary between HTTP transport and application services.

## Testing

The existing HTTP API tests continue to pass after introducing the interface boundary.

The API package was verified with:

```text
go test ./internal/api
```

The complete project test suite was also verified with:

```text
go test ./...
```

Both commands completed successfully.

## Result

The registry HTTP layer is now coupled to an explicit application-service contract rather than a concrete service implementation:

```text
HTTP transport
      ↓
service interface
      ↓
application service
      ↓
repository abstraction
      ↓
in-memory registry
      ↓
domain model
```

This provides a cleaner separation of concerns and establishes a foundation for future API testing, alternative service implementations, and continued evolution of the registry application architecture.
