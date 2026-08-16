# Session 070 — SQL JOIN USING

## Overview

This session introduces SQL JOIN USING behavior.

A JOIN USING explicitly specifies one or more columns that should be used to match rows between two datasets.

## Key Concepts

- JOIN USING specifies the shared join column explicitly.
- Matching rows are joined when the specified field values are equal.
- Multiple matching rows are supported.
- Unlike NATURAL JOIN, the join fields are explicitly defined.
- Left-side attributes are preserved.
- Right-side attributes are preserved using the existing `right_` attribute convention.
- JOIN USING avoids requiring fully qualified join expressions.

## Testing

The execution package will be verified with:

`go test ./...`
