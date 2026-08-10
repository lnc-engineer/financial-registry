# Session 065 — SQL FULL OUTER JOIN

## Overview

Session 065 adds support for SQL FULL OUTER JOIN semantics to the query execution layer.

A FULL OUTER JOIN preserves:

- Matching records from both sides
- Unmatched records from the left side
- Unmatched records from the right side

This extends the existing JOIN functionality implemented in previous sessions.

## Implementation

Added:

func ApplyFullOuterJoin(
    left []ExecutionContext,
    right []ExecutionContext,
    condition JoinCondition,
) []ExecutionContext

The implementation:

1. Iterates through every left-side record.
2. Compares the configured join fields against every right-side record.
3. Produces joined results for matching records.
4. Preserves unmatched left-side records.
5. Tracks matched right-side records.
6. Appends unmatched right-side records after processing the left side.

Right-side attributes are prefixed with `right_` when records are joined.

## Example

Given:

LEFT

account-001
account-002

RIGHT

account-001
account-003

The FULL OUTER JOIN produces:

account-001 + account-001
account-002
account-003

This ensures that records existing on either side are preserved.

## Tests

Added coverage for:

- Matching records
- Unmatched left records
- Unmatched right records
- Both sides containing unmatched records
- Empty left input
- Empty right input
- Multiple matching records

## Validation

The implementation was validated with:

go fmt ./...
go test ./...

All tests passed successfully.

## Result

The query execution layer now supports FULL OUTER JOIN semantics alongside the existing JOIN, LEFT JOIN, and RIGHT JOIN implementations.
