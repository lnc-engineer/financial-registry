# Session 087 — SQL EXCEPT Empty-Result and Order Semantics

## Overview

Session 087 extends the SQL `EXCEPT` test coverage with focused validation of empty-input behavior, duplicate elimination, first-occurrence preservation, and left-side ordering.

The implementation already supported these semantics; this session formalizes them through additional regression tests.

## Scope

The following `EXCEPT` behaviors are covered:

* An empty right-hand input leaves the left-hand results subject to normal duplicate elimination.
* An empty left-hand input produces no results.
* Both inputs being empty produces no results.
* When duplicate values exist on the left, the first matching left context is preserved.
* When the right-hand input contains no matching values, the complete left-side ordering is preserved.
* Right-side ordering does not influence the ordering of surviving left-side results.

## Test Coverage

### Empty Right Side with Duplicate Left Values

A left input containing repeated values is evaluated against an empty right input.

Expected behavior:

* The result contains distinct left-side values.
* The first occurrence of each distinct value is retained.
* Original left-side ordering is preserved.

### Both Inputs Empty

When both operands contain no records, `EXCEPT` returns an empty result.

This provides an explicit base-case regression test for the operator.

### First Left Context Preservation

When duplicate left-side values contain different attributes, the first context associated with the distinct value is retained.

This verifies that deduplication does not replace the first left-side context with a later duplicate.

### Complete Left Ordering

When the right-hand side contains only unrelated values, all left-side records survive in their original order.

For example:

```text
Left:  D, B, A, C
Right: X, Y

Result: D, B, A, C
```

This confirms that `EXCEPT` does not reorder surviving records.

## Validation

The test suite was formatted and executed with:

```bash
gofmt -w internal/execution/query_except_test.go
go test ./...
```

All packages passed successfully.

## Files Changed

* `internal/execution/query_except_test.go`
* `docs/session-087-sql-except-empty-result-and-order-semantics.md`

## Result

Session 087 establishes regression coverage for `EXCEPT` empty-result behavior, duplicate handling, first-occurrence preservation, and left-side ordering.

No production implementation changes were required.

