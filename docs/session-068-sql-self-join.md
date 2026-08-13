# Session 068 — SQL SELF JOIN

## Overview

This session continues the SQL JOIN work by documenting SELF JOIN behavior.

A SELF JOIN allows rows from the same dataset to be compared with or related to other rows from that dataset.

## Key Concepts

- A table can be joined to itself.
- The left and right sides represent the same underlying dataset.
- Join conditions determine which rows are related.
- Right-side attributes are preserved using the existing join attribute convention.
- Multiple matching rows are supported.

## Testing

The execution package was verified with:

```bash
go test ./...
