package execution

// ApplyIntersect returns distinct left contexts whose field value
// also exists in the right contexts.
func ApplyIntersect(
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

		if rightValues[value] && !seen[value] {
			results = append(results, leftCtx)
			seen[value] = true
		}
	}

	return results
}
