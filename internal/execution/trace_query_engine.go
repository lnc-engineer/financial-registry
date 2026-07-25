package execution

type QueryCondition struct {
	Field    string
	Operator QueryOperator
	Value    string
}

func MatchesQuery(ctx ExecutionContext, conditions []QueryCondition) bool {
	for _, condition := range conditions {
		actual := resolveField(ctx, condition.Field)

		if !MatchValue(actual, condition.Value, condition.Operator) {
			return false
		}
	}

	return true
}

func resolveField(ctx ExecutionContext, field string) string {
	switch field {
	case "trace_id":
		return ctx.TraceID

	case "span_id":
		return ctx.SpanID

	case "parent_span_id":
		return ctx.ParentSpanID

	case "span_name":
		return ctx.SpanName

	case "status":
		return ctx.Status

	default:
		if value, ok := ctx.Attributes[field]; ok {
			return value
		}

		return ""
	}
}
