# Response Status Tracking Middleware

## Purpose

The response status tracking middleware provides visibility into
the HTTP status codes returned by the API.

This enables the system to distinguish successful requests from
failed requests and contributes to overall observability.

---

## Why It Exists

The standard `http.ResponseWriter` does not expose the final
status code after a response is written.

Without additional tracking, middleware cannot determine whether
a request completed successfully or returned an error.

To solve this problem, the application wraps the
`http.ResponseWriter` and captures status codes as they are
written.

---

## How It Works

1. A custom response writer wraps the original
   `http.ResponseWriter`.

2. The wrapper intercepts calls to `WriteHeader()`.

3. The HTTP status code is recorded before being passed to the
   underlying writer.

4. Middleware can access the captured status code after request
   processing completes.

---

## Observability Benefits

Response status tracking enables:

* Success and failure counting
* Error rate monitoring
* Request outcome visibility
* Future alerting and reporting capabilities

This information complements request duration tracking and
execution metrics.

---

## Architectural Importance

The middleware establishes a reusable observability pattern.

Request processing can now be evaluated using:

* Request counts
* Response status codes
* Request duration
* Execution events

Together these capabilities provide a stronger foundation for
execution monitoring and operational visibility.
