# Session 062 – SQL UNION Query Support

## Overview

This session introduces SQL UNION-style query support into the execution engine.

UNION allows multiple query results to be combined into a single result set while removing duplicate records.

Example:

```sql
SELECT *
FROM traces
WHERE status = 'success'

UNION

SELECT *
FROM traces
WHERE trace_id = 'trace-002';
