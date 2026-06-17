# Session 031 — Parent/Child Span Relationships

## Summary

Introduced hierarchical tracing by adding `ParentSpanID` to `ExecutionContext` and enabling child span creation from a parent span.

This allows execution traces to form a tree structure instead of flat spans.

---

## Changes

- Added `ParentSpanID` to `ExecutionContext`
- Implemented `NewChildSpan(parent ExecutionContext)` helper
- Updated middleware to:
  - create root span per request
  - generate child span from root
  - propagate child span through context
- Added span logging for ROOT and CHILD spans

---

## Result

Each request now produces a trace structure:

Root Span → Child Span

With shared `TraceID` and linked `ParentSpanID`.

---

## Outcome

This is the foundation for multi-layer distributed tracing and execution tree visualization.
