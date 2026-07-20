package execution

// CompositeFilter applies multiple trace filtering criteria.
type CompositeFilter struct {
	TraceID string
	Status  string
}

// Apply executes all configured filters using logical AND.
func (f CompositeFilter) Apply(spans []ExecutionContext) []ExecutionContext {
	filtered := spans

	if f.TraceID != "" {
		filtered = FilterByTraceID(filtered, f.TraceID)
	}

	if f.Status != "" {
		filtered = FilterByStatus(filtered, f.Status)
	}

	return filtered
}
