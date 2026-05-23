concept

middleware enables cross-cutting concerns in backend systems to be handled outside of business logic

this improves separation of concerns, observability, and system maintainability

goal is to introduce request lifecycle visibility into the ingestion processor

How it works

HTTP Request
↓
Logging Middleware
↓
Status Recorder (captures response metadata)
↓
Handler (request processing)
↓
Service Layer (business orchestration)
↓
Processing Engine
↓
HTTP Response

What we implemented

- middleware package introduced
- logging middleware added
- response status tracking implemented via ResponseWriter wrapper
- request duration measurement added
- handler wrapped with middleware in main.go

Key outcomes

- every request is now observable
- status codes are captured centrally
- execution time is tracked per request
- business logic remains clean and isolated

Why it matters

this is a foundational step toward:
- distributed tracing
- execution observability
- audit logging systems
- financial-grade backend infrastructure
- AI-native workflow tracking systems