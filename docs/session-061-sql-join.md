# Session 061 – SQL JOIN

## Overview

Implemented SQL INNER JOIN support in the query execution engine.

This session introduces relational join behaviour, allowing records from two datasets to be matched using configurable join conditions.

## Features

* Added `JoinCondition` to define join relationships.
* Implemented INNER JOIN execution logic.
* Matches records using selected fields from left and right datasets.
* Supports multiple matching records.
* Merges matching right-side attributes into the resulting execution context.
* Reuses existing query field resolution logic.

## Implementation

A join is performed by comparing values from two execution contexts.

Example:

```go
JoinCondition{
    LeftField:  "trace_id",
    RightField: "trace_id",
}
```

This matches records where both contexts contain the same `trace_id`.

SQL equivalent:

```sql
SELECT *
FROM traces
INNER JOIN spans
ON traces.trace_id = spans.trace_id;
```

## Testing

Added unit tests covering:

* Successful joins between matching records.
* No-match scenarios.
* Multiple matching records.

All tests pass successfully:

```bash
go fmt ./...
go test ./...
```

## Learning Outcome

* Understanding SQL INNER JOIN behaviour.
* Translating relational database concepts into Go code.
* Building reusable query execution components.
* Handling one-to-many relationships during joins.
