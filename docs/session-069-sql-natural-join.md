# Session 069 — SQL NATURAL JOIN

## Overview

This session introduces SQL NATURAL JOIN behavior.

A NATURAL JOIN automatically joins two datasets using columns that share the same name, without requiring an explicit join condition.

## Key Concepts

- A NATURAL JOIN derives join fields from shared attribute names.
- Matching rows are joined when their shared field values are equal.
- Multiple matching rows are supported.
- Left-side attributes are preserved.
- Right-side attributes are preserved using the existing `right_` attribute convention.
- No explicit `JoinCondition` is required.

## Testing

The execution package will be verified with:

`go test ./...`
