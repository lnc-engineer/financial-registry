# Session 082 — SQL INTERSECT Order Preservation

## Objective

Add focused test coverage verifying that SQL `INTERSECT` preserves the ordering of matching values from the left relation while maintaining distinct-result semantics.

## Scope

This session verifies that `ApplyIntersect`:

- Preserves the order established by the left input.
- Ignores the ordering of the right input when producing results.
- Emits each matching field value at most once.
- Preserves the first matching left context for duplicate values.
- Maintains result ordering when unmatched left rows are interleaved with matching rows.

The implementation already establishes these semantics by iterating through the left relation and appending the first occurrence of each value that also exists in the right relation.

## Ordering Semantics

`INTERSECT` is evaluated using the left relation as the source of result ordering.

For example:

```text
Left:   C, A, B
Right:  B, C, A

Result: C, A, B
