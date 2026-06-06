# Session 023 — Request Duration Tracking

## Overview

In this session, request latency tracking was introduced into the execution observability layer of the financial-registry system. This marks the first step toward performance visibility within the request lifecycle.

---

## What was added

### 1. Duration Metric

A new metric was introduced in the execution package:

* `LastRequestDurationMs`

This stores the duration of the most recent request in milliseconds.

---

### 2. Duration Recording Function

A new function was added:

```go
func RecordDuration(durationMs uint64)
```

This function safely stores the latest request duration using atomic operations.

---

### 3. Metrics Snapshot Update

The `MetricsSnapshot` struct was extended to include:

* `last_request_duration_ms`

This allows the duration metric to be exposed via the `/metrics` endpoint.

---

### 4. Middleware Integration

The execution middleware was updated to:

* Capture request start time using `time.Now()`
* Measure total request duration using `time.Since(start)`
* Record duration after request completion

This ensures end-to-end request timing is captured for every processed request.

---

## Observability Impact

The system now tracks:

* Total request count
* Success count
* Failure count
* Last request execution time

This provides the first performance signal in the observability pipeline.

---

## Current State of System

```
Request
  ↓
Middleware (Execution Context + Timing)
  ↓
Processor
  ↓
Metrics (Counters + Duration)
  ↓
/metrics endpoint
```

---

## Notes

* Duration is stored in milliseconds
* Sub-millisecond requests may appear as `0ms`
* This implementation tracks only the latest request duration (not averages or percentiles yet)

---

## Next Direction

Future improvements may include:

* Average latency tracking
* P95 / P99 latency computation
* Per-request execution traces
* Structured execution event logs
