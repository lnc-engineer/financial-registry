# Session 088 — SQL EXCEPT Execution-Context Field Semantics

## Overview

Session 088 extends SQL `EXCEPT` coverage to execution-context fields supported by the shared field resolver.

The existing `ApplyExcept` implementation resolves comparison fields through `resolveField`. This session verifies that `EXCEPT` correctly performs set difference using built-in execution-context fields rather than only values stored in `Attributes`.

## Scope

The following execution-context fields are covered:

* `trace_id`
* `span_name`
* `status`

Each test verifies that matching values from the right-hand input are excluded from the left-hand result.

## Trace ID Semantics

`EXCEPT` can use `trace_id` as its comparison field.

A left context whose trace ID exists in the right-hand input is excluded, while non-matching trace IDs remain in their original left-side order.

## Span Name Semantics

`EXCEPT` can use `span_name` as its comparison field.

Matching span names are removed from the result while the original left execution context is preserved for retained records.

## Status Semantics

`EXCEPT` can use `status` as its comparison field.

Contexts with statuses present in the right-hand input are excluded, while non-matching statuses remain in the result.

## Implementation Relationship

The tests rely on the existing `resolveField` behavior:

* `trace_id` resolves to `ExecutionContext.TraceID`
* `span_name` resolves to `ExecutionContext.SpanName`
* `status` resolves to `ExecutionContext.Status`
* other fields continue to resolve through `ExecutionContext.Attributes`

No implementation changes were required during this session.

## Validation

Formatting and the complete Go test suite were executed successfully:

```text
gofmt -w internal/execution/query_except_test.go
go test ./...
