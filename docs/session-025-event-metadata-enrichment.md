# Session 025 – Event Metadata Enrichment

## Overview

In this session, the execution event system was extended to support structured metadata enrichment.

Previously, execution events only captured:

* Event type
* Request ID
* Timestamp

While this provided basic request tracing, events could not carry contextual information about what occurred during execution.

This session introduced event metadata, allowing execution events to include structured key-value information that can be used for observability, auditing, diagnostics, and future replay capabilities.

---

## Motivation

The event pipeline already provided end-to-end request correlation through Request IDs.

Example:

```text
request_started
ingestion_started
records_received
records_processed
ingestion_completed
request_completed
```

However, events only described that something happened and did not include additional execution details.

For example:

```json
{
  "type": "records_processed"
}
```

does not indicate how many records were processed.

Metadata enrichment allows events to carry structured execution information.

Example:

```json
{
  "type": "records_processed",
  "metadata": {
    "records_processed": "2"
  }
}
```

---

## Execution Event Extension

The ExecutionEvent type was extended with a Metadata field.

Before:

```go
type ExecutionEvent struct {
	Type      string
	RequestID string
	Timestamp time.Time
}
```

After:

```go
type ExecutionEvent struct {
	Type      string            `json:"type"`
	RequestID string            `json:"request_id"`
	Timestamp time.Time         `json:"timestamp"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}
```

The `omitempty` tag ensures metadata is only included when present.

---

## Metadata Propagation

The execution context already contained a metadata map:

```go
type ExecutionContext struct {
	RequestID string
	StartTime time.Time
	Metadata  map[string]string
}
```

Events now capture a snapshot of the execution context metadata when they are created.

---

## Metadata Isolation

An important issue was discovered during implementation.

Because Go maps are reference types, directly assigning:

```go
Metadata: ctx.Metadata
```

caused all events to share the same metadata map.

As metadata was updated later in the request lifecycle, previously recorded events appeared to change retroactively.

To prevent this, event creation now performs a metadata copy:

```go
metadataCopy := make(map[string]string)

for k, v := range ctx.Metadata {
	metadataCopy[k] = v
}
```

This ensures each event stores an immutable snapshot of metadata at the time it was recorded.

---

## First Metadata-Enriched Event

The ingestion service now records the number of successfully processed records.

Example:

```go
ctx.Metadata["records_processed"] =
	strconv.Itoa(len(validRecords))

execution.LogEvent(ctx, "records_processed")
```

Resulting event:

```json
{
  "type": "records_processed",
  "request_id": "eb6736c1-b482-4044-8016-7f6b164cfeac",
  "metadata": {
    "records_processed": "2"
  }
}
```

---

## Verification

Execution was verified through the `/events` endpoint.

Observed behaviour:

* Request IDs remained consistent across the entire request lifecycle.
* Metadata appeared only on events recorded after enrichment.
* Earlier events remained unchanged.
* Event snapshots were isolated correctly.

Example sequence:

```text
request_started
ingestion_started
records_received
records_processed
ingestion_completed
request_completed
```

with metadata attached to enriched events.

---

## Architectural Impact

This session transitions the event system from simple event logging toward structured observability.

The execution pipeline can now capture both:

* What happened
* Context about what happened

This provides a foundation for:

* Distributed tracing
* Audit trails
* Event replay systems
* Persistent execution storage
* Intelligent workflow observability

---

## Session Outcome

Completed:

* Added Metadata support to ExecutionEvent
* Introduced structured event enrichment
* Implemented metadata snapshot isolation
* Recorded first business-level execution metadata
* Verified enriched events through the event store API

The execution event system now supports structured observability data while maintaining request-level traceability through Request IDs.
