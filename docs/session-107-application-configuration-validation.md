# Session 107 — Application Configuration Validation

**Date:** 28 September 2026

## Objective

Add an explicit validation boundary for application configuration.

Session 106 introduced the `Config` type and default HTTP address. Session 107 builds on that foundation by ensuring invalid configuration is rejected before it is used to construct the application server.

## Changes

### Configuration validation

Added:

```go
func (c Config) Validate() error