package execution

// CompositeFilter applies multiple trace filtering criteria.
type CompositeFilter struct {
	TraceID   string
	Status    string
	Lifecycle string

	TraceIDOperator   QueryOperator
	StatusOperator    QueryOperator
	LifecycleOperator QueryOperator
}

// Apply executes all configured filters using logical AND.
func (f CompositeFilter) Apply(spans []ExecutionContext) []ExecutionContext {
	filtered := spans

	if f.TraceID != "" {
		operator := f.TraceIDOperator

		if operator == "" {
			operator = OperatorEquals
		}

		var result []ExecutionContext

		for _, span := range filtered {
			if MatchValue(span.TraceID, f.TraceID, operator) {
				result = append(result, span)
			}
		}

		filtered = result
	}

	if f.Status != "" {
		operator := f.StatusOperator

		if operator == "" {
			operator = OperatorEquals
		}

		var result []ExecutionContext

		for _, span := range filtered {
			if MatchValue(span.Status, f.Status, operator) {
				result = append(result, span)
			}
		}

		filtered = result
	}

	if f.Lifecycle != "" {
		operator := f.LifecycleOperator

		if operator == "" {
			operator = OperatorEquals
		}

		var result []ExecutionContext

		for _, span := range filtered {
			for _, event := range span.Lifecycle {
				if MatchValue(event.Name, f.Lifecycle, operator) {
					result = append(result, span)
					break
				}
			}
		}

		filtered = result
	}

	return filtered
}
