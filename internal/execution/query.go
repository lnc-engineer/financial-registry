package execution

func FindByTraceID(contexts []ExecutionContext, traceID string) []ExecutionContext {
	var result []ExecutionContext

	for _, ctx := range contexts {
		if ctx.TraceID == traceID {
			result = append(result, ctx)
		}
	}

	return result
}
