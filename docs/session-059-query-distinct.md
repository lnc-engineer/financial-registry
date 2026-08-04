# Session 059 - Query DISTINCT

## Overview

This session introduces DISTINCT query support for removing duplicate query results.

DISTINCT behaves similarly to the SQL DISTINCT clause by ensuring that only unique projected values are returned.

Example:

```sql
SELECT DISTINCT status
FROM traces;
```

Duplicate status values are removed from the result set.

## Implementation

The DISTINCT query model is represented by:

```go
type DistinctQuery struct {
	Enabled bool
}
```

The DISTINCT operation is implemented through:

```go
func ApplyDistinct(
	records []ExecutionContext,
	projection []string,
) []ExecutionContext
```

The function:

- evaluates projected fields
- creates a unique key for each record
- removes duplicate results
- preserves the order of the first occurrence

## Multi-field DISTINCT

Multiple projection fields are supported.

Example:

```sql
SELECT DISTINCT status, span_name
FROM traces;
```

A composite key is generated from all projected fields to identify unique rows.

## Query Execution Order

DISTINCT is applied after projection because uniqueness is determined from the selected fields.

Query execution order:

```
FROM
WHERE
GROUP BY
HAVING
SELECT (projection)
DISTINCT
ORDER BY
LIMIT/OFFSET
```

## Testing

The DISTINCT implementation includes tests covering:

- duplicate removal
- multi-field projections
- preservation of result order
- empty projection behaviour

Validation:

```bash
go test ./...
```

All tests pass.

## Session Completion

Session 059 complete.
