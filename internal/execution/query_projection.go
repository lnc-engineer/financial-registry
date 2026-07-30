package execution

type QueryProjection struct {
	Fields []string
}

func ApplyProjection(
	ctx ExecutionContext,
	projection QueryProjection,
) map[string]string {
	result := make(map[string]string)

	for _, field := range projection.Fields {
		switch field {
		case "trace_id":
			result[field] = ctx.TraceID
		case "span_id":
			result[field] = ctx.SpanID
		case "parent_span_id":
			result[field] = ctx.ParentSpanID
		case "span_name":
			result[field] = ctx.SpanName
		case "status":
			result[field] = ctx.Status
		default:
			if value, ok := ctx.Attributes[field]; ok {
				result[field] = value
			}
		}
	}

	return result
}

func ApplyProjectionToResults(
	results []ExecutionContext,
	projection QueryProjection,
) []map[string]string {
	projected := make([]map[string]string, 0, len(results))

	for _, result := range results {
		projected = append(projected, ApplyProjection(result, projection))
	}

	return projected
}
