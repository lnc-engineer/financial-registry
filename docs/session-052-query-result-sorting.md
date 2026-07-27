# Session 052 – Query Result Sorting

## Goal

Introduce sorting support for query results to improve trace analysis and query flexibility.

## Work completed

- Added `QuerySort` model
- Introduced sortable fields:
  - Start Time
  - End Time
  - Span Name
- Added ascending and descending sort orders
- Implemented `SortExecutionContexts`
- Added unit tests covering:
  - Span name ascending
  - Span name descending
  - Start time sorting
  - End time sorting

## Result

Query results can now be ordered independently of filtering, providing a clean foundation for future features such as pagination, ranking, and advanced query execution.
