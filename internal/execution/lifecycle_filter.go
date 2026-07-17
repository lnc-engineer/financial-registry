package execution

// FilterByLifecycle returns spans containing a lifecycle event name.
func FilterByLifecycle(spans []ExecutionContext, lifecycle string) []ExecutionContext {
	var filtered []ExecutionContext

	for _, span := range spans {
		for _, event := range span.Lifecycle {
			if event.Name == lifecycle {
				filtered = append(filtered, span)
				break
			}
		}
	}

	return filtered
}
