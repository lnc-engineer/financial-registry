# Session 030 — Structured Tracing with SpanID Support

## Overview

This session extended the existing observability system by introducing **SpanID propagation** across the execution lifecycle. The system now supports three core identifiers:

* RequestID → identifies a single HTTP request
* TraceID → links the full lifecycle of a request across the system
* SpanID → identifies individual execution scopes within a request

---

## Changes Introduced

### 1. ExecutionContext Update

Added `SpanID` to the execution context:

* Ensures every request carries a span-level identifier
* Allows downstream logging and event emission to reference execution scope

---

### 2. ExecutionEvent Update

All events now include:

* RequestID
* TraceID
* SpanID
* Timestamp
* Metadata (optional)

This enables full reconstruction of request execution flows.

---

### 3. Middleware Enhancement

The `ExecutionContextMiddleware` now initializes:

* RequestID (UUID)
* TraceID (UUID)
* SpanID (UUID)

These values are injected into request context for downstream usage.

---

### 4. Logging Pipeline Update

All emitted execution events now propagate SpanID from context, ensuring trace consistency across:

* request_started
* ingestion_started
* records_received
* records_processed
* ingestion_completed
* request_completed

---

## Observability Model (Current State)

```
Request
 ├── Trace (TraceID)
 │     ├── Span: request_started
 │     ├── Span: ingestion_started
 │     ├── Span: processing
 │     └── Span: completion
```

---

## Key Insight

At this stage, SpanID is still globally assigned per request. Future improvements will evolve SpanID into **phase-level execution identifiers** to represent real operational boundaries.

---

## Next Steps

Planned evolution:

1. Replace single SpanID per request with **phase-based spans**
2. Introduce semantic spans:

   * validation
   * processing
   * persistence
   * response
3. Begin trace reconstruction endpoint design

---
