# Session 066 — SQL CROSS JOIN

## Overview

Session 066 adds support for SQL CROSS JOIN semantics to the query execution layer.

A CROSS JOIN produces the Cartesian product of two result sets.

Every record from the left side is combined with every record from the right side.

For example:

- 2 left records
- 3 right records

produce:

- 2 × 3 = 6 joined results

This extends the existing JOIN functionality implemented in previous sessions.

## Implementation

Added:

func ApplyCrossJoin(
    left []ExecutionContext,
    right []ExecutionContext,
) []ExecutionContext

The implementation:

1. Iterates through every left-side record.
2. Iterates through every right-side record for each left record.
3. Creates a joined result for every left/right combination.
4. Preserves left-side attributes.
5. Copies right-side attributes using the `right_` prefix.
6. Returns an empty result when either input is empty.

Unlike the existing JOIN implementations, CROSS JOIN does not use a `JoinCondition`.

Every left record is combined with every right record.

## Example

Given:

LEFT

- account-001
- account-002

RIGHT

- currency-GBP
- currency-USD

The CROSS JOIN produces:

- account-001 + currency-GBP
- account-001 + currency-USD
- account-002 + currency-GBP
- account-002 + currency-USD

The result count is:

2 × 2 = 4

This demonstrates the Cartesian product behavior of CROSS JOIN.

## Attribute Handling

Left-side attributes are preserved using their original names.

Right-side attributes are prefixed with `right_`.

For example:

left:

`account_id = account-001`

right:

`currency = GBP`

produces:

`account_id = account-001`

`right_currency = GBP`

This follows the attribute handling convention already used by the existing JOIN implementations.

## Tests

Added coverage for:

- Cartesian product generation
- Left attribute preservation
- Right attribute prefixing
- Multiple left and right combinations
- Empty left input
- Empty right input

The tests verify both result counts and attribute preservation.

## Validation

The implementation was validated with:

```bash
go fmt ./internal/execution/query_join.go
go fmt ./internal/execution/query_join_test.go
go test ./...
