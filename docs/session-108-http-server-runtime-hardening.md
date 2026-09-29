# Session 108 — HTTP Server Runtime Hardening

## Objective

Harden the registry application's HTTP server boundary by configuring explicit
runtime timeouts.

The application already provides HTTP server construction through the
application package. This session makes the server runtime behavior more
explicit by configuring request and connection timeout limits.

## Changes

### HTTP server timeouts

`NewHTTPServer` now configures four HTTP server timeout values:

- `ReadHeaderTimeout`: 5 seconds
- `ReadTimeout`: 10 seconds
- `WriteTimeout`: 10 seconds
- `IdleTimeout`: 60 seconds

The timeout values are defined as package-level constants so the runtime
policy is centralized and visible in the application boundary.

### Tests

The HTTP server tests now verify:

- the configured server address
- the HTTP handler is present
- the read-header timeout
- the read timeout
- the write timeout
- the idle timeout

This ensures the runtime configuration remains explicit and protected against
accidental removal or change.

## Architecture

The HTTP server remains composed through the existing application boundaries:

```text
Config
  ↓
NewConfiguredHTTPServer
  ↓
NewHTTPServer
  ↓
NewHTTPHandler
  ↓
SystemHandler
  ↓
SystemService
  ↓
Registry