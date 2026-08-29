# Session 079 — SQL JOIN Cardinality

## Overview

Session 079 adds test coverage for SQL JOIN cardinality and match multiplicity.

The purpose of this session is to verify that the JOIN execution layer produces the expected number of joined execution contexts when one or more records on either side satisfy the JOIN condition.

## Cardinality

JOIN cardinality depends on the number of matching records on each side.

For matching records:

- One left record × one right record produces one joined result.
- One left record × two right records produces two joined results.
- Two left records × two right records produces four joined results.

This confirms that each matching pair is represented independently rather than being collapsed or duplicated incorrectly.

## Test Coverage

Session 079 adds coverage for:

- One-to-many JOIN cardinality.
- Many-to-many JOIN cardinality.
- Duplicate matching keys.
- Unmatched records alongside matching records.
- Exact result counts for each scenario.

## One-to-Many

A single left-side record matching two right-side records must produce two results.

This verifies that multiple right-side matches are preserved.

## Many-to-Many

Two left-side records matching two right-side records must produce four results.

The result count follows:

`left matches × right matches`

This is the expected multiplicity for a JOIN where every record on the left matches every qualifying record on the right.

## Unmatched Records

The test suite also verifies JOIN behavior when the left input contains a record for which no corresponding right-side record exists.

Matching records continue to produce their expected joined results while the unmatched record follows the existing `ApplyJoin` behavior.

## Validation

The implementation was formatted with `gofmt` and the complete repository test suite was executed.

```text
go test ./...
