# Session 027 — Observability Upgrade: Metrics & Duration Precision Fix

## Overview

In this session, the execution and observability layer of the financial-registry system was refined to improve accuracy, consistency, and reliability of runtime metrics.

The primary focus was fixing incorrect request duration reporting and improving metric precision.

---

## Problem Identified

The system originally reported:

- `last_request_duration_ms = 0`

This was misleading because:
- The system was executing requests faster than millisecond resolution
- Duration was being truncated via `.Milliseconds()`

This resulted in loss of precision for fast API calls.

---

## Key Changes Made

### 1. Improved Duration Precision

Switched from millisecond precision to microsecond precision in LoggingMiddleware:

duration := time.Since(start)
execution.RecordDuration(uint64(duration.Microseconds()))

---

### 2. Updated Metrics Model

- Replaced millisecond tracking with microsecond tracking
- Updated snapshot field:

LastRequestDurationUs

This now reflects high-resolution latency values.

---

### 3. Fixed Middleware Responsibility Separation

Clear separation of concerns:

ExecutionContextMiddleware:
- Generates RequestID
- Generates TraceID
- Injects execution context into request
- No timing logic

LoggingMiddleware:
- Measures request duration
- Captures HTTP status
- Updates metrics
- Writes log output

---

### 4. Verified End-to-End Observability Flow

Each request now produces:

- Execution context (RequestID + TraceID)
- Event stream (request lifecycle events)
- Metrics counters (total/success/failure)
- Accurate latency measurement (microseconds)

---

## Final Metrics Output Example

{
  "total_requests": 1,
  "successes": 1,
  "failures": 0,
  "last_request_duration_us": 526
}

---

## Outcome

The system now supports:

- High-resolution latency tracking
- Clean middleware separation of concerns
- Reliable observability signals
- Accurate execution timing for fast requests

---

## Key Learning

Millisecond precision is insufficient for modern high-performance APIs.

Microsecond-level measurement is required even in simple Go services to avoid misleading observability data.

---

## Next Direction

Potential next improvements:

- Per-event span timing (mini tracing system)
- Latency percentiles (p50/p95/p99)
- Persistent event storage (Redis or DB-backed)
