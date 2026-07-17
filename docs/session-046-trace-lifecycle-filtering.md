# Session 046 – Trace Lifecycle Filtering

## Goal

Introduce lifecycle-based filtering helpers to improve trace analysis capabilities.

## Completed Work

- Added `FilterByLifecycle` helper.
- Supports searching spans by lifecycle event name.
- Integrated with existing `LifecycleEvent` architecture.
- Added unit tests covering lifecycle filtering scenarios.

## Testing

Verified with:

```bash
go fmt ./...
go test ./...
