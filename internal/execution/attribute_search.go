package execution

func FindByAttribute(spans []ExecutionContext, key, value string) []ExecutionContext {
	var matches []ExecutionContext

	for _, span := range spans {
		if span.Attributes[key] == value {
			matches = append(matches, span)
		}
	}

	return matches
}
