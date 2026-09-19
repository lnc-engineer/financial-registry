# Session 098 — Registry Service Error Handling

## Objective

Document the error-handling direction for the registry service layer introduced in Sessions 096 and 097.

The goal is to establish a clear boundary between domain, repository, and service-layer errors before additional application behavior is introduced.

## Scope

This session is documentation-only.

No production code or tests are changed in this session.

## Current Registry Architecture

The registry currently separates several responsibilities:

* `System` represents the registry domain entity.
* `Registry` manages the in-memory collection of registered systems.
* `SystemRepository` provides a persistence-oriented abstraction for system storage.
* `SystemService` provides an application/service-layer boundary for system operations.

The service layer sits above the repository and provides a stable application-facing boundary.

## Error Handling Responsibility

Errors should remain meaningful as they move between architectural layers.

The service layer should not unnecessarily hide or replace errors originating from lower layers.

For example, repository errors may indicate:

* a system does not exist
* a system already exists
* persistence failed
* an underlying storage operation failed

Domain validation errors may indicate:

* a missing system ID
* a missing system name
* an invalid system status
* an invalid domain state transition

These errors provide useful information to callers and should remain distinguishable where that distinction is important.

## Service Layer Direction

The `SystemService` should coordinate application behavior without becoming responsible for domain rules that belong elsewhere.

The intended responsibility boundary is:

```text
System
    |
    | domain rules
    v
SystemRepository
    |
    | storage abstraction
    v
SystemService
    |
    | application orchestration
    v
API / External Caller
```

The service layer may translate or enrich errors when necessary for an application-facing contract, but it should avoid obscuring the original cause without a clear reason.

## Future Error Design

As the registry becomes more production-oriented, error handling may evolve toward:

* sentinel errors for important domain conditions
* typed errors where callers need structured information
* error wrapping using Go's standard error mechanisms
* consistent error classification at API boundaries
* observability attributes associated with important failures
* preservation of underlying repository and persistence errors

These capabilities should be introduced when they are required by actual application behavior rather than prematurely adding complexity.

## Architectural Principle

The registry should preserve useful failure information across its layers.

A caller should be able to determine whether an operation failed because of:

1. invalid domain input
2. an invalid domain state
3. a missing or duplicate registry system
4. a repository failure
5. a persistence or infrastructure failure

Keeping these categories distinguishable will make the system easier to test, observe, integrate, and operate as the registry grows.

## Relationship to the Broader Registry Vision

Error handling is part of the trust boundary of the Financial Systems Registry.

The long-term platform is intended to provide reliable registry, execution, validation, observability, and orchestration capabilities for financial systems.

Reliable error propagation supports that direction by ensuring that failures are:

* understandable
* testable
* observable
* diagnosable
* suitable for API consumers

The current implementation remains intentionally small. Future sessions can introduce concrete error types and handling mechanisms when they are justified by the next application requirement.

## Outcome

Session 098 establishes the architectural direction for registry service-layer error handling.

No production code is changed.

The registry now has a documented principle that service-layer operations should preserve meaningful domain and repository failure information while keeping application orchestration separate from lower-level responsibilities.
