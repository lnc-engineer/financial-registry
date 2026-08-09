# Session 064 — SQL RIGHT JOIN

## Overview

Implemented SQL RIGHT JOIN support in the execution package.

A RIGHT JOIN preserves every record from the right-hand input, including records that have no matching record on the left.

## Implementation

Added:

* `ApplyRightJoin`
* RIGHT-side preservation semantics
* Left-side attribute projection using the `left_` prefix
* Independent attribute maps for each joined result

The existing `ApplyJoin` LEFT JOIN implementation was also updated to use independent attribute maps for each joined result.

This prevents Go map aliasing when one record matches multiple records on the opposite side.

## RIGHT JOIN Behaviour

Given:

```text
LEFT
account-001

RIGHT
account-001
account-002
```

The RIGHT JOIN preserves both right-side records:

```text
account-001 → matched with LEFT
account-002 → preserved without LEFT attributes
```

## Attribute Handling

LEFT JOIN adds right-side attributes using:

```text
right_<attribute>
```

RIGHT JOIN adds left-side attributes using:

```text
left_<attribute>
```

Each joined result receives a new attribute map so multiple matches do not overwrite one another.

## Tests

Added coverage for:

* Matching RIGHT JOIN records
* Unmatched right records
* Empty left input
* Empty right input
* Multiple left matches
* LEFT JOIN multiple-right-match regression
* Attribute isolation between joined results

## Validation

```text
go fmt ./...
go test ./...
```

All tests pass.

## Files Changed

```text
internal/execution/query_join.go
internal/execution/query_join_test.go
docs/session-064-sql-right-join.md
```

