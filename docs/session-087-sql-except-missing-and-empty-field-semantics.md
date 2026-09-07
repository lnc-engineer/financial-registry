# Session 087 — SQL EXCEPT Missing and Empty-Field Semantics

## Overview

Session 087 extends SQL `EXCEPT` test coverage for missing and empty field values.

The existing `ApplyExcept` implementation resolves fields through the execution-context attribute lookup mechanism. Because execution-context attributes are represented as `map[string]string`, a missing field resolves to the empty string.

This session verifies that behavior explicitly and ensures it remains consistent with the existing distinct-result semantics.

## Test Coverage

The following cases were added:

- Deduplication of multiple left-side contexts with a missing field.
- A missing left-side field matching an explicit empty-string value on the right.
- An explicit empty-string left-side value matching a missing field on the right.
- Missing fields on both sides being excluded from the result.
- Deduplication of missing fields alongside normal field values.
- Deduplication of repeated explicit empty-string values.

## Observed Semantics

For the current execution-context representation:

- Missing fields resolve to `""`.
- Explicit empty-string values also resolve to `""`.
- Consequently, missing and empty-string values are treated as equivalent by `ApplyExcept`.
- The right-side value set excludes matching values from the result.
- The left-side `seen` set preserves distinct-result semantics.
- The first left-side context for a distinct value is preserved.
- Existing left-side result ordering remains unchanged.

## Implementation Impact

No production implementation changes were required for this session.

The existing `ApplyExcept` behavior already provides the required semantics for the current `map[string]string` execution-context model. Session 087 therefore adds regression coverage rather than introducing a new nullable-value abstraction.

## Validation

Formatting and the complete test suite were run successfully:

```text
gofmt -w internal/execution/query_except_test.go
go test ./...
