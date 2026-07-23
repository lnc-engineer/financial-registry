# Session 048 – Trace Query Operators

## Goal

Extend the trace query system with reusable query operators for flexible filtering.

## Implemented

- Added `QueryOperator` type.
- Introduced supported operators:
  - `equals`
  - `contains`
  - `starts_with`
- Implemented `MatchValue()` helper for evaluating query conditions.
- Added unit tests covering each operator and non-matching scenarios.

## Result

The trace query system now supports reusable comparison operators that can be integrated into future query builders and composite filters while keeping the API extensible.
