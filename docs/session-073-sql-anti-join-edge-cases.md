# Session 073 — SQL ANTI JOIN Edge Cases

## Overview

Session 073 adds edge-case test coverage for SQL LEFT ANTI JOIN and RIGHT ANTI JOIN behaviour.

Sessions 071 and 072 introduced LEFT ANTI JOIN and RIGHT ANTI JOIN support. This session strengthens that implementation by validating behaviour across unmatched records, multiple matches, empty inputs, and attribute preservation.

## LEFT ANTI JOIN

A LEFT ANTI JOIN returns records from the left input that have no matching record in the right input.

Session 073 verifies that:

- All left records are returned when there are no matches.
- Matched left records are excluded even when multiple right records match.
- Original left attributes are preserved.
- Right-side attributes are not added to anti-join results.

## RIGHT ANTI JOIN

A RIGHT ANTI JOIN returns records from the right input that have no matching record in the left input.

Session 073 verifies that:

- All right records are returned when there are no matches.
- Matched right records are excluded even when multiple left records match.
- Original right attributes are preserved.
- Left-side attributes are not added to anti-join results.

## Edge Cases Covered

The test suite now covers:

- No matching records.
- Multiple matching records.
- Empty left input.
- Empty right input.
- Preservation of original attributes.
- Exclusion of matched records.
- Preservation of unmatched records.

## Testing

The complete Go test suite passes:

```text
go test ./...

?       github.com/lnc-engineer/financial-registry/cmd/hello    [no test files]
ok      github.com/lnc-engineer/financial-registry/cmd/ingestion-processor
ok      github.com/lnc-engineer/financial-registry/internal/execution
?       github.com/lnc-engineer/financial-registry/internal/middleware  [no test files]
