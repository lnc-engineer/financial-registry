# Session 085 — SQL EXCEPT Ordering and Result Semantics

## Overview

Session 085 extends SQL `EXCEPT` coverage with tests focused on result ordering and set-difference semantics.

The existing `EXCEPT` implementation accepts a left result set, a right result set, and the field used to determine matching rows. This session verifies that the operation behaves consistently across ordered, empty, identical, disjoint, and chained result sets.

## Scope

The session covers:

- Preservation of left-side result order.
- Removal of matching rows from the left result.
- Behaviour when both result sets are identical.
- Behaviour when result sets are disjoint.
- Empty left-side results.
- Empty right-side results.
- Chained `EXCEPT` operations.

## Result Ordering

`EXCEPT` preserves the order of rows that remain from the left-hand result set.

For example, given:

```text
Left:
txn-003
txn-001
txn-004
txn-002

Right:
txn-004
