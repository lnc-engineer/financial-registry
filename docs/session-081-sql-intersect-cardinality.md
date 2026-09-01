# Session 081 — SQL INTERSECT Cardinality Coverage

## Objective

Add focused test coverage for SQL INTERSECT cardinality and duplicate-value semantics.

## Scope

This session verifies that `ApplyIntersect` returns distinct matching left contexts based on the configured field value.

Unlike JOIN execution, where every valid matching pair is preserved, INTERSECT returns at most one result for each distinct matching field value.

## Cardinality Semantics

For each value that exists in both the left and right relations:

- The value produces one result.
- Repeated left-side values do not produce additional results.
- Repeated right-side values do not produce additional results.
- Repeated values on both sides still produce one result.
- The first matching left context is preserved in the result.

For example:

```text
Left:   A, A, B
Right:  A, B, B

Result: A, B
