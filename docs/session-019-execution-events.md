concept

execution observability extends request-level logging into structured visibility of what happens inside a request

instead of only tracking HTTP metadata (method, status, duration), the system now tracks execution stages across the business pipeline

this is the first step from “request observability” → “execution observability”

How it works

HTTP Request
↓
Execution Context Middleware (injects request identity + timestamps)
↓
Handler (extracts execution context)
↓
Service Layer (business processing with context)
↓
Execution Event Logger (stage-based tracing)
↓
Processing Engine
↓
HTTP Response
↓
Request + execution trace printed to logs

What has been implemented
execution context middleware introduced (request ID + start time)
execution context injected into context.Context
handler updated to extract execution context via execution.FromContext
service layer updated to accept execution context explicitly
execution event system added (LogEvent)
execution stages introduced:
ingestion_started
records_received
ingestion_completed
end-to-end request tracing implemented across handler → service flow
Key outcomes
every request now has a unique execution identity
business logic is linked to request traceability
execution stages are visible in logs
first structured observability layer introduced beyond HTTP logging
clear separation between:
HTTP lifecycle tracking
execution lifecycle tracking
Why it matters

this is a foundational step toward:

distributed tracing systems
execution audit trails
financial-grade transaction observability
workflow replay systems
AI-native backend execution tracking

it moves the system from:

“what happened to the request?”

to:

“what happened inside the request?”
