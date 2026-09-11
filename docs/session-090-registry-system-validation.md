# Session 090 — Registry System Validation

## Overview

Session 090 adds domain validation for the registry `System` model introduced in Session 089.

The goal is to establish a small, reliable validation boundary around registered systems before introducing persistence or API layers.

## System Validation Rules

A `System` is considered valid when:

* `ID` is present.
* `Name` is present.
* `Status` is either `active` or `inactive`.
* `Description` may be empty because it is optional.

## Implementation

Added:

```text
internal/registry/system_validation.go
```

The file provides:

```go
func (s System) Validate() error
```

The method returns an error when required fields are missing or when an unsupported system status is supplied.

Valid statuses are defined by the existing domain constants:

```go
SystemStatusActive
SystemStatusInactive
```

## Tests

Added:

```text
internal/registry/system_validation_test.go
```

The tests cover:

* Valid system acceptance.
* Missing system ID.
* Missing system name.
* Invalid system status.
* Valid system without a description.

## Verification

The complete Go test suite was executed:

```bash
go test ./...
```

All packages passed successfully.

## Result

The registry domain now has a basic validation boundary for `System` objects.

This keeps validation inside the domain layer rather than placing domain rules inside future API, persistence, or transport code.

## Next Direction

The next registry-focused work can build on this foundation by introducing a repository/storage abstraction for registered systems, while keeping domain validation independent from persistence concerns.
