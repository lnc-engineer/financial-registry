# Session 038 – Span Attributes Unification

## Overview

This session introduces a unified attribute system for spans, replacing the previous `Metadata` concept across the entire tracing system.

The goal is to standardise how contextual information is attached to spans and improve trace readability, consistency, and long-term maintainability.

---

## What Changed

### 1. Metadata Fully Removed

The `Metadata` field has been completely removed from:

- ExecutionContext
- ExecutionEvent
- Middleware
- Span creation logic
- Trace rendering logic
- All service-layer usage

There is no remaining usage of Metadata in the system.

---

### 2. Attributes Introduced

A new unified field has been introduced:

Attributes map[string]string

This is now the single source of truth for all span-level contextual information.

---

### 3. API Changes

Old API:

WithMetadata(key, value)

New API:

WithAttribute(key, value)

This ensures a consistent naming model and removes ambiguity in span context handling.

---

### 4. ExecutionEvent Updated

Execution events now carry attributes instead of metadata:

Attributes map[string]string

This ensures consistent propagation of contextual data across:
- spans
- lifecycle events
- trace output
- event logging system

---

### 5. Trace Tree Rendering Improvements

Attributes are now:

- Rendered in the trace tree output
- Sorted for deterministic ordering
- Displayed per span alongside lifecycle events

Example output:

TRACE TREE

Validation (span-id)
    • stage = record_processing
    • record_index = 1
    • result = success
    ◦ Span Started (2026-07-03T...)
    ◦ Span Completed (2026-07-03T...)

---

## Architecture Impact

The tracing system now cleanly separates concerns:

### Span Structure
- TraceID, SpanID, ParentSpanID
- Tree hierarchy

### Attributes
- Contextual metadata describing execution state

### Lifecycle Events
- Ordered execution timeline within each span

### Timing
- StartTime, EndTime, Duration

---

## Why This Matters

This change removes duplication and ambiguity in the tracing model and aligns the system closer to production-grade observability systems such as OpenTelemetry.

It improves:
- Debugging clarity
- Trace consistency
- System extensibility
- Event propagation correctness

---

## Result

The system now has a single unified model for span context:

- No Metadata duplication
- Consistent attribute handling across all layers
- Deterministic trace output
- Cleaner architecture for future observability features
