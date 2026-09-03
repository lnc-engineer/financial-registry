# Session 083 — SQL EXCEPT Support

## Overview

Session 083 adds SQL `EXCEPT`-style set difference support to the execution layer.

The new `ApplyExcept` operation compares a left and right collection of `ExecutionContext` values using a selected field. Records whose field value exists on the right side are excluded, while distinct left-side values that do not appear on the right are preserved.

The implementation follows the same set-operation semantics established by `ApplyIntersect`.

## Implementation

Added:

```text
internal/execution/query_except.go
```

The implementation provides:

```go
ApplyExcept(
    left []ExecutionContext,
    right []ExecutionContext,
    field string,
) []ExecutionContext
```

The operation:

1. Builds a lookup set containing field values from the right-hand contexts.
2. Iterates through the left-hand contexts.
3. Excludes values present in the right-hand set.
4. Deduplicates remaining left-side values.
5. Preserves the first left context associated with each distinct value.
6. Preserves the ordering of the left input.

## Test Coverage

Added:

```text
internal/execution/query_except_test.go
```

Coverage includes:

* Returning left-only records.
* Excluding matching records.
* Excluding right-only records from the output.
* Removing duplicate left values.
* Empty left input.
* Empty right input.
* No matching values.
* Preserving the complete left-side context.
* Preserving the first left context for duplicate values.
* Preserving left-side ordering.
* Ensuring right-side ordering does not affect the result.
* Handling repeated values on both sides.

## Semantics

For inputs conceptually represented as:

```text
LEFT:  A, B, C
RIGHT: B, C, D
```

`EXCEPT` produces:

```text
A
```

Only values originating from the left input can appear in the result.

Duplicate values are treated as a single set member. Therefore:

```text
LEFT:  A, A, B
RIGHT: C
```

produces:

```text
A, B
```

The first left-side context for a retained value is preserved.

## Ordering

The result follows the ordering of the left input rather than the ordering of the right input.

For example:

```text
LEFT:  C, A, B
RIGHT: A
```

produces:

```text
C, B
```

This provides deterministic left-side ordering consistent with the existing `INTERSECT` implementation.

## Validation

Formatting and the complete Go test suite were executed successfully:

```bash
gofmt -w internal/execution/query_except.go internal/execution/query_except_test.go
go test ./...
```

Result:

```text
ok   github.com/lnc-engineer/financial-registry/cmd/ingestion-processor
ok   github.com/lnc-engineer/financial-registry/internal/execution
?    github.com/lnc-engineer/financial-registry/internal/middleware [no test files]
```

## Outcome

Session 083 establishes `EXCEPT` set-difference semantics in the execution layer with distinct-value handling, left-context preservation, and deterministic left-side ordering.
