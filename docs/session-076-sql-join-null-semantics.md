# Session 076 — SQL JOIN NULL Semantics

## Overview

Session 076 documents NULL semantics for SQL JOIN operations.

JOIN conditions involving NULL values require explicit handling because NULL does not behave like an ordinary value in SQL equality comparisons.

In SQL, comparing NULL with another value using `=` does not produce a true match. Instead, the comparison evaluates to UNKNOWN.

This distinction is important for JOIN execution because records containing NULL join fields must not be treated as matching ordinary values.

## NULL JOIN Behavior

For an equality JOIN condition:

```sql
ON left.field = right.field
```

the following cases apply:

| Left value | Right value | Match |
| ---------- | ----------- | ----- |
| `100`      | `100`       | Yes   |
| `100`      | `200`       | No    |
| `NULL`     | `100`       | No    |
| `100`      | `NULL`      | No    |
| `NULL`     | `NULL`      | No    |

Two NULL values therefore do not match through ordinary equality JOIN semantics.

## Implementation

The JOIN execution layer preserves the existing field-resolution and comparison model while ensuring that unresolved or NULL-like join values do not incorrectly produce equality matches.

JOIN matching remains based on explicit condition evaluation rather than treating the absence of a value as an ordinary matching value.

This prevents accidental joins between records where both sides lack the requested field.

## Outer JOIN Considerations

NULL semantics are particularly important for outer joins.

A LEFT JOIN must retain an unmatched left-side record even when the right-side join field is missing or NULL.

Similarly, RIGHT JOIN and FULL OUTER JOIN operations must preserve unmatched records from their respective inputs.

The absence of a matching right-side record must not be confused with a successful NULL-to-NULL equality comparison.

## Example

Consider:

```text
Left:
account_id = NULL
name       = Alice

Right:
account_id = NULL
status     = Active
```

A normal equality JOIN using:

```sql
ON left.account_id = right.account_id
```

does not produce a match.

For a LEFT JOIN, the left record is therefore retained as an unmatched record.

## Testing

JOIN NULL semantics should cover:

* NULL on the left join field
* NULL on the right join field
* NULL on both join fields
* Missing join attributes
* Outer JOIN preservation of unmatched records
* Multiple records containing NULL join values
* Prevention of accidental NULL-to-NULL matches

## Result

Session 076 establishes explicit NULL-aware JOIN semantics and ensures that missing or NULL join values do not incorrectly satisfy equality conditions.

This provides a clearer foundation for subsequent JOIN features and edge-case handling.
