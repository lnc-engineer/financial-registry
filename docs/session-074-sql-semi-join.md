# Session 074 — SQL SEMI JOIN

## Overview

Session 074 adds SQL SEMI JOIN support to the query execution layer.

A SEMI JOIN returns rows from the left input when at least one matching row exists in the right input. Unlike an INNER JOIN, it does not return right-side attributes and does not duplicate left-side rows when multiple right-side records match.

## Implementation

The new `ApplySemiJoin` function was added to:

`internal/execution/query_join.go`

The implementation:

* Iterates through the left input.
* Resolves the configured left join field.
* Searches the right input for a matching value.
* Returns the left context when a match is found.
* Stops searching after the first match.
* Does not merge right-side attributes into the result.
* Does not produce duplicate results when multiple right-side records match.

## SEMI JOIN Semantics

Given:

```text
Left:
A
B
C

Right:
A
A
C
```

A SEMI JOIN returns:

```text
A
C
```

The two matching `A` records on the right produce only one result for the left-side `A`.

This distinguishes SEMI JOIN from INNER JOIN, where every matching pair would be returned.

## Test Coverage

Session 074 adds coverage for:

* Basic SEMI JOIN matching.
* Multiple right-side matches.
* No matching records.
* Empty right input.
* Preservation of left-side attributes.
* Confirmation that right-side attributes are not included.
* Prevention of duplicate left-side results.

## Validation

The implementation was formatted with:

```bash
gofmt -w internal/execution/query_join.go internal/execution/query_join_test.go
```

The complete test suite passes:

```text
go test ./...
```

All packages passed successfully.

## Result

Session 074 establishes SEMI JOIN behavior as a distinct join operation in the query execution layer, providing existence-based matching while preserving the left-side result set.
