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

## Behaviour

LEFT ANTI JOIN evaluates each record from the left input against the right input using the supplied `JoinCondition`.

A left-side record is included in the result only when no matching right-side record exists.

When a match is found, the left-side record is excluded and the remaining right-side records are not evaluated for that left-side record.

Unlike a regular JOIN, LEFT ANTI JOIN does not merge or add attributes from the right input. The result contains only the original unmatched left-side contexts.

## Example

Given:

### Left Input

| Account ID | Record |
|---|---|
| account-001 | left-001 |
| account-002 | left-002 |

### Right Input

| Account ID | Record |
|---|---|
| account-001 | right-001 |

Using:

```text
LEFT ANTI JOIN
ON left.account_id = right.account_id
