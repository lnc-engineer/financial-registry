# Session 103 — Registry API Composition

## Overview

Session 103 introduces the application composition layer for the registry HTTP API.

The previous sessions established separate boundaries for:

* Registry domain storage
* System repository
* System service
* HTTP API
* HTTP service

This session connects those boundaries into a single application-level construction path.

## Objective

The objective of this session is to compose the existing registry components without moving HTTP, storage, or domain responsibilities into the wrong layer.

The application package is responsible for constructing the dependencies required by the HTTP API.

## Architecture

The resulting dependency flow is:

HTTP Server
↓
HTTP Handler
↓
System Service
↓
System Repository
↓
Registry

The application composition flow is:

NewHTTPServer()
↓
NewHTTPHandler()
├── NewRegistry()
├── NewSystemService(repository)
└── NewSystemHandler(service)
    ↓
    Routes()

## Application Package

A new `internal/app` package was introduced.

### HTTP handler composition

`NewHTTPHandler` constructs:

1. A new registry.
2. A registry system service using that registry as its repository.
3. An HTTP system handler using the system service.
4. The HTTP routes exposed by the handler.

The function returns the resulting `http.Handler`.

This keeps dependency construction outside the HTTP API package.

### HTTP server construction

`NewHTTPServer` creates an `http.Server` using:

* The supplied server address.
* The application HTTP handler returned by `NewHTTPHandler`.

This provides a clear boundary between application composition and executable startup.

## Testing

Session 103 adds composition-level tests covering:

* Creating a system through the composed HTTP handler.
* Retrieving the created system through the same composed handler.
* HTTP server address configuration.
* HTTP server handler configuration.

The tests verify that the application can be used through its external HTTP boundary without manually constructing the underlying registry and service dependencies.

## Design Principles

### Explicit dependency composition

Dependencies are constructed explicitly:

Registry
↓
SystemService
↓
SystemHandler
↓
HTTP Handler

No dependency injection framework is introduced.

### Layer separation

The application package knows how to compose the system.

The API package handles HTTP concerns.

The registry package owns domain and persistence abstractions.

This prevents the individual layers from becoming responsible for application startup or unrelated infrastructure.

### Small core

The composition layer remains deliberately small.

No database, configuration framework, authentication system, middleware framework, or microservice infrastructure is introduced in this session.

Those concerns can be added later when they provide a concrete architectural requirement.

## Result

Session 103 establishes the first application-level composition path for the registry HTTP API.

The registry can now be constructed as a coherent application rather than as a collection of independently tested layers.

Validation

The full Go test suite passes:

go test ./...

This confirms that the new application composition layer does not break the existing API, registry, execution, ingestion, or other project packages.
