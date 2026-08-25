# Session 075 — SQL Multi-Condition JOIN

## Overview

Session 075 extends SQL JOIN support with multi-condition JOIN predicates.

Previously, JOIN operations supported a single pair of fields through `JoinCondition`. This session introduces `JoinPredicate` and allows a `JoinCondition` to contain multiple field comparisons.

A multi-condition JOIN requires **all configured predicates to match** before two execution contexts are considered joined.

This provides support for SQL-style joins such as:

```sql
SELECT *
FROM transactions
JOIN payments
  ON transactions.account_id = payments.account_id
 AND transactions.currency = payments.currency;
```

The implementation is shared across the supported conditional JOIN operations.

## Multi-Condition Join Model

The JOIN condition now supports individual predicates:

```go
type JoinPredicate struct {
    LeftField  string
    RightField string
}
```

A `JoinCondition` can contain multiple predicates:

```go
type JoinCondition struct {
    LeftField  string
    RightField string
    Conditions []JoinPredicate
}
```

The existing `LeftField` and `RightField` fields are retained for backward compatibility with single-condition JOINs.

When `Conditions` is populated, all predicates are evaluated.

## AND Semantics

Multi-condition JOINs use logical AND semantics.

For example:

```go
JoinCondition{
    Conditions: []JoinPredicate{
        {
            LeftField:  "account_id",
            RightField: "account_id",
        },
        {
            LeftField:  "currency",
            RightField: "currency",
        },
    },
}
```

A left and right context only match when both comparisons succeed:

```text
left.account_id == right.account_id
AND
left.currency == right.currency
```

If any predicate fails, the contexts are not considered a match.

## Shared Predicate Evaluation

The implementation introduces a common predicate evaluation path through:

```go
func joinContextsMatch(
    leftCtx ExecutionContext,
    rightCtx ExecutionContext,
    condition JoinCondition,
) bool
```

This ensures that all conditional JOIN operations use the same matching semantics.

The helper evaluates every configured `JoinPredicate` and returns `false` as soon as a predicate does not match.

This avoids duplicating multi-condition matching logic across individual JOIN implementations.

## Supported JOIN Operations

Multi-condition matching is applied to:

* INNER JOIN
* RIGHT JOIN
* FULL OUTER JOIN
* SELF JOIN
* LEFT ANTI JOIN
* RIGHT ANTI JOIN
* SEMI JOIN

For example, `ApplySemiJoin` now determines whether a left context has at least one right context satisfying **all** JOIN predicates.

## Backward Compatibility

Single-condition JOINs continue to work using the existing structure:

```go
JoinCondition{
    LeftField:  "account_id",
    RightField: "account_id",
}
```

The implementation converts this legacy representation into a single `JoinPredicate` internally.

This means existing JOIN callers do not need to change.

## Joined Attribute Handling

The existing attribute behaviour is preserved.

For standard JOIN operations, attributes from the left context remain unprefixed while attributes from the right context receive the `right_` prefix.

For example:

```text
account_id=account-001
currency=GBP
right_status=completed
```

Right JOIN operations retain their existing behaviour of preserving right-side attributes and prefixing left-side attributes with `left_`.

Anti JOIN and SEMI JOIN operations continue to return only the original side's context attributes.

## Example

Consider the following left contexts:

```text
account_id=account-001
currency=GBP

account_id=account-001
currency=USD
```

and right contexts:

```text
account_id=account-001
currency=GBP

account_id=account-001
currency=EUR
```

Using:

```go
JoinCondition{
    Conditions: []JoinPredicate{
        {
            LeftField:  "account_id",
            RightField: "account_id",
        },
        {
            LeftField:  "currency",
            RightField: "currency",
        },
    },
}
```

only the GBP records match.

The account ID matches in both cases, but the currency predicate eliminates the USD/EUR combination.

## Empty and Multiple Predicate Behaviour

An empty input continues to follow the existing JOIN semantics.

Multiple matching records continue to produce multiple results.

For example, if two right-side contexts satisfy every predicate for a single left context, the JOIN produces two joined results.

This preserves normal one-to-many JOIN behaviour.

## Test Coverage

Session 075 adds coverage for multi-condition JOIN behaviour, including:

* Matching multiple predicates
* Rejecting records when one predicate fails
* Matching when all predicates succeed
* Multiple matching right-side records
* Multi-condition JOINs with different fields
* Backward compatibility with single-condition JOINs
* Multi-condition SEMI JOIN behaviour
* Multi-condition ANTI JOIN behaviour
* Multi-condition RIGHT JOIN behaviour
* Multi-condition FULL OUTER JOIN behaviour

The tests verify that predicates are evaluated collectively rather than independently.

## Implementation Summary

Session 075 introduces reusable multi-condition JOIN matching without changing the established behaviour of existing JOIN types.

The main additions are:

```go
type JoinPredicate struct {
    LeftField  string
    RightField string
}
```

and:

```go
type JoinCondition struct {
    LeftField  string
    RightField string
    Conditions []JoinPredicate
}
```

A shared predicate matcher provides consistent **AND-based JOIN semantics** across the execution layer.

This establishes a more expressive JOIN foundation for future query features while maintaining compatibility with the existing single-condition JOIN API.

## Verification

The implementation was formatted and verified with:

```bash
gofmt -w internal/execution/query_join.go internal/execution/query_join_test.go
go test ./...
```

All packages and tests pass successfully.
