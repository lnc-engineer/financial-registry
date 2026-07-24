# Session 049 – Operator-Aware Composite Filtering

## Goal

Extend composite trace filtering to support configurable query operators.

## Completed work

- Added QueryOperator support to CompositeFilter
- Integrated MatchValue() into trace filtering
- Added operator support for:
  - Trace ID filtering
  - Status filtering
  - Lifecycle filtering
- Preserved existing equality filtering behaviour
- Added unit tests for operator-based filtering

## Status

Completed.
