# Session 063 — SQL LEFT JOIN

## Overview

Implemented LEFT JOIN behavior for the execution query engine.

A LEFT JOIN preserves every record from the left-hand dataset, even when there is no matching record in the right-hand dataset.

## Implementation

Updated:

* `internal/execution/query_join.go`
* `internal/execution/query_join_test.go`

The existing `ApplyJoin` function was extended to track whether each left-side record matched a right-side record.

When a match exists, right-side attributes are copied into the result using the `right_` prefix.

When no match exists, the original left-side record is still appended to the results.

## Example

Given:

```text
LEFT

account-001
account-002

RIGHT

account-001
```

A LEFT JOIN produces:

```text
account-001 + matching right record
account-002 + no right record
```

The second left-side record is preserved even though no matching right-side record exists.

## Tests

Added tests covering:

* matched left and right records
* unmatched left records
* empty right-side datasets
* empty left-side datasets

## Validation

```bash
gofmt -w internal/execution/query_join.go internal/execution/query_join_test.go
go test ./...
```

All tests pass.

## Key Concept

A LEFT JOIN guarantees:

> Every record from the left side appears in the result.

Matching right-side data is included when available, but the absence of a match does not remove the left-side record.
