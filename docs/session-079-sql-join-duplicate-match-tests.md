# Session 079 — SQL JOIN Duplicate Match Tests

## Objective

Document and verify the expected result semantics when SQL JOIN predicates match multiple rows.

## Scope

This session focuses on duplicate-match behavior for JOIN execution.

The key requirement is that a matching pair of rows produces one joined result for each valid match. JOIN execution must not create additional duplicate results beyond those implied by the input matches.

## Duplicate Match Semantics

For a JOIN between a left relation and a right relation:

* Each left row may match zero, one, or multiple right rows.
* Each matching left/right pair produces one joined result.
* Multiple matching right rows therefore produce multiple joined results.
* Non-matching rows must not generate additional results for an INNER JOIN.
* The implementation must not duplicate a valid match accidentally during predicate evaluation.

For example:

```sql
SELECT *
FROM transactions t
JOIN accounts a
  ON t.account_id = a.account_id;
```

If one transaction matches two account rows, the execution result contains two joined rows because there are two valid matching pairs.

## Test Coverage

The JOIN test suite covers:

* A single matching pair producing one result.
* Multiple right-side matches producing multiple results.
* Multiple left-side matches producing their corresponding joined results.
* Non-matching rows producing no additional INNER JOIN results.
* Multi-condition JOIN predicates preserving correct match cardinality.
* FULL OUTER JOIN behavior when duplicate matches are present.
* Prevention of unintended duplicate results.

## Expected Behavior

JOIN result cardinality is determined by the number of valid matching row pairs.

For a left row with `N` matching right rows:

```text
1 left row × N matching right rows = N joined results
```

The execution layer should therefore preserve legitimate multiplicity while avoiding artificial duplication.

## Validation

The implementation was formatted with:

```bash
gofmt -w internal/execution/query_join.go internal/execution/query_join_test.go
```

The complete test suite was validated with:

```bash
go test ./...
```

All tests must pass before the session is committed.

## Outcome

Session 079 establishes explicit test coverage for SQL JOIN duplicate-match semantics and protects JOIN result cardinality against accidental duplication during future changes.
