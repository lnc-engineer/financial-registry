# Session 039 – Trace Status Export

## Overview

This session enhances the trace export system by adding execution status information to every exported span.

Previously, traces contained structural execution information including trace IDs, span relationships, duration tracking, and attributes. However, there was no explicit indication of whether an individual span completed successfully or failed.

This session introduces span status as a first-class observability signal and improves trace state consistency by ensuring exported traces use the final recorded execution state.

---

## What Was Implemented

### Added Status Tracking to Execution Context

The execution context already contained a status field. This session added helper methods for controlled status updates:

- `MarkSuccess()`
- `MarkFailure()`

This provides a consistent way for execution components to update completion state.

---

## Updated Trace Export

The trace export representation now includes execution status alongside existing trace information.

Each exported span now reports:

- Trace ID
- Span ID
- Parent Span ID
- Span name
- Duration
- Child spans
- Execution status

Supported execution states:

- `SUCCESS`
- `FAILURE`

This removes the need to infer execution outcome from other trace fields.

---

## Status Propagation

Execution paths now explicitly mark successful completion before finishing spans.

The successful execution flow is:


Process Execution
|
v
MarkSuccess()
|
v
FinishSpan()
|
v
RecordSpan()


This ensures that the stored execution span contains the final execution outcome before trace export.

---

## Trace State Consistency Improvement

During validation, a stale trace state issue was identified.

Previously, middleware maintained its own span buffer:


middleware spanBuffer
|
v
BuildTraceTree()
|
v
Trace Export


At the same time, completed spans were stored separately through:


FinishSpan()
|
v
RecordSpan()
|
v
Execution Store


This created two sources of trace state, causing exported traces to contain outdated information such as:


status: unknown


The middleware was updated to use the execution span store as the single source of truth.

The new flow is:


FinishSpan()
|
v
RecordSpan()
|
v
GetSpans()
|
v
BuildTraceTree()
|
v
Trace Export


This guarantees that exported traces represent the final completed execution state.

---

## Added Trace Summary

A trace summary component was introduced to provide aggregated execution information.

The summary reports:

- Total spans
- Successful spans
- Failed spans
- Maximum trace depth
- Total execution duration

Example:


TotalSpans: 2
SuccessCount: 2
FailureCount: 0


This provides a foundation for future monitoring and analytics features.

---

## Status Normalisation

Execution status values were standardised across the tracing system.

Before:


success


After:


SUCCESS


This ensures consistent aggregation and allows summary calculations to correctly identify successful spans.

---

## Files Modified


cmd/ingestion-processor/handler.go
cmd/ingestion-processor/processor.go
cmd/ingestion-processor/service.go

internal/execution/context.go
internal/execution/trace_export.go
internal/execution/trace_summary.go

internal/middleware/execution.go


---

## Result

The Financial Registry now exports traces containing explicit execution outcomes.

Each completed span reports whether it succeeded or failed, making execution history easier to inspect and analyse.

This improvement enables future observability capabilities including:

- failed span detection
- success-rate calculations
- trace filtering
- monitoring dashboards
- execution analytics

The tracing architecture now maintains a single source of truth for completed spans, improving reliability and preparing the system for more advanced distributed observability features.
