# Session 101 — Registry HTTP API Boundary

## Objective

Establish the first HTTP API boundary for the registry system while preserving the existing separation between transport, application, repository, and domain layers.

## Architecture

The registry now follows this application flow:

```text
HTTP API
   │
   ▼
SystemHandler
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

The HTTP layer is responsible for transport concerns such as:

* decoding JSON request bodies
* reading URL path and query parameters
* selecting HTTP status codes
* encoding JSON responses
* translating transport-level errors into HTTP responses

The `SystemService` remains responsible for application-level registry operations and continues to depend on the `SystemRepository` abstraction.

## HTTP Handler

Added:

```text
internal/api/system_handler.go
```

The `SystemHandler` depends on `*registry.SystemService` and exposes the registry through HTTP.

The handler provides:

* `CreateSystem`
* `GetSystem`
* `ListSystems`
* `Routes`

The API layer does not access the registry implementation directly.

## Routes

The HTTP router currently exposes:

```text
POST /systems
GET  /systems
GET  /systems/{id}
```

The list endpoint also supports status filtering:

```text
GET /systems?status=active
GET /systems?status=inactive
```

An unsupported status value returns:

```text
400 Bad Request
```

## Create System

`POST /systems` accepts a JSON request containing:

```json
{
  "id": "system-1",
  "name": "Payments System",
  "description": "Core payments platform",
  "status": "active"
}
```

The request is decoded into an API-specific request structure and then converted into the existing `registry.System` domain type.

Registration is delegated to `SystemService`.

Successful registration returns:

```text
201 Created
```

Invalid JSON and validation errors return:

```text
400 Bad Request
```

## Get System

`GET /systems/{id}` retrieves a registered system through `SystemService.Get`.

A successful request returns:

```text
200 OK
```

with the system encoded as JSON.

A system that does not exist returns:

```text
404 Not Found
```

## List Systems

`GET /systems` delegates to `SystemService.List`.

The response is encoded as JSON.

When a status query parameter is provided, the handler delegates to:

```text
SystemService.ListByStatus
```

This keeps filtering logic outside the HTTP transport layer.

## Testing

Added:

```text
internal/api/system_handler_test.go
```

The tests cover:

* successful system creation
* invalid JSON
* validation errors
* successful system retrieval
* missing system retrieval
* HTTP route registration
* listing systems
* filtering systems by status
* invalid status filtering

The API package tests pass with:

```text
go test ./internal/api
```

The complete project test suite also passes with:

```text
go test ./...
```

## Design Boundary

Session 101 intentionally does not introduce:

* a database
* authentication or authorization
* middleware changes
* external HTTP frameworks
* application startup changes
* API versioning
* pagination
* request validation beyond the existing domain validation
* orchestration logic

The purpose of this session is to establish a small, testable HTTP transport boundary above the existing registry application layer.

## Result

The registry now has a concrete HTTP-facing boundary while retaining the existing architecture:

```text
HTTP transport
      ↓
application service
      ↓
repository abstraction
      ↓
in-memory registry
      ↓
domain model
```

This provides a foundation for future API expansion without moving business logic into HTTP handlers.
