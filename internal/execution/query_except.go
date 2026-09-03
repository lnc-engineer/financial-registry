package execution

// ApplyExcept returns distinct left contexts whose field value
// does not exist in the right contexts.
func ApplyExcept(
	left []ExecutionContext,
	right []ExecutionContext,
	field string,
) []ExecutionContext {
	var results []ExecutionContext

	rightValues := make(map[string]bool)

	for _, rightCtx := range right {
		value := resolveField(rightCtx, field)
		rightValues[value] = true
	}

	seen := make(map[string]bool)

	for _, leftCtx := range left {
		value := resolveField(leftCtx, field)

		if !rightValues[value] && !seen[value] {
			results = append(results, leftCtx)
			seen[value] = true
		}
	}

	return results
}
