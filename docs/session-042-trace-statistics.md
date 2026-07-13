# Session 042 - Trace Statistics

## Objective

Improve the tracing system by introducing aggregated trace statistics that provide a concise overview of execution activity. This allows trace summaries to present useful metrics instead of requiring manual inspection of every span.

---

## What Was Added

### Trace Statistics

Implemented a statistics model capable of summarising a trace.

Statistics include:

- Total spans
- Root spans
- Leaf spans
- Maximum tree depth
- Successful spans
- Failed spans
- Unknown status spans

---

### Summary Generation

Extended the trace summary functionality so statistics are calculated automatically from the execution trace.

The summary now reports:

- Overall span counts
- Success and failure totals
- Tree depth information
- Root and leaf counts

This provides a quick health overview of the trace.

---

### Testing

Added unit tests covering:

- Empty traces
- Single-span traces
- Nested span trees
- Mixed success and failure states
- Depth calculation
- Root and leaf counting

These tests help ensure future tracing changes do not break summary calculations.

---

## Result

The tracing package now provides both structural trace information and high-level execution statistics, making trace output easier to understand while maintaining automated test coverage.
