# Session 060 – SQL Pagination

## Overview

This session introduces SQL-style pagination to the query execution engine. Pagination allows query results to be divided into smaller, manageable sets, improving performance and enabling efficient navigation through large collections of execution contexts.

The implementation supports both `LIMIT` and `OFFSET` semantics, making the query engine behave similarly to relational database systems.

## Objectives

The objectives of this session were to:

* Implement SQL-style pagination.
* Support configurable `LIMIT` and `OFFSET` values.
* Handle invalid or edge-case pagination parameters safely.
* Provide comprehensive unit test coverage.

## Implementation

A new `QueryPagination` structure was introduced to represent pagination parameters.

```go
type QueryPagination struct {
    Limit  int
    Offset int
}
```

The pagination logic is implemented by the `ApplyPagination` function.

```go
func ApplyPagination(
    results []ExecutionContext,
    pagination QueryPagination,
) []ExecutionContext
```

The function applies pagination after filtering, sorting, projection, aggregation, or other query operations have produced the final result set.

## Pagination Behaviour

The implementation follows these rules:

* `Offset` specifies the number of records to skip.
* `Limit` specifies the maximum number of records returned.
* Negative offsets are normalised to zero.
* A limit less than or equal to zero returns all remaining results.
* If the offset exceeds the available records, an empty result set is returned.
* Empty input collections are returned unchanged.

This behaviour provides predictable and safe pagination without requiring additional validation by callers.

## Edge Case Handling

Several edge cases are handled explicitly:

* Empty result collections.
* Unlimited pagination (`Limit <= 0`).
* Negative offsets.
* Offsets beyond the available records.
* Partial final pages where fewer records remain than the requested limit.

These checks ensure that pagination behaves consistently regardless of the size of the input data.

## Testing

Comprehensive unit tests were added covering:

* Retrieval of the first page.
* Pagination using offsets.
* Offsets beyond the available results.
* Unlimited pagination.
* Negative offset handling.

These tests verify both normal operation and boundary conditions.

## Result

The query engine now supports SQL-style pagination through `LIMIT` and `OFFSET`, allowing clients to retrieve query results in predictable pages while safely handling invalid or out-of-range pagination values.

## Next Steps

The next stage of development will extend the query engine with support for SQL joins, beginning with `INNER JOIN` operations to enable querying across multiple datasets.
