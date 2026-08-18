# Session 071 — SQL LEFT ANTI JOIN

## Overview

Session 071 adds LEFT ANTI JOIN support to the query execution layer.

A LEFT ANTI JOIN returns records from the left input that have no matching record in the right input according to the supplied join condition.

This is useful for identifying records that exist in one dataset but have no corresponding record in another dataset.

## Implementation

The implementation was added to:

- `internal/execution/query_join.go`

The new function is:

```go
ApplyLeftAntiJoin(
    left []ExecutionContext,
    right []ExecutionContext,
    condition JoinCondition,
) []ExecutionContext
