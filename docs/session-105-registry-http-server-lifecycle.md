# Session 105 — Registry HTTP Server Lifecycle

## Overview

Session 105 adds explicit HTTP server lifecycle management to the registry
application.

The registry HTTP entry point introduced in Session 104 previously started the
server directly from `main()` and treated every return from `ListenAndServe()`
as a fatal error.

This session introduces:

- server lifecycle orchestration
- graceful shutdown
- SIGINT and SIGTERM handling
- a bounded shutdown timeout
- a testable `runServer` function
- lifecycle verification for the registry command

## Previous Lifecycle

Before Session 105, the registry command followed this simple flow:

```text
main
 |
 +-- app.NewHTTPServer(":8080")
 |
 +-- server.ListenAndServe()
 |
 +-- log.Fatal(error)