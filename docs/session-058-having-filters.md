# Session 058 - HAVING Filters

## Overview

This session introduces HAVING filter support for aggregated query results.

The HAVING stage allows filtering after grouping and aggregation has completed, similar to the SQL HAVING clause.

Example:

```sql
GROUP BY status
HAVING count > 5
