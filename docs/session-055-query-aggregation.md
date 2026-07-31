# Session 055 – Query Aggregation

## Overview

This session introduces aggregation capabilities to the Financial Registry query engine.
Aggregation allows query results to be summarized by counting the occurrence of values
for a selected field.

## Objectives

- Introduce aggregation support
- Reuse the existing field resolution logic
- Keep aggregation independent from filtering and sorting
- Provide unit test coverage

## Implementation

### QueryAggregation

Describe the `QueryAggregation` struct and the `AggregationCount` type.

### Aggregate Function

Explain that the function:

1. Creates an empty result map.
2. Iterates over each `ExecutionContext`.
3. Resolves the requested field using `resolveField`.
4. Increments the count for each value.
5. Returns the aggregated counts.

## Example

Input:

success
success
failure

Output:

success: 2
failure: 1

## Test Coverage

- Aggregate by status
- Aggregate by span name
- Aggregate by custom attribute
- Empty records
- Unknown field

## Outcome

The query engine now supports basic count aggregation while remaining modular and
extensible for future aggregation types such as sum, average, minimum, maximum,
and grouped aggregations.
