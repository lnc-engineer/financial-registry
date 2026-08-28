# Session 078 — SQL JOIN Duplicate-Match Semantics

## Overview

Session 078 adds integration coverage for SQL JOIN duplicate-match semantics.

The tests verify that JOIN execution preserves every valid combination when multiple records on either side satisfy the JOIN condition.

This establishes the expected result cardinality for one-to-many and many-to-many relationships while also confirming that unmatched left-side records continue to be preserved by the existing JOIN implementation.

## Coverage

The Session 078 test coverage includes:

* One left record matching multiple right records.
* Multiple left records matching one right record.
* Multiple left records matching multiple right records.
* Preservation of unmatched left records.
* Verification that matched records retain both left-side and right-side attributes.

## Duplicate-Match Semantics

JOIN execution evaluates each left context against each right context.

When a matching condition is satisfied, a joined result is produced for that specific pair.

For example, if two left records match two right records, the expected result contains four joined combinations:

```text
L1 → R1
L1 → R2
L2 → R1
L2 → R2
```

This ensures that valid one-to-many and many-to-many relationships are not collapsed or deduplicated during execution.

## Unmatched Left Records

The existing `ApplyJoin` behavior preserves a left-side context when no matching right-side context exists.

Session 078 explicitly verifies this behavior so that duplicate-match coverage does not accidentally obscure the established unmatched-row semantics.

For example:

```text
Left:
L1 → A
L2 → B

Right:
R1 → A
R2 → C
```

The result contains:

```text
L1 + R1
L2
```

The unmatched `L2` record remains present in the result.

## Attribute Preservation

For matched records, the joined execution context preserves the left-side attributes and adds right-side attributes using the existing `right_` prefix.

For example:

```text
id       = L1
code     = A
right_id = R1
```

This allows both sides of a JOIN to remain accessible without overwriting the original left-side attributes.

## Implementation

No production JOIN implementation changes were required for Session 078.

The existing nested-loop JOIN implementation already produces one result for every matching left/right pair. Session 078 adds tests that formally verify this behavior.

The relevant test coverage is contained in:

```text
internal/execution/query_join_test.go
```

## Validation

The following commands were executed:

```bash
gofmt -w internal/execution/query_join_test.go
go test ./...
```

All tests passed successfully.

## Result

Session 078 establishes regression coverage for JOIN result cardinality across one-to-many and many-to-many relationships.

The query execution layer now has explicit tests confirming that duplicate matching relationships are preserved and that unmatched left-side records retain their existing behavior.
