# Session 026 – Trace Identifier Propagation

## Overview

In this session, the execution observability system was extended with Trace ID support.

Previous sessions introduced:

* Request lifecycle tracking
* Execution event recording
* Event metadata enrichment
* Request correlation through Request IDs

While Request IDs allowed events belonging to a single HTTP request to be grouped together, there was no concept of a higher-level execution trace that could eventually span multiple services, workers, or processing components.

This session introduced Trace IDs as a foundational distributed tracing primitive.

---

## Motivation

The execution system already produced correlated event streams.

Example:

```text
request_started
ingestion_started
records_received
records_processed
ingestion_completed
request_completed
```

Each event carried:

```text
RequestID
Timestamp
Metadata
```

This was sufficient for tracing activity within a single request.

However, future distributed execution workflows may involve:

```text
API Request
    ↓
Execution Service
    ↓
Worker
    ↓
Validator
    ↓
Persistence Layer
```

A shared Trace ID allows all related operations to be associated with the same end-to-end workflow.

---

## Execution Context Extension

The ExecutionContext structure was extended to include a Trace ID.

Before:

```go
type ExecutionContext struct {
	RequestID string
	StartTime time.Time
	Metadata  map[string]string
}
```

After:

```go
type ExecutionContext struct {
	RequestID string
	TraceID   string
	StartTime time.Time
	Metadata  map[string]string
}
```

The Trace ID now becomes part of the execution context propagated through the request lifecycle.

---

## Trace Generation

ExecutionContextMiddleware was updated to generate a Trace ID when creating a new execution context.

Example:

```go
execCtx := execution.ExecutionContext{
	RequestID: uuid.NewString(),
	TraceID:   uuid.NewString(),
	StartTime: start,
	Metadata:  make(map[string]string),
}
```

For the current single-service architecture:

* Each request receives a unique Request ID
* Each request receives a unique Trace ID

This establishes the tracing infrastructure required for future multi-component workflows.

---

## Event Propagation

ExecutionEvent was extended to include Trace ID support.

Before:

```go
type ExecutionEvent struct {
	Type      string
	RequestID string
	Timestamp time.Time
	Metadata  map[string]string
}
```

After:

```go
type ExecutionEvent struct {
	Type      string            `json:"type"`
	RequestID string            `json:"request_id"`
	TraceID   string            `json:"trace_id"`
	Timestamp time.Time         `json:"timestamp"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}
```

Every execution event now records both identifiers.

---

## Event Creation Updates

The event creation pipeline was updated to propagate Trace IDs automatically.

Example:

```go
return ExecutionEvent{
	Type:      eventType,
	RequestID: ctx.RequestID,
	TraceID:   ctx.TraceID,
	Timestamp: time.Now().UTC(),
	Metadata:  metadataCopy,
}
```

This ensures all events generated within the same execution context share the same Trace ID.

---

## Verification

Execution was verified through the `/events` endpoint.

Observed output:

```json
{
  "type": "records_processed",
  "request_id": "2c8db4c9-6b3f-4796-9c7d-13d0309ab437",
  "trace_id": "40d50353-ccb1-4f0f-a4bc-1afbd23badde",
  "metadata": {
    "records_processed": "2"
  }
}
```

Verification confirmed:

* Request IDs propagated correctly
* Trace IDs propagated correctly
* Trace IDs remained consistent across all events in a request lifecycle
* Event metadata enrichment continued to function correctly

---

## Runtime Logging Improvements

Execution context logging was updated to expose both identifiers.

Example:

```text
[EXECUTION] request=2c8db4c9-6b3f-4796-9c7d-13d0309ab437 trace=40d50353-ccb1-4f0f-a4bc-1afbd23badde
```

This provides immediate visibility into execution correlation during development and debugging.

---

## Architectural Impact

This session introduces the first distributed tracing concept into the platform.

The execution model now supports two layers of correlation:

```text
TraceID
    ↓
RequestID
    ↓
Execution Events
```

Current architecture:

```text
One TraceID per request
```

Future architecture:

```text
One TraceID
    ├─ Multiple RequestIDs
    ├─ Multiple Services
    ├─ Multiple Workers
    └─ Multiple Execution Components
```

This aligns with modern observability systems and distributed execution platforms.

---

## Session Outcome

Completed:

* Added TraceID to ExecutionContext
* Added TraceID to ExecutionEvent
* Implemented Trace ID generation
* Implemented Trace ID propagation
* Verified Trace IDs through the event store API
* Improved execution logging visibility

The platform now supports foundational distributed tracing concepts while maintaining full request-level observability and structured execution events.
