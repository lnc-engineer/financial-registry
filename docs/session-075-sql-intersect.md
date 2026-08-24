# Session 075 — SQL INTERSECT

## Overview

Session 075 adds SQL `INTERSECT` support to the query execution layer.

`INTERSECT` returns the distinct records that are present in both input result sets according to a specified field.

Unlike a regular JOIN, `INTERSECT` does not combine attributes from both inputs. The resulting records preserve the `ExecutionContext` from the left input.

## Implementation

The implementation is provided by:

`internal/execution/query_intersect.go`

The main operation is:

```go
ApplyIntersect(
    left []ExecutionContext,
    right []ExecutionContext,
    field string,
) []ExecutionContext
```

The implementation:

1. Builds a lookup set of field values from the right input.
2. Iterates through the left input.
3. Resolves the configured field using `resolveField`.
4. Keeps the left context when its value exists in the right-side lookup.
5. Uses a `seen` set to ensure distinct results.
6. Returns only the matching left-side contexts.

## Example

Given the following inputs:

```text
Left:
A
B
C

Right:
B
C
D
```

The intersection is:

```text
B
C
```

`A` is excluded because it does not exist in the right input, while `D` is excluded because it does not exist in the left input.

## Duplicate Handling

SQL `INTERSECT` returns distinct results by default.

For example:

```text
Left:
A
A
B

Right:
A
B
B
```

The result is:

```text
A
B
```

The implementation achieves this using a `seen` lookup map.

## Attribute Preservation

Only the matching `ExecutionContext` from the left input is returned.

For example:

```text
Left:
id=A
name=Alice

Right:
id=A
other=value
```

The resulting context contains:

```text
id=A
name=Alice
```

The right-side `other` attribute is not added to the result.

This keeps `INTERSECT` distinct from JOIN operations, where attributes from both inputs may be combined.

## Edge Cases

The implementation handles:

* Empty left input.
* Empty right input.
* No matching records.
* Duplicate values.
* Matching records with additional attributes.
* Multiple matching records.

## Tests

Session 075 adds `internal/execution/query_intersect_test.go`.

Test coverage includes:

* Common records are returned.
* Left-only records are excluded.
* Right-only records are excluded.
* Duplicate results are removed.
* Empty left input returns no results.
* Empty right input returns no results.
* No matches return no results.
* Left-side context and attributes are preserved.
* Right-only attributes are not included.

## Validation

Formatting was applied with:

```bash
gofmt -w internal/execution/query_intersect.go internal/execution/query_intersect_test.go
```

The execution package tests passed:

```text
ok github.com/lnc-engineer/financial-registry/internal/execution
```

The complete project test suite also passed:

```bash
go test ./...
```

Result:

```text
?    github.com/lnc-engineer/financial-registry/cmd/hello    [no test files]
ok   github.com/lnc-engineer/financial-registry/cmd/ingestion-processor
ok   github.com/lnc-engineer/financial-registry/internal/execution
?    github.com/lnc-engineer/financial-registry/internal/middleware [no test files]
```

## Outcome

Session 075 successfully adds SQL `INTERSECT` support to the query execution layer with distinct-result semantics, left-context preservation, duplicate elimination, edge-case coverage, and full project test validation.
