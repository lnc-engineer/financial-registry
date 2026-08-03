# Session 057 - Query DISTINCT Support

## Overview

This session introduces DISTINCT support for the financial registry query engine.

The DISTINCT operator removes duplicate query results based on selected projected fields, similar to SQL `SELECT DISTINCT`.

Example:

```sql
SELECT DISTINCT status
FROM traces;
