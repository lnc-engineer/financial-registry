# Session 036 – Span Duration Tracking

## Overview

This session enhances the tracing system by introducing execution duration visibility for each span in the trace tree.

While previous sessions established parent-child span relationships and execution metrics, this session adds time-based observability, allowing each span to report how long it took to execute.

## What Was Implemented

- Added `Duration()` method to `ExecutionContext`
- Calculated span duration using `StartTime` and `EndTime`
- Integrated duration display into trace tree rendering
- Updated output format to include execution timing per span

## Result

Trace output now includes execution timing per node:


TRACE TREE

├── ValidateRecord (span-1) [2ms]
├── ProcessRecord (span-2) [5ms]
└── StoreRecord (span-3) [1ms]


## Design Decisions

- Duration is calculated dynamically rather than stored to avoid state duplication
- ExecutionContext remains the single source of truth for span lifecycle data
- Trace rendering layer is responsible only for presentation, not computation

## Impact

This session significantly improves debugging and performance visibility across the ingestion pipeline.

It enables:
- Identification of slow spans
- Better understanding of pipeline bottlenecks
- Foundation for distributed tracing style observability

## Next Steps

Future improvements may include:
- Human-readable duration formatting (ms / µs / s scaling)
- Aggregated parent span timing
- Status propagation (success/failure bubbling)
- Structured trace export (JSON/OpenTelemetry style)
