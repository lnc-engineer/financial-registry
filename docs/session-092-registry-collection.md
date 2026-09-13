# Session 092 — Registry Collection

## Goal

Introduce the first registry collection abstraction responsible for storing and retrieving registered financial systems.

## Changes

Added `internal/registry/registry.go` with:

- `Registry` type backed by a map of system IDs to `System` values.
- `NewRegistry()` for creating an empty registry.
- `Register()` for adding valid systems.
- Duplicate system IDs are rejected.
- Invalid systems are rejected through `System.Validate()`.
- `Get()` for retrieving a registered system by ID.

Added `internal/registry/registry_test.go` covering:

- Empty registry creation.
- Successful system registration.
- Duplicate ID rejection.
- Invalid system rejection.
- Successful system lookup.
- Unknown system ID lookup.

## Domain Responsibility

The registry now separates system-level behavior from collection-level behavior:

- `System` owns its domain state, validation, and lifecycle transitions.
- `Registry` owns system registration, uniqueness, storage, and lookup.

This establishes the first collection boundary for the Financial Systems Registry domain.

## Verification

The following commands pass:

```bash
go test ./internal/registry
go test ./...
