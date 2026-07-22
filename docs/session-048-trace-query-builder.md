# Session 048 – Trace Query Builder

## Goal

Introduce a fluent builder for constructing trace queries.

## Implemented work

- Added a `TraceQueryBuilder` for fluent query construction
- Added support for chaining `WithTraceID()` and `WithStatus()`
- Added `Build()` to produce a `CompositeFilter`
- Added unit tests covering builder behavior
- Kept the API simple and extensible for future filter types

## Status

Implemented.
