package execution

// FilterByTraceID returns all spans belonging to a trace.
func FilterByTraceID(spans []ExecutionContext, traceID string) []ExecutionContext {
	var filtered []ExecutionContext

	for _, span := range spans {
		if span.TraceID == traceID {
			filtered = append(filtered, span)
		}
	}

	return filtered
}

// FilterByStatus returns spans matching a status.
func FilterByStatus(spans []ExecutionContext, status string) []ExecutionContext {
	var filtered []ExecutionContext

	for _, span := range spans {
		if span.Status == status {
			filtered = append(filtered, span)
		}
	}

	return filtered
}

// FilterRootSpans returns spans without parents.
func FilterRootSpans(spans []ExecutionContext) []ExecutionContext {
	var filtered []ExecutionContext

	for _, span := range spans {
		if span.ParentSpanID == "" {
			filtered = append(filtered, span)
		}
	}

	return filtered
}
