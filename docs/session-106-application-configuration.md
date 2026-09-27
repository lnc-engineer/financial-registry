# Session 106 — Application Configuration Boundary

**Date:** 27 September 2026

## Objective

Introduce a small application configuration boundary so server configuration can be represented explicitly rather than being tied directly to the HTTP server constructor.

The goal is to establish a simple foundation for future application configuration without introducing an external configuration framework prematurely.

## Changes

### Application configuration

Added `internal/app/config.go`.

The application now defines:

```go
type Config struct {
	HTTPAddr string
}
