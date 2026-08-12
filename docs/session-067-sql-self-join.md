# Session 067 — SQL SELF JOIN

## Overview

This session adds SQL SELF JOIN support to the execution layer.

A SELF JOIN joins a dataset against itself using a join condition. It is useful when records within the same dataset need to be compared or related to one another.

## Implementation

Added `ApplySelfJoin` to `internal/execution/query_join.go`.

The implementation reuses the existing `ApplyJoin` behavior by passing the same execution context slice as both the left and right sides of the join.

```go
func ApplySelfJoin(
	contexts []ExecutionContext,
	condition JoinCondition,
) []ExecutionContext {
	return ApplyJoin(contexts, contexts, condition)
}
```

This keeps SELF JOIN behavior consistent with the existing INNER JOIN implementation.

## Join Behavior

The SELF JOIN:

* Uses the same dataset for both sides of the join.
* Applies the existing `JoinCondition`.
* Preserves left-side attributes.
* Prefixes right-side attributes with `right_`.
* Supports multiple matches.
* Returns an empty result when the input dataset is empty.

## Tests

Added tests covering:

* Matching records in a self join.
* Multiple matching records.
* Empty input handling.

## Validation

Executed:

```bash
go fmt ./...
go test ./...
```

All tests passed successfully.

## Result

Session 067 adds SQL SELF JOIN execution support while reusing the existing join implementation rather than duplicating join logic.
