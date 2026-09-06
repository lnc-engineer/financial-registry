# Session 086 — SQL EXCEPT Missing-Field Semantics

## Overview

Session 086 extends SQL `EXCEPT` test coverage to explicitly define behavior when the field used for comparison is missing from one or both input contexts.

The implementation continues to use `resolveField()` for field lookup. When an attribute is not present, `resolveField()` returns an empty string. Session 086 adds regression coverage around this behavior without modifying the existing `EXCEPT` implementation.

## Scope

This session covers:

* Missing comparison fields on the right-hand input.
* Deduplication of repeated missing-field values.
* Preservation of valid left-hand values when right-hand values are present.
* Continued preservation of the original left-side context.
* Regression protection for existing `EXCEPT` semantics.

## Implementation Behavior

`ApplyExcept()` builds a set of comparison values from the right-hand contexts and then evaluates each left-hand context against that set.

Field resolution is delegated to `resolveField()`.

For attributes that do not exist, `resolveField()` returns:

```text
""
```

Consequently, a missing comparison field participates in `EXCEPT` comparison as an empty-string value.

This behavior is now explicitly covered by tests.

## Test Coverage

### Missing Field on Right

Verifies that a right-hand context containing no value for the comparison field contributes the resolved empty-string value to the exclusion set.

A left-hand context with a valid field value remains eligible for the result.

### Deduplication of Missing Field Values

Verifies that multiple left-hand contexts with the comparison field missing are treated as the same resolved value and therefore do not produce duplicate `EXCEPT` results.

The first qualifying left-hand context is preserved.

### Missing Field Does Not Affect Valid Right Values

Verifies that a missing comparison field does not incorrectly exclude unrelated valid values.

For example, when the right side contains `id = "B"`, a left-side `id = "C"` remains in the result even when another left-side context has no `id`.

## Result Semantics

The Session 086 tests confirm that `EXCEPT` continues to:

* Return only qualifying left-hand contexts.
* Exclude values represented on the right-hand side.
* Deduplicate results by the resolved comparison value.
* Preserve the original left-hand context.
* Preserve left-side result ordering.
* Treat missing fields consistently through `resolveField()`.

## Validation

Formatting and the complete repository test suite were run:

```bash
gofmt -w internal/execution/query_except_test.go
go test ./...
```

All packages passed successfully.

## Files

### Modified

```text
internal/execution/query_except_test.go
```

### Added

```text
docs/session-086-sql-except-missing-field-semantics.md
```

### Implementation

No changes were made to:

```text
internal/execution/query_except.go
```

## Conclusion

Session 086 establishes explicit regression coverage for missing-field behavior in SQL `EXCEPT`.

The existing implementation remains unchanged because the observed behavior is consistent and testable. The new tests document how unresolved comparison fields participate in exclusion and deduplication semantics.
