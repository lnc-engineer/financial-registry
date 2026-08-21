# Session 072 — SQL RIGHT ANTI JOIN

## Overview

Session 072 adds and validates RIGHT ANTI JOIN support in the query execution layer.

A RIGHT ANTI JOIN returns records from the right input that have no matching record in the left input according to the supplied join condition.

This is useful for identifying records that exist in one dataset but have no corresponding record in another dataset.

## Implementation

RIGHT ANTI JOIN is implemented by `ApplyRightAntiJoin` in:

`internal/execution/query_join.go`

The implementation:

1. Iterates over the right input.
2. Resolves the configured right-side join field.
3. Searches the left input for a matching join value.
4. Excludes the right record when a match is found.
5. Preserves the right record when no match exists.

Unlike a standard RIGHT JOIN, unmatched right records are the only records returned.

## Example

Given the following inputs:

```text
Left:
account-001

Right:
account-001
account-002
