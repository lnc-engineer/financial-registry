# Session 084 — SQL EXCEPT Edge-Case Coverage

## Overview

Session 084 extends the SQL `EXCEPT` test suite with field-resolution and edge-case coverage.

The tests verify that `ApplyExcept` correctly handles missing fields, empty-string values, built-in execution context fields, and comparisons based only on the selected field.

The existing `EXCEPT` implementation remains unchanged.

## Test Coverage

Added coverage for:

* Missing fields on the left side.
* Missing fields on both sides.
* Explicit empty-string field values.
* `trace_id` as a built-in execution context field.
* Comparisons based exclusively on the selected field.

## Missing Field Semantics

`ApplyExcept` delegates field resolution to `resolveField`.

When an attribute is not present, `resolveField` returns an empty string.

For example:

```text
LEFT:
  missing id
  B

RIGHT:
  A
