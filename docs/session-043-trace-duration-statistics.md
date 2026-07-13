# Session 043 - Trace Duration Statistics

## Objective

Extend the tracing system by calculating execution duration metrics for spans. The goal was to provide additional performance insights alongside the existing trace statistics.

---

## What Was Added

### Duration Statistics

Introduced a `TraceStats` structure capable of tracking execution timing information across all recorded spans.

The statistics now include:

- Total spans
- Root spans
- Child spans
- Successful spans
- Failed spans
- Unknown status spans
- Average execution duration
- Longest execution duration
- Shortest execution duration

---

### Statistics Calculation

Implemented the `CalculateTraceStats` function to analyse a collection of execution spans.

During calculation the function:

- Counts root and child spans
- Groups spans by execution status
- Measures the duration of every span
- Calculates the average execution time
- Identifies the longest-running span
- Identifies the shortest-running span

This provides a concise performance summary for an execution trace.

---

## Testing

Added comprehensive unit tests to verify:

- Total span counting
- Root and child span detection
- Success, failure and unknown status counts
- Longest duration calculation
- Shortest duration calculation
- Average duration calculation

The tests use sample execution spans with known durations to ensure all calculated statistics are accurate.

---

## Result

The tracing package now produces both structural and performance statistics for execution traces. In addition to reporting span counts and execution outcomes, it can summarise timing information, making it easier to identify slow-running operations and analyse overall execution performance.
