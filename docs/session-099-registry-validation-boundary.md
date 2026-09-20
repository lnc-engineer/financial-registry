# Session 099 — Registry Validation Boundary

## Objective

Define the validation boundary for the Financial Systems Registry and document how validation fits between registry domain behavior, application services, and future execution workflows.

The goal is to establish a clear architectural boundary before introducing more validation capabilities.

## Scope

This session is documentation-only.

No production code or tests are changed in this session.

## Current Architecture

The registry currently separates several responsibilities:

* `System` represents the registry domain entity.
* `Registry` manages the in-memory collection of registered systems.
* `SystemRepository` provides a persistence-oriented abstraction for system storage.
* `SystemService` provides an application/service-layer boundary for system operations.

The current structure provides a foundation for introducing validation as a distinct concern.

## Validation Boundary

Validation is responsible for determining whether a registry system satisfies the requirements necessary for a particular operation.

Validation should remain separate from persistence and execution.

The intended boundary is:

```text
API / External Caller
        |
        v
SystemService
        |
        +------> Validation
        |
        v
SystemRepository
        |
        v
Persistence
```

The service layer coordinates the operation, while validation determines whether the relevant requirements are satisfied.

## Domain Validation

Some validation belongs directly to the domain model.

The existing `System.Validate()` method is an example of domain-level validation.

It protects basic system invariants such as:

* required system identity
* required system name
* valid system status

These rules describe whether a `System` is structurally valid as a domain entity.

Domain validation should remain close to the domain model because these rules apply regardless of how the system is accessed.

## Application Validation

Other validation concerns depend on the operation being performed.

Examples include:

* whether a system may be registered
* whether a requested system exists
* whether an operation is allowed for the current system state
* whether required application inputs are present
* whether a system satisfies prerequisites for a future execution workflow

These checks belong at the application/service boundary when they require coordination between multiple components.

## Execution Validation

Future execution workflows may require additional validation before a registered system can be executed.

Potential examples include:

* required execution configuration
* supported execution capabilities
* dependency availability
* input compatibility
* execution permissions
* runtime prerequisites

These checks should not be placed inside the basic `System` domain entity unless they are genuine domain invariants.

This keeps the registry model small while allowing execution-specific validation to evolve independently.

## Validation and Trust

Validation is an important part of the future Registry + Runtime + Trust Layer architecture.

The registry should not only answer:

> "Is this system registered?"

It should eventually be able to support questions such as:

> "Is this system valid?"

and:

> "Is this system ready for the requested operation?"

These are different questions and should remain distinguishable.

Registration establishes identity and availability within the registry.

Validation establishes whether defined requirements are satisfied.

Execution establishes whether the system can actually participate in a particular runtime workflow.

## Error Handling

Validation failures should remain distinguishable from infrastructure failures.

For example:

```text
Validation failure
        |
        v
Invalid system or operation requirements

Repository failure
        |
        v
Storage or persistence problem

Execution failure
        |
        v
Runtime or downstream problem
```

Keeping these categories separate will support clearer API responses, testing, observability, and future operational diagnostics.

## Architectural Principles

The validation boundary follows several principles:

1. Keep domain invariants close to the domain model.
2. Keep application-specific checks at the service boundary.
3. Keep execution-specific requirements outside the basic registry entity.
4. Avoid coupling validation directly to persistence.
5. Preserve meaningful validation errors.
6. Introduce additional validation only when required by actual system behavior.

## Relationship to the Broader Registry Vision

The Financial Systems Registry is intended to evolve toward a platform combining:

* registry
* validation
* execution
* orchestration
* observability
* trust

A clear validation boundary allows these capabilities to grow without turning the registry domain model into a large collection of unrelated responsibilities.

The long-term direction is:

```text
Registry
   |
   v
Validation
   |
   v
Execution
   |
   v
Observability / Trust
```

The current system remains intentionally small.

Future sessions can introduce concrete validation abstractions and implementations when they are justified by the next application requirement.

## Outcome

Session 099 establishes the validation boundary for the Financial Systems Registry.

No production code is changed.

The architecture now distinguishes:

* domain validation
* application validation
* execution validation

This provides a foundation for future validation, execution, observability, and trust capabilities without prematurely increasing the complexity of the registry core.
