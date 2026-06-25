Trace Query Support

Previous tracing sessions introduced TraceID, SpanID, parent-child relationships, and trace tree rendering. These improvements made it possible to represent execution flows as structured traces.

Session 034 introduces trace query support through the FindByTraceID helper function. This allows execution contexts belonging to a specific trace to be retrieved from a collection.

Example:

TraceID
├── Validation Span
├── Transformation Span
├── Storage Span
└── Notification Span

By filtering execution contexts using a TraceID, future observability features can inspect complete execution flows rather than individual spans in isolation.

This capability provides a foundation for trace inspection, debugging, analytics, and future trace statistics functionality.

