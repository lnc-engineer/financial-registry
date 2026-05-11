# Session 16 — Multi-File Architecture

## Goal

Refactor the ingestion processor from a single-file backend
into a cleaner multi-file architecture.

---

## Why This Matters

Large Go applications should not keep all logic inside `main.go`.

Separating files improves:

- readability
- maintainability
- scalability
- debugging
- backend organization

This introduces the concept of separation of concerns.

---

## New Project Structure

cmd/ingestion-processor/

main.go
models.go
processor.go
handler.go

---

## File Responsibilities

### main.go

Responsible only for:

- starting the HTTP server
- registering routes

Example:

```go
http.HandleFunc("/process", processHandler)
